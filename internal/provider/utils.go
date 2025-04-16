// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	tfReflect "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/provider/reflect"
)

func debugResponse(response *http.Response) string {
	if v := response.Request.Header.Get("Authorization"); v != "" {
		response.Request.Header.Set("Authorization", "(sensitive)")
	}
	dumpReq, err := httputil.DumpRequest(response.Request, true)
	if err != nil {
		dumpReq, err = httputil.DumpRequest(response.Request, false)
		if err != nil {
			return err.Error()
		}
	}
	dumpRes, err := httputil.DumpResponse(response, true)
	if err != nil {
		dumpRes, err = httputil.DumpResponse(response, false)
		if err != nil {
			return err.Error()
		}
	}
	return fmt.Sprintf("**Request**:\n%s\n**Response**:\n%s", string(dumpReq), string(dumpRes))
}

func merge(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, target interface{}) {
	var plan types.Object
	var state types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	val, err := state.ToTerraformValue(ctx)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Object Conversion Error", "An unexpected error was encountered trying to convert object. This is always an error in the provider. Please report the following to the provider developer:\n\n"+err.Error()))
		return
	}
	resp.Diagnostics.Append(tfReflect.Into(ctx, types.ObjectType{AttrTypes: state.AttributeTypes(ctx)}, val, target, tfReflect.Options{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	}, path.Empty())...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, target)...)
}

func refreshPlan(ctx context.Context, plan types.Object, target any) diag.Diagnostics {
	var diags diag.Diagnostics

	obj := types.ObjectType{AttrTypes: plan.AttributeTypes(ctx)}
	val, err := plan.ToTerraformValue(ctx)
	if err != nil {
		diags.AddError(
			"Object Conversion Error",
			"An unexpected error was encountered trying to convert object. This is always an error in the provider. Please report the following to the provider developer:\n\n"+err.Error(),
		)
		return diags
	}

	diags.Append(tfReflect.Into(ctx, obj, val, target, tfReflect.Options{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
		SourceType:              tfReflect.SourceTypePlan,
	}, path.Empty())...)

	return diags
}

// Configurable options for the provider HTTP transport.
type ProviderHTTPTransportOpts struct {
	// HTTP headers to set on all requests.
	SetHeaders map[string]string

	// Underlying HTTP transport.
	Transport http.RoundTripper
}

// Note: this is taken as a more minimal/specific version of https://github.com/hashicorp/terraform-plugin-sdk/blob/main/helper/logging/logging_http_transport.go
func NewProviderHTTPTransport(opts ProviderHTTPTransportOpts) *providerHttpTransport {
	return &providerHttpTransport{
		setHeaders: opts.SetHeaders,
		transport:  opts.Transport,
	}
}

const (
	FieldHttpOperationType        = "tf_http_op_type"
	OperationHttpRequest          = "request"
	OperationHttpResponse         = "response"
	FieldHttpRequestMethod        = "tf_http_req_method"
	FieldHttpRequestUri           = "tf_http_req_uri"
	FieldHttpRequestProtoVersion  = "tf_http_req_version"
	FieldHttpRequestBody          = "tf_http_req_body"
	FieldHttpResponseProtoVersion = "tf_http_res_version"
	FieldHttpResponseStatusCode   = "tf_http_res_status_code"
	FieldHttpResponseStatusReason = "tf_http_res_status_reason"
	FieldHttpResponseBody         = "tf_http_res_body"
	FieldHttpTransactionId        = "tf_http_trans_id"
)

type providerHttpTransport struct {
	setHeaders map[string]string
	transport  http.RoundTripper
}

func (t *providerHttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	ctx = t.addTransactionIdField(ctx)

	// Set globally defined HTTP headers in the request
	t.setRequestHeaders(req)

	// Decompose the request bytes in a message (HTTP body) and fields (HTTP headers), then log it
	fields, err := decomposeRequestForLogging(req)
	if err != nil {
		tflog.Error(ctx, "Failed to parse request bytes for logging", []map[string]interface{}{map[string]interface{}{
			"error": err,
		}}...)
	} else {
		tflog.Debug(ctx, "Sending HTTP Request", []map[string]interface{}{fields}...)
	}

	// Invoke the wrapped RoundTrip now
	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return res, err
	}

	// Decompose the response bytes in a message (HTTP body) and fields (HTTP headers), then log it
	fields, err = decomposeResponseForLogging(res)
	if err != nil {
		tflog.Error(ctx, "Failed to parse response bytes for logging", []map[string]interface{}{map[string]interface{}{
			"error": err,
		}}...)
	} else {
		tflog.Debug(ctx, "Received HTTP Response", []map[string]interface{}{fields}...)
	}

	return res, nil
}

