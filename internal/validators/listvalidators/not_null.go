// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package listvalidators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.List = ListNotNullValidator{}

// ListNotNullValidator validates that an attribute is not null. Most
// attributes should set Required: true instead, however in certain scenarios,
// such as a computed nested attribute, all underlying attributes must also be
// computed for planning to not show unexpected differences.
type ListNotNullValidator struct{}

// Description describes the validation in plain text formatting.
func (v ListNotNullValidator) Description(_ context.Context) string {
	return "value must be configured"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v ListNotNullValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v ListNotNullValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing Attribute Value",
		req.Path.String()+": "+v.Description(ctx),
	)
}

// NotNull returns an validator which ensures that the attribute is
// configured. Most attributes should set Required: true instead, however in
// certain scenarios, such as a computed nested attribute, all underlying
// attributes must also be computed for planning to not show unexpected
// differences.
func NotNull() validator.List {
	return ListNotNullValidator{}
}
