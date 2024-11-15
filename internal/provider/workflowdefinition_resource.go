// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/validators"
	speakeasy_numbervalidators "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/validators/numbervalidators"
	speakeasy_objectvalidators "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WorkflowDefinitionResource{}
var _ resource.ResourceWithImportState = &WorkflowDefinitionResource{}

func NewWorkflowDefinitionResource() resource.Resource {
	return &WorkflowDefinitionResource{}
}

// WorkflowDefinitionResource defines the resource implementation.
type WorkflowDefinitionResource struct {
	client *sdk.SDK
}

// WorkflowDefinitionResourceModel describes the resource data model.
type WorkflowDefinitionResourceModel struct {
	AssignedTo             []types.String                   `tfsdk:"assigned_to"`
	ClosingReasons         []tfTypes.ClosingReasonID        `tfsdk:"closing_reasons"`
	CreationTime           types.String                     `tfsdk:"creation_time"`
	Description            types.String                     `tfsdk:"description"`
	DueDate                types.String                     `tfsdk:"due_date"`
	DynamicDueDate         *tfTypes.DynamicDueDate          `tfsdk:"dynamic_due_date"`
	EnableECPWorkflow      types.Bool                       `tfsdk:"enable_ecp_workflow"`
	Enabled                types.Bool                       `tfsdk:"enabled"`
	Flow                   types.String                     `tfsdk:"flow"`
	ID                     types.String                     `tfsdk:"id"`
	LastUpdateTime         types.String                     `tfsdk:"last_update_time"`
	Name                   types.String                     `tfsdk:"name"`
	Taxonomies             []types.String                   `tfsdk:"taxonomies"`
	UpdateEntityAttributes []tfTypes.UpdateEntityAttributes `tfsdk:"update_entity_attributes"`
	UserIds                []types.Number                   `tfsdk:"user_ids"`
}

func (r *WorkflowDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_definition"
}

func (r *WorkflowDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "WorkflowDefinition Resource",
		Attributes: map[string]schema.Attribute{
			"assigned_to": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"closing_reasons": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
					},
				},
			},
			"creation_time": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `ISO String Date & Time`,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"due_date": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"dynamic_due_date": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"action_type_condition": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null; must be one of ["WORKFLOW_STARTED", "STEP_CLOSED"]`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf(
								"WORKFLOW_STARTED",
								"STEP_CLOSED",
							),
						},
					},
					"number_of_units": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null`,
						Validators: []validator.Number{
							speakeasy_numbervalidators.NotNull(),
						},
					},
					"step_id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"time_period": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Not Null; must be one of ["minutes", "hours", "days", "weeks", "months"]`,
						Validators: []validator.String{
							speakeasy_stringvalidators.NotNull(),
							stringvalidator.OneOf(
								"minutes",
								"hours",
								"days",
								"weeks",
								"months",
							),
						},
					},
				},
				Description: `set a Duedate for a step then a specific`,
			},
			"enable_ecp_workflow": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Indicates whether this workflow is available for End Customer Portal or not. By default it's not.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Description: `Whether the workflow is enabled or not. Default: true`,
			},
			"flow": schema.StringAttribute{
				Required:    true,
				Description: `Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"last_update_time": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `ISO String Date & Time`,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"taxonomies": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `Taxonomy ids that are associated with this workflow and used for filtering`,
			},
			"update_entity_attributes": schema.ListNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null; must be one of ["workflow_status", "current_section", "current_step"]`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
								stringvalidator.OneOf(
									"workflow_status",
									"current_section",
									"current_step",
								),
							},
						},
						"target": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"entity_attribute": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
								"entity_schema": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
							},
							Description: `Not Null`,
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
						},
					},
				},
			},
			"user_ids": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.NumberType,
				Description: `This field is deprecated. Please use assignedTo`,
			},
		},
	}
}

func (r *WorkflowDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WorkflowDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkflowDefinitionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedWorkflowDefinition()
	res, err := r.client.Workflows.CreateDefinition(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.WorkflowDefinition != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedWorkflowDefinition(res.WorkflowDefinition)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkflowDefinitionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var definitionID string
	definitionID = data.ID.ValueString()

	request := operations.GetDefinitionRequest{
		DefinitionID: definitionID,
	}
	res, err := r.client.Workflows.GetDefinition(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.WorkflowDefinition != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedWorkflowDefinition(res.WorkflowDefinition)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkflowDefinitionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	workflowDefinition := *data.ToSharedWorkflowDefinition()
	var definitionID string
	definitionID = data.ID.ValueString()

	request := operations.UpdateDefinitionRequest{
		WorkflowDefinition: workflowDefinition,
		DefinitionID:       definitionID,
	}
	res, err := r.client.Workflows.UpdateDefinition(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.WorkflowDefinition != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedWorkflowDefinition(res.WorkflowDefinition)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkflowDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkflowDefinitionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var definitionID string
	definitionID = data.ID.ValueString()

	request := operations.DeleteDefinitionRequest{
		DefinitionID: definitionID,
	}
	res, err := r.client.Workflows.DeleteDefinition(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *WorkflowDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