// Generates UUID and sets it into the tf_http_trans_id logging field.
func (t *providerHttpTransport) addTransactionIdField(ctx context.Context) context.Context {
	tId, err := uuid.GenerateUUID()

	if err != nil {
		tId = "Unable to assign Transaction ID: " + err.Error()
	}

	return tflog.SetField(ctx, FieldHttpTransactionId, tId)
}

// Sets globally defined HTTP headers in the request.
func (t *providerHttpTransport) setRequestHeaders(req *http.Request) {
	for name, value := range t.setHeaders {
		req.Header.Set(name, value)
	}
}

func decomposeRequestForLogging(req *http.Request) (map[string]interface{}, error) {
	fields := make(map[string]interface{}, len(req.Header)+4)
	fields[FieldHttpOperationType] = OperationHttpRequest

	fields[FieldHttpRequestMethod] = req.Method
	fields[FieldHttpRequestUri] = req.URL.RequestURI()
	fields[FieldHttpRequestProtoVersion] = req.Proto

	// Get the full body of the request, including headers appended by http.Transport:
	// this is necessary because the http.Request at this stage doesn't contain
	// all the headers that will be eventually sent.
	// We rely on `httputil.DumpRequestOut` to obtain the actual bytes that will be sent out.
	reqBytes, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}

	// Create a reader around the request full body
	reqReader := textproto.NewReader(bufio.NewReader(bytes.NewReader(reqBytes)))

	err = fieldHeadersFromRequestReader(reqReader, fields)
	if err != nil {
		return nil, err
	}

	// Read the rest of the body content
	fields[FieldHttpRequestBody] = bodyFromRestOfRequestReader(reqReader)
	return fields, nil
}

func fieldHeadersFromRequestReader(reader *textproto.Reader, fields map[string]interface{}) error {
	// Ignore the first line: it contains non-header content
	// that we have already captured.
	// Skipping this step, would cause the following call to `ReadMIMEHeader()`
	// to fail as it cannot parse the first line.
	_, err := reader.ReadLine()
	if err != nil {
		return err
	}

	// Read the MIME-style headers
	mimeHeader, err := reader.ReadMIMEHeader()
	if err != nil {
		return err
	}

	// Set the headers as fields to log
	for k, v := range mimeHeader {
		if len(v) == 1 {
			fields[k] = v[0]
		} else {
			fields[k] = v
		}
	}
	if _, ok := fields["Authorization"]; ok {
		fields["Authorization"] = "(sensitive)"
	}

	return nil
}

func bodyFromRestOfRequestReader(reader *textproto.Reader) string {
	var builder strings.Builder
	for {
		line, err := reader.ReadContinuedLine()
		if errors.Is(err, io.EOF) {
			break
		}
		builder.WriteString(line)
	}

	return builder.String()
}

func decomposeResponseForLogging(res *http.Response) (map[string]interface{}, error) {
	fields := make(map[string]interface{}, len(res.Header)+4)
	fields[FieldHttpOperationType] = OperationHttpResponse

	fields[FieldHttpResponseProtoVersion] = res.Proto
	fields[FieldHttpResponseStatusCode] = res.StatusCode
	fields[FieldHttpResponseStatusReason] = res.Status

	// Set the headers as fields to log
	for k, v := range res.Header {
		if len(v) == 1 {
			fields[k] = v[0]
		} else {
			fields[k] = v
		}
	}

	// Read the whole response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Wrap the bytes from the response body, back into an io.ReadCloser,
	// to respect the interface of http.Response, as expected by users of the
	// http.Client
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))

	fields[FieldHttpResponseBody] = string(resBody)

	return fields, nil
}
