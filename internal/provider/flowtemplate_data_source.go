// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

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
var _ datasource.DataSource = &FlowTemplateDataSource{}
var _ datasource.DataSourceWithConfigure = &FlowTemplateDataSource{}

func NewFlowTemplateDataSource() datasource.DataSource {
	return &FlowTemplateDataSource{}
}

// FlowTemplateDataSource is the data source implementation.
type FlowTemplateDataSource struct {
	client *sdk.SDK
}

// FlowTemplateDataSourceModel describes the data model.
type FlowTemplateDataSourceModel struct {
	AssignedTo             []types.String                   `tfsdk:"assigned_to"`
	AvailableInEcp         types.Bool                       `tfsdk:"available_in_ecp"`
	ClosingReasons         []tfTypes.ClosingReason1         `tfsdk:"closing_reasons"`
	CreatedAt              types.String                     `tfsdk:"created_at"`
	Description            types.String                     `tfsdk:"description"`
	DueDate                types.String                     `tfsdk:"due_date"`
	DueDateConfig          *tfTypes.DueDateConfig           `tfsdk:"due_date_config"`
	Edges                  []tfTypes.Edge                   `tfsdk:"edges"`
	Enabled                types.Bool                       `tfsdk:"enabled"`
	ID                     types.String                     `tfsdk:"id"`
	IsFlowMigrated         types.Bool                       `tfsdk:"is_flow_migrated"`
	Name                   types.String                     `tfsdk:"name"`
	OrgID                  types.String                     `tfsdk:"org_id"`
	Phases                 []tfTypes.Phase                  `tfsdk:"phases"`
	Tasks                  []tfTypes.Task                   `tfsdk:"tasks"`
	Taxonomies             []types.String                   `tfsdk:"taxonomies"`
	Trigger                *tfTypes.Trigger                 `tfsdk:"trigger"`
	UpdateEntityAttributes []tfTypes.UpdateEntityAttributes `tfsdk:"update_entity_attributes"`
	UpdatedAt              types.String                     `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *FlowTemplateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flow_template"
}

// Schema defines the schema for the data source.
func (r *FlowTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "FlowTemplate DataSource",

		Attributes: map[string]schema.Attribute{
			"assigned_to": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"available_in_ecp": schema.BoolAttribute{
				Computed:    true,
				Description: `Indicates whether this workflow is available for End Customer Portal or not. By default it's not.`,
			},
			"closing_reasons": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"creation_time": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"last_update_time": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"title": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `ISO String Date & Time`,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"due_date": schema.StringAttribute{
				Computed: true,
			},
			"due_date_config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"duration": schema.NumberAttribute{
						Computed: true,
					},
					"phase_id": schema.StringAttribute{
						Computed: true,
					},
					"task_id": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Computed: true,
					},
					"unit": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Set due date for the task based on a dynamic condition`,
			},
			"edges": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"condition_id": schema.StringAttribute{
							Computed: true,
						},
						"from_id": schema.StringAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"none_met": schema.BoolAttribute{
							Computed:    true,
							Description: `Indicates a default case for a decision task. Only decision task edges can have this field and the flow advances using this edge if no conditions are met.`,
						},
						"to_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the workflow is enabled or not`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"is_flow_migrated": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the workflow is migrated from workflows to flows or not`,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"org_id": schema.StringAttribute{
				Computed: true,
			},
			"phases": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"assigned_to": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"due_date": schema.StringAttribute{
							Computed: true,
						},
						"due_date_config": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"duration": schema.NumberAttribute{
									Computed: true,
								},
								"phase_id": schema.StringAttribute{
									Computed: true,
								},
								"task_id": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
								"unit": schema.StringAttribute{
									Computed: true,
								},
							},
							Description: `Set due date for the task based on a dynamic condition`,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"taxonomies": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `Taxonomy ids that are associated with this workflow and used for filtering`,
						},
					},
				},
			},
			"tasks": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"automation_task": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"assigned_to": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"automation_config": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"flow_id": schema.StringAttribute{
											Computed:    true,
											Description: `Id of the configured automation to run`,
										},
									},
									Description: `Configuration for automation execution to run`,
								},
								"description": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"value": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Longer information regarding Task`,
								},
								"due_date": schema.StringAttribute{
									Computed: true,
								},
								"due_date_config": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.NumberAttribute{
											Computed: true,
										},
										"phase_id": schema.StringAttribute{
											Computed: true,
										},
										"task_id": schema.StringAttribute{
											Computed: true,
										},
										"type": schema.StringAttribute{
											Computed: true,
										},
										"unit": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Set due date for the task based on a dynamic condition`,
								},
								"ecp": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"journey": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"journey_id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"label": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Details regarding ECP for the workflow step`,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"installer": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"journey": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"journey_id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"label": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Details regarding ECP for the workflow step`,
								},
								"journey": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"journey_id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"phase_id": schema.StringAttribute{
									Computed: true,
								},
								"requirements": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"phase_id": schema.StringAttribute{
												Computed: true,
											},
											"task_id": schema.StringAttribute{
												Computed: true,
											},
											"when": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									Description: `requirements that need to be fulfilled in order to enable the task while flow instances are running`,
								},
								"schedule": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"delayed_schedule": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"duration": schema.NumberAttribute{
													Computed: true,
												},
												"mode": schema.StringAttribute{
													Computed: true,
												},
												"unit": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"immediate_schedule": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"mode": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"relative_schedule": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"direction": schema.StringAttribute{
													Computed: true,
												},
												"duration": schema.NumberAttribute{
													Computed: true,
												},
												"mode": schema.StringAttribute{
													Computed: true,
												},
												"reference": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"attribute": schema.StringAttribute{
															Computed:    true,
															Description: `An entity attribute that identifies a date / datetime`,
														},
														"id": schema.StringAttribute{
															Computed:    true,
															Description: `The id of the entity / workflow / task, based on the origin of the schedule`,
														},
														"origin": schema.StringAttribute{
															Computed: true,
														},
														"schema": schema.StringAttribute{
															Computed:    true,
															Description: `The schema of the entity`,
														},
													},
												},
												"unit": schema.StringAttribute{
													Computed: true,
												},
											},
										},
									},
								},
								"task_type": schema.StringAttribute{
									Computed: true,
								},
								"taxonomies": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Taxonomy ids that are associated with this workflow and used for filtering`,
								},
								"trigger_mode": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"decision_task": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"assigned_to": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"conditions": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"branch_name": schema.StringAttribute{
												Computed:    true,
												Description: `The name of the branch`,
											},
											"id": schema.StringAttribute{
												Computed: true,
											},
											"logical_operator": schema.StringAttribute{
												Computed: true,
											},
											"statements": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"id": schema.StringAttribute{
															Computed: true,
														},
														"operator": schema.StringAttribute{
															Computed: true,
														},
														"source": schema.SingleNestedAttribute{
															Computed: true,
															Attributes: map[string]schema.Attribute{
																"attribute": schema.StringAttribute{
																	Computed: true,
																},
																"attribute_operation": schema.StringAttribute{
																	Computed: true,
																},
																"attribute_repeatable": schema.BoolAttribute{
																	Computed: true,
																},
																"attribute_type": schema.StringAttribute{
																	Computed: true,
																},
																"id": schema.StringAttribute{
																	Computed:    true,
																	Description: `The id of the action or trigger`,
																},
																"origin": schema.StringAttribute{
																	Computed: true,
																},
																"origin_type": schema.StringAttribute{
																	Computed: true,
																},
																"schema": schema.StringAttribute{
																	Computed: true,
																},
															},
														},
														"values": schema.ListAttribute{
															Computed:    true,
															ElementType: types.StringType,
														},
													},
												},
											},
										},
									},
								},
								"description": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"value": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Longer information regarding Task`,
								},
								"due_date": schema.StringAttribute{
									Computed: true,
								},
								"due_date_config": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.NumberAttribute{
											Computed: true,
										},
										"phase_id": schema.StringAttribute{
											Computed: true,
										},
										"task_id": schema.StringAttribute{
											Computed: true,
										},
										"type": schema.StringAttribute{
											Computed: true,
										},
										"unit": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Set due date for the task based on a dynamic condition`,
								},
								"ecp": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"journey": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"journey_id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"label": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Details regarding ECP for the workflow step`,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"installer": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"journey": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"journey_id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"label": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Details regarding ECP for the workflow step`,
								},
								"journey": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"journey_id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"phase_id": schema.StringAttribute{
									Computed: true,
								},
								"requirements": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"phase_id": schema.StringAttribute{
												Computed: true,
											},
											"task_id": schema.StringAttribute{
												Computed: true,
											},
											"when": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									Description: `requirements that need to be fulfilled in order to enable the task while flow instances are running`,
								},
								"schedule": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"delayed_schedule": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"duration": schema.NumberAttribute{
													Computed: true,
												},
												"mode": schema.StringAttribute{
													Computed: true,
												},
												"unit": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"relative_schedule": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"direction": schema.StringAttribute{
													Computed: true,
												},
												"duration": schema.NumberAttribute{
													Computed: true,
												},
												"mode": schema.StringAttribute{
													Computed: true,
												},
												"reference": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"attribute": schema.StringAttribute{
															Computed:    true,
															Description: `An entity attribute that identifies a date / datetime`,
														},
														"id": schema.StringAttribute{
															Computed:    true,
															Description: `The id of the entity / workflow / task, based on the origin of the schedule`,
														},
														"origin": schema.StringAttribute{
															Computed: true,
														},
														"schema": schema.StringAttribute{
															Computed:    true,
															Description: `The schema of the entity`,
														},
													},
												},
												"unit": schema.StringAttribute{
													Computed: true,
												},
											},
										},
									},
								},
								"task_type": schema.StringAttribute{
									Computed: true,
								},
								"taxonomies": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Taxonomy ids that are associated with this workflow and used for filtering`,
								},
							},
						},
						"task_base": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"assigned_to": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"description": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"value": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Longer information regarding Task`,
								},
								"due_date": schema.StringAttribute{
									Computed: true,
								},
								"due_date_config": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"duration": schema.NumberAttribute{
											Computed: true,
										},
										"phase_id": schema.StringAttribute{
											Computed: true,
										},
										"task_id": schema.StringAttribute{
											Computed: true,
										},
										"type": schema.StringAttribute{
											Computed: true,
										},
										"unit": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Set due date for the task based on a dynamic condition`,
								},
								"ecp": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"journey": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"journey_id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"label": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Details regarding ECP for the workflow step`,
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"installer": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Computed: true,
										},
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
										"journey": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Computed: true,
												},
												"journey_id": schema.StringAttribute{
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										"label": schema.StringAttribute{
											Computed: true,
										},
									},
									Description: `Details regarding ECP for the workflow step`,
								},
								"journey": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"journey_id": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"phase_id": schema.StringAttribute{
									Computed: true,
								},
								"requirements": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"phase_id": schema.StringAttribute{
												Computed: true,
											},
											"task_id": schema.StringAttribute{
												Computed: true,
											},
											"when": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									Description: `requirements that need to be fulfilled in order to enable the task while flow instances are running`,
								},
								"task_type": schema.StringAttribute{
									Computed: true,
								},
								"taxonomies": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `Taxonomy ids that are associated with this workflow and used for filtering`,
								},
							},
						},
					},
				},
			},
			"taxonomies": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `Taxonomy ids that are associated with this workflow and used for filtering`,
			},
			"trigger": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"automation_trigger": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"automation_id": schema.StringAttribute{
								Computed:    true,
								Description: `Id of the automation config that triggers this workflow`,
							},
							"id": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"journey_submission_trigger": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"automation_id": schema.StringAttribute{
								Computed: true,
							},
							"id": schema.StringAttribute{
								Computed: true,
							},
							"journey_id": schema.StringAttribute{
								Computed:    true,
								Description: `ID of the journey that will trigger this flow`,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"manual_trigger": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"entity_schema": schema.StringAttribute{
								Computed: true,
							},
							"id": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"update_entity_attributes": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Computed: true,
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
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `ISO String Date & Time`,
			},
		},
	}
}

func (r *FlowTemplateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *FlowTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *FlowTemplateDataSourceModel
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

	var flowID string
	flowID = data.ID.ValueString()

	request := operations.GetFlowTemplateRequest{
		FlowID: flowID,
	}
	res, err := r.client.FlowsV2.GetFlowTemplate(ctx, request)
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
	if !(res.FlowTemplate != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedFlowTemplate(res.FlowTemplate)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
