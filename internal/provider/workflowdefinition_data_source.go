// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &WorkflowDefinitionDataSource{}
var _ datasource.DataSourceWithConfigure = &WorkflowDefinitionDataSource{}

func NewWorkflowDefinitionDataSource() datasource.DataSource {
	return &WorkflowDefinitionDataSource{}
}

// WorkflowDefinitionDataSource is the data source implementation.
type WorkflowDefinitionDataSource struct {
	client *sdk.SDK
}

// WorkflowDefinitionDataSourceModel describes the data model.
type WorkflowDefinitionDataSourceModel struct {
	AssignedTo             []types.String                   `tfsdk:"assigned_to"`
	ClosingReasons         []tfTypes.ClosingReasonID        `tfsdk:"closing_reasons"`
	CreationTime           types.String                     `tfsdk:"creation_time"`
	Description            types.String                     `tfsdk:"description"`
	DueDate                types.String                     `tfsdk:"due_date"`
	DynamicDueDate         *tfTypes.DynamicDueDate          `tfsdk:"dynamic_due_date"`
	EnableECPWorkflow      types.Bool                       `tfsdk:"enable_ecp_workflow"`
	Flow                   types.String                     `tfsdk:"flow"`
	ID                     types.String                     `tfsdk:"id"`
	LastUpdateTime         types.String                     `tfsdk:"last_update_time"`
	Name                   types.String                     `tfsdk:"name"`
	UpdateEntityAttributes []tfTypes.UpdateEntityAttributes `tfsdk:"update_entity_attributes"`
	UserIds                []types.Number                   `tfsdk:"user_ids"`
}

// Metadata returns the data source type name.
func (r *WorkflowDefinitionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_definition"
}

// Schema defines the schema for the data source.
func (r *WorkflowDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "WorkflowDefinition DataSource",

		Attributes: map[string]schema.Attribute{
			"assigned_to": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"closing_reasons": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"creation_time": schema.StringAttribute{
				Computed:    true,
				Description: `ISO String Date & Time`,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"due_date": schema.StringAttribute{
				Computed: true,
			},
			"dynamic_due_date": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"action_type_condition": schema.StringAttribute{
						Computed:    true,
						Description: `must be one of ["WORKFLOW_STARTED", "STEP_CLOSED"]`,
					},
					"number_of_units": schema.NumberAttribute{
						Computed: true,
					},
					"step_id": schema.StringAttribute{
						Computed: true,
					},
					"time_period": schema.StringAttribute{
						Computed:    true,
						Description: `must be one of ["days", "weeks", "months"]`,
					},
				},
				Description: `set a Duedate for a step then a specific`,
			},
			"enable_ecp_workflow": schema.BoolAttribute{
				Computed:    true,
				Description: `Indicates whether this workflow is available for End Customer Portal or not. By default it's not.`,
			},
			"flow": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"last_update_time": schema.StringAttribute{
				Computed:    true,
				Description: `ISO String Date & Time`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"update_entity_attributes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["workflow_status", "current_section", "current_step"]`,
						},
						"target": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"entity_attribute": schema.StringAttribute{
									Computed: true,
								},
								"entity_schema": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
			"user_ids": schema.ListAttribute{
				Computed:    true,
				ElementType: types.NumberType,
				Description: `This field is deprecated. Please use assignedTo`,
			},
		},
	}
}

func (r *WorkflowDefinitionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WorkflowDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *WorkflowDefinitionDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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

	definitionID := data.ID.ValueString()
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
