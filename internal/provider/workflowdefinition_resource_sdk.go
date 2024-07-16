// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-workflow/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-workflow/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
)

func (r *WorkflowDefinitionResourceModel) ToSharedWorkflowDefinition() *shared.WorkflowDefinition {
	var assignedTo []string = []string{}
	for _, assignedToItem := range r.AssignedTo {
		assignedTo = append(assignedTo, assignedToItem.ValueString())
	}
	var closingReasons []shared.ClosingReasonID = []shared.ClosingReasonID{}
	for _, closingReasonsItem := range r.ClosingReasons {
		id := closingReasonsItem.ID.ValueString()
		closingReasons = append(closingReasons, shared.ClosingReasonID{
			ID: id,
		})
	}
	creationTime := new(string)
	if !r.CreationTime.IsUnknown() && !r.CreationTime.IsNull() {
		*creationTime = r.CreationTime.ValueString()
	} else {
		creationTime = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	dueDate := new(string)
	if !r.DueDate.IsUnknown() && !r.DueDate.IsNull() {
		*dueDate = r.DueDate.ValueString()
	} else {
		dueDate = nil
	}
	var dynamicDueDate *shared.DynamicDueDate
	if r.DynamicDueDate != nil {
		actionTypeCondition := new(shared.ActionTypeCondition)
		if !r.DynamicDueDate.ActionTypeCondition.IsUnknown() && !r.DynamicDueDate.ActionTypeCondition.IsNull() {
			*actionTypeCondition = shared.ActionTypeCondition(r.DynamicDueDate.ActionTypeCondition.ValueString())
		} else {
			actionTypeCondition = nil
		}
		numberOfUnits := new(float64)
		if !r.DynamicDueDate.NumberOfUnits.IsUnknown() && !r.DynamicDueDate.NumberOfUnits.IsNull() {
			*numberOfUnits, _ = r.DynamicDueDate.NumberOfUnits.ValueBigFloat().Float64()
		} else {
			numberOfUnits = nil
		}
		stepID := new(string)
		if !r.DynamicDueDate.StepID.IsUnknown() && !r.DynamicDueDate.StepID.IsNull() {
			*stepID = r.DynamicDueDate.StepID.ValueString()
		} else {
			stepID = nil
		}
		timePeriod := new(shared.TimePeriod)
		if !r.DynamicDueDate.TimePeriod.IsUnknown() && !r.DynamicDueDate.TimePeriod.IsNull() {
			*timePeriod = shared.TimePeriod(r.DynamicDueDate.TimePeriod.ValueString())
		} else {
			timePeriod = nil
		}
		dynamicDueDate = &shared.DynamicDueDate{
			ActionTypeCondition: actionTypeCondition,
			NumberOfUnits:       numberOfUnits,
			StepID:              stepID,
			TimePeriod:          timePeriod,
		}
	}
	enableECPWorkflow := new(bool)
	if !r.EnableECPWorkflow.IsUnknown() && !r.EnableECPWorkflow.IsNull() {
		*enableECPWorkflow = r.EnableECPWorkflow.ValueBool()
	} else {
		enableECPWorkflow = nil
	}
	var flow []shared.Flow = []shared.Flow{}
	for _, flowItem := range r.Flow {
		if flowItem.Section != nil {
			id1 := new(string)
			if !flowItem.Section.ID.IsUnknown() && !flowItem.Section.ID.IsNull() {
				*id1 = flowItem.Section.ID.ValueString()
			} else {
				id1 = nil
			}
			name := flowItem.Section.Name.ValueString()
			order, _ := flowItem.Section.Order.ValueBigFloat().Float64()
			var steps []shared.Step = []shared.Step{}
			for _, stepsItem := range flowItem.Section.Steps {
				var assignedTo1 []string = []string{}
				for _, assignedToItem1 := range stepsItem.AssignedTo {
					assignedTo1 = append(assignedTo1, assignedToItem1.ValueString())
				}
				var automationConfig *shared.AutomationConfig
				if stepsItem.AutomationConfig != nil {
					flowID := stepsItem.AutomationConfig.FlowID.ValueString()
					automationConfig = &shared.AutomationConfig{
						FlowID: flowID,
					}
				}
				var description1 *shared.StepDescription
				if stepsItem.Description != nil {
					enabled := new(bool)
					if !stepsItem.Description.Enabled.IsUnknown() && !stepsItem.Description.Enabled.IsNull() {
						*enabled = stepsItem.Description.Enabled.ValueBool()
					} else {
						enabled = nil
					}
					value := new(string)
					if !stepsItem.Description.Value.IsUnknown() && !stepsItem.Description.Value.IsNull() {
						*value = stepsItem.Description.Value.ValueString()
					} else {
						value = nil
					}
					description1 = &shared.StepDescription{
						Enabled: enabled,
						Value:   value,
					}
				}
				dueDate1 := new(string)
				if !stepsItem.DueDate.IsUnknown() && !stepsItem.DueDate.IsNull() {
					*dueDate1 = stepsItem.DueDate.ValueString()
				} else {
					dueDate1 = nil
				}
				var dynamicDueDate1 *shared.DynamicDueDate
				if stepsItem.DynamicDueDate != nil {
					actionTypeCondition1 := new(shared.ActionTypeCondition)
					if !stepsItem.DynamicDueDate.ActionTypeCondition.IsUnknown() && !stepsItem.DynamicDueDate.ActionTypeCondition.IsNull() {
						*actionTypeCondition1 = shared.ActionTypeCondition(stepsItem.DynamicDueDate.ActionTypeCondition.ValueString())
					} else {
						actionTypeCondition1 = nil
					}
					numberOfUnits1 := new(float64)
					if !stepsItem.DynamicDueDate.NumberOfUnits.IsUnknown() && !stepsItem.DynamicDueDate.NumberOfUnits.IsNull() {
						*numberOfUnits1, _ = stepsItem.DynamicDueDate.NumberOfUnits.ValueBigFloat().Float64()
					} else {
						numberOfUnits1 = nil
					}
					stepId1 := new(string)
					if !stepsItem.DynamicDueDate.StepID.IsUnknown() && !stepsItem.DynamicDueDate.StepID.IsNull() {
						*stepId1 = stepsItem.DynamicDueDate.StepID.ValueString()
					} else {
						stepId1 = nil
					}
					timePeriod1 := new(shared.TimePeriod)
					if !stepsItem.DynamicDueDate.TimePeriod.IsUnknown() && !stepsItem.DynamicDueDate.TimePeriod.IsNull() {
						*timePeriod1 = shared.TimePeriod(stepsItem.DynamicDueDate.TimePeriod.ValueString())
					} else {
						timePeriod1 = nil
					}
					dynamicDueDate1 = &shared.DynamicDueDate{
						ActionTypeCondition: actionTypeCondition1,
						NumberOfUnits:       numberOfUnits1,
						StepID:              stepId1,
						TimePeriod:          timePeriod1,
					}
				}
				var ecp *shared.ECPDetails
				if stepsItem.Ecp != nil {
					description2 := new(string)
					if !stepsItem.Ecp.Description.IsUnknown() && !stepsItem.Ecp.Description.IsNull() {
						*description2 = stepsItem.Ecp.Description.ValueString()
					} else {
						description2 = nil
					}
					enabled1 := new(bool)
					if !stepsItem.Ecp.Enabled.IsUnknown() && !stepsItem.Ecp.Enabled.IsNull() {
						*enabled1 = stepsItem.Ecp.Enabled.ValueBool()
					} else {
						enabled1 = nil
					}
					var journey *shared.StepJourney
					if stepsItem.Ecp.Journey != nil {
						id2 := new(string)
						if !stepsItem.Ecp.Journey.ID.IsUnknown() && !stepsItem.Ecp.Journey.ID.IsNull() {
							*id2 = stepsItem.Ecp.Journey.ID.ValueString()
						} else {
							id2 = nil
						}
						journeyID := new(string)
						if !stepsItem.Ecp.Journey.JourneyID.IsUnknown() && !stepsItem.Ecp.Journey.JourneyID.IsNull() {
							*journeyID = stepsItem.Ecp.Journey.JourneyID.ValueString()
						} else {
							journeyID = nil
						}
						name1 := new(string)
						if !stepsItem.Ecp.Journey.Name.IsUnknown() && !stepsItem.Ecp.Journey.Name.IsNull() {
							*name1 = stepsItem.Ecp.Journey.Name.ValueString()
						} else {
							name1 = nil
						}
						journey = &shared.StepJourney{
							ID:        id2,
							JourneyID: journeyID,
							Name:      name1,
						}
					}
					label := new(string)
					if !stepsItem.Ecp.Label.IsUnknown() && !stepsItem.Ecp.Label.IsNull() {
						*label = stepsItem.Ecp.Label.ValueString()
					} else {
						label = nil
					}
					ecp = &shared.ECPDetails{
						Description: description2,
						Enabled:     enabled1,
						Journey:     journey,
						Label:       label,
					}
				}
				executionType := new(shared.StepType)
				if !stepsItem.ExecutionType.IsUnknown() && !stepsItem.ExecutionType.IsNull() {
					*executionType = shared.StepType(stepsItem.ExecutionType.ValueString())
				} else {
					executionType = nil
				}
				id3 := new(string)
				if !stepsItem.ID.IsUnknown() && !stepsItem.ID.IsNull() {
					*id3 = stepsItem.ID.ValueString()
				} else {
					id3 = nil
				}
				var installer *shared.ECPDetails
				if stepsItem.Installer != nil {
					description3 := new(string)
					if !stepsItem.Installer.Description.IsUnknown() && !stepsItem.Installer.Description.IsNull() {
						*description3 = stepsItem.Installer.Description.ValueString()
					} else {
						description3 = nil
					}
					enabled2 := new(bool)
					if !stepsItem.Installer.Enabled.IsUnknown() && !stepsItem.Installer.Enabled.IsNull() {
						*enabled2 = stepsItem.Installer.Enabled.ValueBool()
					} else {
						enabled2 = nil
					}
					var journey1 *shared.StepJourney
					if stepsItem.Installer.Journey != nil {
						id4 := new(string)
						if !stepsItem.Installer.Journey.ID.IsUnknown() && !stepsItem.Installer.Journey.ID.IsNull() {
							*id4 = stepsItem.Installer.Journey.ID.ValueString()
						} else {
							id4 = nil
						}
						journeyId1 := new(string)
						if !stepsItem.Installer.Journey.JourneyID.IsUnknown() && !stepsItem.Installer.Journey.JourneyID.IsNull() {
							*journeyId1 = stepsItem.Installer.Journey.JourneyID.ValueString()
						} else {
							journeyId1 = nil
						}
						name2 := new(string)
						if !stepsItem.Installer.Journey.Name.IsUnknown() && !stepsItem.Installer.Journey.Name.IsNull() {
							*name2 = stepsItem.Installer.Journey.Name.ValueString()
						} else {
							name2 = nil
						}
						journey1 = &shared.StepJourney{
							ID:        id4,
							JourneyID: journeyId1,
							Name:      name2,
						}
					}
					label1 := new(string)
					if !stepsItem.Installer.Label.IsUnknown() && !stepsItem.Installer.Label.IsNull() {
						*label1 = stepsItem.Installer.Label.ValueString()
					} else {
						label1 = nil
					}
					installer = &shared.ECPDetails{
						Description: description3,
						Enabled:     enabled2,
						Journey:     journey1,
						Label:       label1,
					}
				}
				var journey2 *shared.StepJourney
				if stepsItem.Journey != nil {
					id5 := new(string)
					if !stepsItem.Journey.ID.IsUnknown() && !stepsItem.Journey.ID.IsNull() {
						*id5 = stepsItem.Journey.ID.ValueString()
					} else {
						id5 = nil
					}
					journeyId2 := new(string)
					if !stepsItem.Journey.JourneyID.IsUnknown() && !stepsItem.Journey.JourneyID.IsNull() {
						*journeyId2 = stepsItem.Journey.JourneyID.ValueString()
					} else {
						journeyId2 = nil
					}
					name3 := new(string)
					if !stepsItem.Journey.Name.IsUnknown() && !stepsItem.Journey.Name.IsNull() {
						*name3 = stepsItem.Journey.Name.ValueString()
					} else {
						name3 = nil
					}
					journey2 = &shared.StepJourney{
						ID:        id5,
						JourneyID: journeyId2,
						Name:      name3,
					}
				}
				name4 := stepsItem.Name.ValueString()
				order1, _ := stepsItem.Order.ValueBigFloat().Float64()
				var requirements []shared.StepRequirement = []shared.StepRequirement{}
				for _, requirementsItem := range stepsItem.Requirements {
					condition := shared.Condition(requirementsItem.Condition.ValueString())
					definitionID := requirementsItem.DefinitionID.ValueString()
					typeVar := shared.ItemType(requirementsItem.Type.ValueString())
					requirements = append(requirements, shared.StepRequirement{
						Condition:    condition,
						DefinitionID: definitionID,
						Type:         typeVar,
					})
				}
				type1 := shared.ItemType(stepsItem.Type.ValueString())
				var userIds []float64 = []float64{}
				for _, userIdsItem := range stepsItem.UserIds {
					userIdsTmp, _ := userIdsItem.ValueBigFloat().Float64()
					userIds = append(userIds, userIdsTmp)
				}
				steps = append(steps, shared.Step{
					AssignedTo:       assignedTo1,
					AutomationConfig: automationConfig,
					Description:      description1,
					DueDate:          dueDate1,
					DynamicDueDate:   dynamicDueDate1,
					Ecp:              ecp,
					ExecutionType:    executionType,
					ID:               id3,
					Installer:        installer,
					Journey:          journey2,
					Name:             name4,
					Order:            order1,
					Requirements:     requirements,
					Type:             type1,
					UserIds:          userIds,
				})
			}
			typeVar1 := shared.ItemType(flowItem.Section.Type.ValueString())
			section := shared.Section{
				ID:    id1,
				Name:  name,
				Order: order,
				Steps: steps,
				Type:  typeVar1,
			}
			flow = append(flow, shared.Flow{
				Section: &section,
			})
		}
		if flowItem.Step != nil {
			var assignedTo2 []string = []string{}
			for _, assignedToItem2 := range flowItem.Step.AssignedTo {
				assignedTo2 = append(assignedTo2, assignedToItem2.ValueString())
			}
			var automationConfig1 *shared.AutomationConfig
			if flowItem.Step.AutomationConfig != nil {
				flowId1 := flowItem.Step.AutomationConfig.FlowID.ValueString()
				automationConfig1 = &shared.AutomationConfig{
					FlowID: flowId1,
				}
			}
			var description4 *shared.StepDescription
			if flowItem.Step.Description != nil {
				enabled3 := new(bool)
				if !flowItem.Step.Description.Enabled.IsUnknown() && !flowItem.Step.Description.Enabled.IsNull() {
					*enabled3 = flowItem.Step.Description.Enabled.ValueBool()
				} else {
					enabled3 = nil
				}
				value1 := new(string)
				if !flowItem.Step.Description.Value.IsUnknown() && !flowItem.Step.Description.Value.IsNull() {
					*value1 = flowItem.Step.Description.Value.ValueString()
				} else {
					value1 = nil
				}
				description4 = &shared.StepDescription{
					Enabled: enabled3,
					Value:   value1,
				}
			}
			dueDate2 := new(string)
			if !flowItem.Step.DueDate.IsUnknown() && !flowItem.Step.DueDate.IsNull() {
				*dueDate2 = flowItem.Step.DueDate.ValueString()
			} else {
				dueDate2 = nil
			}
			var dynamicDueDate2 *shared.DynamicDueDate
			if flowItem.Step.DynamicDueDate != nil {
				actionTypeCondition2 := new(shared.ActionTypeCondition)
				if !flowItem.Step.DynamicDueDate.ActionTypeCondition.IsUnknown() && !flowItem.Step.DynamicDueDate.ActionTypeCondition.IsNull() {
					*actionTypeCondition2 = shared.ActionTypeCondition(flowItem.Step.DynamicDueDate.ActionTypeCondition.ValueString())
				} else {
					actionTypeCondition2 = nil
				}
				numberOfUnits2 := new(float64)
				if !flowItem.Step.DynamicDueDate.NumberOfUnits.IsUnknown() && !flowItem.Step.DynamicDueDate.NumberOfUnits.IsNull() {
					*numberOfUnits2, _ = flowItem.Step.DynamicDueDate.NumberOfUnits.ValueBigFloat().Float64()
				} else {
					numberOfUnits2 = nil
				}
				stepId2 := new(string)
				if !flowItem.Step.DynamicDueDate.StepID.IsUnknown() && !flowItem.Step.DynamicDueDate.StepID.IsNull() {
					*stepId2 = flowItem.Step.DynamicDueDate.StepID.ValueString()
				} else {
					stepId2 = nil
				}
				timePeriod2 := new(shared.TimePeriod)
				if !flowItem.Step.DynamicDueDate.TimePeriod.IsUnknown() && !flowItem.Step.DynamicDueDate.TimePeriod.IsNull() {
					*timePeriod2 = shared.TimePeriod(flowItem.Step.DynamicDueDate.TimePeriod.ValueString())
				} else {
					timePeriod2 = nil
				}
				dynamicDueDate2 = &shared.DynamicDueDate{
					ActionTypeCondition: actionTypeCondition2,
					NumberOfUnits:       numberOfUnits2,
					StepID:              stepId2,
					TimePeriod:          timePeriod2,
				}
			}
			var ecp1 *shared.ECPDetails
			if flowItem.Step.Ecp != nil {
				description5 := new(string)
				if !flowItem.Step.Ecp.Description.IsUnknown() && !flowItem.Step.Ecp.Description.IsNull() {
					*description5 = flowItem.Step.Ecp.Description.ValueString()
				} else {
					description5 = nil
				}
				enabled4 := new(bool)
				if !flowItem.Step.Ecp.Enabled.IsUnknown() && !flowItem.Step.Ecp.Enabled.IsNull() {
					*enabled4 = flowItem.Step.Ecp.Enabled.ValueBool()
				} else {
					enabled4 = nil
				}
				var journey3 *shared.StepJourney
				if flowItem.Step.Ecp.Journey != nil {
					id6 := new(string)
					if !flowItem.Step.Ecp.Journey.ID.IsUnknown() && !flowItem.Step.Ecp.Journey.ID.IsNull() {
						*id6 = flowItem.Step.Ecp.Journey.ID.ValueString()
					} else {
						id6 = nil
					}
					journeyId3 := new(string)
					if !flowItem.Step.Ecp.Journey.JourneyID.IsUnknown() && !flowItem.Step.Ecp.Journey.JourneyID.IsNull() {
						*journeyId3 = flowItem.Step.Ecp.Journey.JourneyID.ValueString()
					} else {
						journeyId3 = nil
					}
					name5 := new(string)
					if !flowItem.Step.Ecp.Journey.Name.IsUnknown() && !flowItem.Step.Ecp.Journey.Name.IsNull() {
						*name5 = flowItem.Step.Ecp.Journey.Name.ValueString()
					} else {
						name5 = nil
					}
					journey3 = &shared.StepJourney{
						ID:        id6,
						JourneyID: journeyId3,
						Name:      name5,
					}
				}
				label2 := new(string)
				if !flowItem.Step.Ecp.Label.IsUnknown() && !flowItem.Step.Ecp.Label.IsNull() {
					*label2 = flowItem.Step.Ecp.Label.ValueString()
				} else {
					label2 = nil
				}
				ecp1 = &shared.ECPDetails{
					Description: description5,
					Enabled:     enabled4,
					Journey:     journey3,
					Label:       label2,
				}
			}
			executionType1 := new(shared.StepType)
			if !flowItem.Step.ExecutionType.IsUnknown() && !flowItem.Step.ExecutionType.IsNull() {
				*executionType1 = shared.StepType(flowItem.Step.ExecutionType.ValueString())
			} else {
				executionType1 = nil
			}
			id7 := new(string)
			if !flowItem.Step.ID.IsUnknown() && !flowItem.Step.ID.IsNull() {
				*id7 = flowItem.Step.ID.ValueString()
			} else {
				id7 = nil
			}
			var installer1 *shared.ECPDetails
			if flowItem.Step.Installer != nil {
				description6 := new(string)
				if !flowItem.Step.Installer.Description.IsUnknown() && !flowItem.Step.Installer.Description.IsNull() {
					*description6 = flowItem.Step.Installer.Description.ValueString()
				} else {
					description6 = nil
				}
				enabled5 := new(bool)
				if !flowItem.Step.Installer.Enabled.IsUnknown() && !flowItem.Step.Installer.Enabled.IsNull() {
					*enabled5 = flowItem.Step.Installer.Enabled.ValueBool()
				} else {
					enabled5 = nil
				}
				var journey4 *shared.StepJourney
				if flowItem.Step.Installer.Journey != nil {
					id8 := new(string)
					if !flowItem.Step.Installer.Journey.ID.IsUnknown() && !flowItem.Step.Installer.Journey.ID.IsNull() {
						*id8 = flowItem.Step.Installer.Journey.ID.ValueString()
					} else {
						id8 = nil
					}
					journeyId4 := new(string)
					if !flowItem.Step.Installer.Journey.JourneyID.IsUnknown() && !flowItem.Step.Installer.Journey.JourneyID.IsNull() {
						*journeyId4 = flowItem.Step.Installer.Journey.JourneyID.ValueString()
					} else {
						journeyId4 = nil
					}
					name6 := new(string)
					if !flowItem.Step.Installer.Journey.Name.IsUnknown() && !flowItem.Step.Installer.Journey.Name.IsNull() {
						*name6 = flowItem.Step.Installer.Journey.Name.ValueString()
					} else {
						name6 = nil
					}
					journey4 = &shared.StepJourney{
						ID:        id8,
						JourneyID: journeyId4,
						Name:      name6,
					}
				}
				label3 := new(string)
				if !flowItem.Step.Installer.Label.IsUnknown() && !flowItem.Step.Installer.Label.IsNull() {
					*label3 = flowItem.Step.Installer.Label.ValueString()
				} else {
					label3 = nil
				}
				installer1 = &shared.ECPDetails{
					Description: description6,
					Enabled:     enabled5,
					Journey:     journey4,
					Label:       label3,
				}
			}
			var journey5 *shared.StepJourney
			if flowItem.Step.Journey != nil {
				id9 := new(string)
				if !flowItem.Step.Journey.ID.IsUnknown() && !flowItem.Step.Journey.ID.IsNull() {
					*id9 = flowItem.Step.Journey.ID.ValueString()
				} else {
					id9 = nil
				}
				journeyId5 := new(string)
				if !flowItem.Step.Journey.JourneyID.IsUnknown() && !flowItem.Step.Journey.JourneyID.IsNull() {
					*journeyId5 = flowItem.Step.Journey.JourneyID.ValueString()
				} else {
					journeyId5 = nil
				}
				name7 := new(string)
				if !flowItem.Step.Journey.Name.IsUnknown() && !flowItem.Step.Journey.Name.IsNull() {
					*name7 = flowItem.Step.Journey.Name.ValueString()
				} else {
					name7 = nil
				}
				journey5 = &shared.StepJourney{
					ID:        id9,
					JourneyID: journeyId5,
					Name:      name7,
				}
			}
			name8 := flowItem.Step.Name.ValueString()
			order2, _ := flowItem.Step.Order.ValueBigFloat().Float64()
			var requirements1 []shared.StepRequirement = []shared.StepRequirement{}
			for _, requirementsItem1 := range flowItem.Step.Requirements {
				condition1 := shared.Condition(requirementsItem1.Condition.ValueString())
				definitionId1 := requirementsItem1.DefinitionID.ValueString()
				type2 := shared.ItemType(requirementsItem1.Type.ValueString())
				requirements1 = append(requirements1, shared.StepRequirement{
					Condition:    condition1,
					DefinitionID: definitionId1,
					Type:         type2,
				})
			}
			typeVar2 := shared.ItemType(flowItem.Step.Type.ValueString())
			var userIds1 []float64 = []float64{}
			for _, userIdsItem1 := range flowItem.Step.UserIds {
				userIdsTmp1, _ := userIdsItem1.ValueBigFloat().Float64()
				userIds1 = append(userIds1, userIdsTmp1)
			}
			step := shared.Step{
				AssignedTo:       assignedTo2,
				AutomationConfig: automationConfig1,
				Description:      description4,
				DueDate:          dueDate2,
				DynamicDueDate:   dynamicDueDate2,
				Ecp:              ecp1,
				ExecutionType:    executionType1,
				ID:               id7,
				Installer:        installer1,
				Journey:          journey5,
				Name:             name8,
				Order:            order2,
				Requirements:     requirements1,
				Type:             typeVar2,
				UserIds:          userIds1,
			}
			flow = append(flow, shared.Flow{
				Step: &step,
			})
		}
	}
	id10 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id10 = r.ID.ValueString()
	} else {
		id10 = nil
	}
	lastUpdateTime := new(string)
	if !r.LastUpdateTime.IsUnknown() && !r.LastUpdateTime.IsNull() {
		*lastUpdateTime = r.LastUpdateTime.ValueString()
	} else {
		lastUpdateTime = nil
	}
	name9 := r.Name.ValueString()
	var updateEntityAttributes []shared.UpdateEntityAttributes = []shared.UpdateEntityAttributes{}
	for _, updateEntityAttributesItem := range r.UpdateEntityAttributes {
		source := shared.Source(updateEntityAttributesItem.Source.ValueString())
		entityAttribute := updateEntityAttributesItem.Target.EntityAttribute.ValueString()
		entitySchema := updateEntityAttributesItem.Target.EntitySchema.ValueString()
		target := shared.Target{
			EntityAttribute: entityAttribute,
			EntitySchema:    entitySchema,
		}
		updateEntityAttributes = append(updateEntityAttributes, shared.UpdateEntityAttributes{
			Source: source,
			Target: target,
		})
	}
	var userIds2 []float64 = []float64{}
	for _, userIdsItem2 := range r.UserIds {
		userIdsTmp2, _ := userIdsItem2.ValueBigFloat().Float64()
		userIds2 = append(userIds2, userIdsTmp2)
	}
	out := shared.WorkflowDefinition{
		AssignedTo:             assignedTo,
		ClosingReasons:         closingReasons,
		CreationTime:           creationTime,
		Description:            description,
		DueDate:                dueDate,
		DynamicDueDate:         dynamicDueDate,
		EnableECPWorkflow:      enableECPWorkflow,
		Flow:                   flow,
		ID:                     id10,
		LastUpdateTime:         lastUpdateTime,
		Name:                   name9,
		UpdateEntityAttributes: updateEntityAttributes,
		UserIds:                userIds2,
	}
	return &out
}

func (r *WorkflowDefinitionResourceModel) RefreshFromSharedWorkflowDefinition(resp *shared.WorkflowDefinition) {
	if resp != nil {
		r.AssignedTo = []types.String{}
		for _, v := range resp.AssignedTo {
			r.AssignedTo = append(r.AssignedTo, types.StringValue(v))
		}
		r.ClosingReasons = []tfTypes.ClosingReasonID{}
		if len(r.ClosingReasons) > len(resp.ClosingReasons) {
			r.ClosingReasons = r.ClosingReasons[:len(resp.ClosingReasons)]
		}
		for closingReasonsCount, closingReasonsItem := range resp.ClosingReasons {
			var closingReasons1 tfTypes.ClosingReasonID
			closingReasons1.ID = types.StringValue(closingReasonsItem.ID)
			if closingReasonsCount+1 > len(r.ClosingReasons) {
				r.ClosingReasons = append(r.ClosingReasons, closingReasons1)
			} else {
				r.ClosingReasons[closingReasonsCount].ID = closingReasons1.ID
			}
		}
		r.CreationTime = types.StringPointerValue(resp.CreationTime)
		r.Description = types.StringPointerValue(resp.Description)
		r.DueDate = types.StringPointerValue(resp.DueDate)
		if resp.DynamicDueDate == nil {
			r.DynamicDueDate = nil
		} else {
			r.DynamicDueDate = &tfTypes.DynamicDueDate{}
			if resp.DynamicDueDate.ActionTypeCondition != nil {
				r.DynamicDueDate.ActionTypeCondition = types.StringValue(string(*resp.DynamicDueDate.ActionTypeCondition))
			} else {
				r.DynamicDueDate.ActionTypeCondition = types.StringNull()
			}
			if resp.DynamicDueDate.NumberOfUnits != nil {
				r.DynamicDueDate.NumberOfUnits = types.NumberValue(big.NewFloat(float64(*resp.DynamicDueDate.NumberOfUnits)))
			} else {
				r.DynamicDueDate.NumberOfUnits = types.NumberNull()
			}
			r.DynamicDueDate.StepID = types.StringPointerValue(resp.DynamicDueDate.StepID)
			if resp.DynamicDueDate.TimePeriod != nil {
				r.DynamicDueDate.TimePeriod = types.StringValue(string(*resp.DynamicDueDate.TimePeriod))
			} else {
				r.DynamicDueDate.TimePeriod = types.StringNull()
			}
		}
		r.EnableECPWorkflow = types.BoolPointerValue(resp.EnableECPWorkflow)
		r.Flow = []tfTypes.Flow{}
		if len(r.Flow) > len(resp.Flow) {
			r.Flow = r.Flow[:len(resp.Flow)]
		}
		for flowCount, flowItem := range resp.Flow {
			var flow1 tfTypes.Flow
			if flowItem.Section != nil {
				flow1.Section = &tfTypes.Section{}
				flow1.Section.ID = types.StringPointerValue(flowItem.Section.ID)
				flow1.Section.Name = types.StringValue(flowItem.Section.Name)
				flow1.Section.Order = types.NumberValue(big.NewFloat(float64(flowItem.Section.Order)))
				flow1.Section.Steps = []tfTypes.Step{}
				for stepsCount, stepsItem := range flowItem.Section.Steps {
					var steps1 tfTypes.Step
					steps1.AssignedTo = []types.String{}
					for _, v := range stepsItem.AssignedTo {
						steps1.AssignedTo = append(steps1.AssignedTo, types.StringValue(v))
					}
					if stepsItem.AutomationConfig == nil {
						steps1.AutomationConfig = nil
					} else {
						steps1.AutomationConfig = &tfTypes.AutomationConfig{}
						steps1.AutomationConfig.FlowID = types.StringValue(stepsItem.AutomationConfig.FlowID)
					}
					if stepsItem.Description == nil {
						steps1.Description = nil
					} else {
						steps1.Description = &tfTypes.StepDescription{}
						steps1.Description.Enabled = types.BoolPointerValue(stepsItem.Description.Enabled)
						steps1.Description.Value = types.StringPointerValue(stepsItem.Description.Value)
					}
					steps1.DueDate = types.StringPointerValue(stepsItem.DueDate)
					if stepsItem.DynamicDueDate == nil {
						steps1.DynamicDueDate = nil
					} else {
						steps1.DynamicDueDate = &tfTypes.DynamicDueDate{}
						if stepsItem.DynamicDueDate.ActionTypeCondition != nil {
							steps1.DynamicDueDate.ActionTypeCondition = types.StringValue(string(*stepsItem.DynamicDueDate.ActionTypeCondition))
						} else {
							steps1.DynamicDueDate.ActionTypeCondition = types.StringNull()
						}
						if stepsItem.DynamicDueDate.NumberOfUnits != nil {
							steps1.DynamicDueDate.NumberOfUnits = types.NumberValue(big.NewFloat(float64(*stepsItem.DynamicDueDate.NumberOfUnits)))
						} else {
							steps1.DynamicDueDate.NumberOfUnits = types.NumberNull()
						}
						steps1.DynamicDueDate.StepID = types.StringPointerValue(stepsItem.DynamicDueDate.StepID)
						if stepsItem.DynamicDueDate.TimePeriod != nil {
							steps1.DynamicDueDate.TimePeriod = types.StringValue(string(*stepsItem.DynamicDueDate.TimePeriod))
						} else {
							steps1.DynamicDueDate.TimePeriod = types.StringNull()
						}
					}
					if stepsItem.Ecp == nil {
						steps1.Ecp = nil
					} else {
						steps1.Ecp = &tfTypes.ECPDetails{}
						steps1.Ecp.Description = types.StringPointerValue(stepsItem.Ecp.Description)
						steps1.Ecp.Enabled = types.BoolPointerValue(stepsItem.Ecp.Enabled)
						if stepsItem.Ecp.Journey == nil {
							steps1.Ecp.Journey = nil
						} else {
							steps1.Ecp.Journey = &tfTypes.StepJourney{}
							steps1.Ecp.Journey.ID = types.StringPointerValue(stepsItem.Ecp.Journey.ID)
							steps1.Ecp.Journey.JourneyID = types.StringPointerValue(stepsItem.Ecp.Journey.JourneyID)
							steps1.Ecp.Journey.Name = types.StringPointerValue(stepsItem.Ecp.Journey.Name)
						}
						steps1.Ecp.Label = types.StringPointerValue(stepsItem.Ecp.Label)
					}
					if stepsItem.ExecutionType != nil {
						steps1.ExecutionType = types.StringValue(string(*stepsItem.ExecutionType))
					} else {
						steps1.ExecutionType = types.StringNull()
					}
					steps1.ID = types.StringPointerValue(stepsItem.ID)
					if stepsItem.Installer == nil {
						steps1.Installer = nil
					} else {
						steps1.Installer = &tfTypes.ECPDetails{}
						steps1.Installer.Description = types.StringPointerValue(stepsItem.Installer.Description)
						steps1.Installer.Enabled = types.BoolPointerValue(stepsItem.Installer.Enabled)
						if stepsItem.Installer.Journey == nil {
							steps1.Installer.Journey = nil
						} else {
							steps1.Installer.Journey = &tfTypes.StepJourney{}
							steps1.Installer.Journey.ID = types.StringPointerValue(stepsItem.Installer.Journey.ID)
							steps1.Installer.Journey.JourneyID = types.StringPointerValue(stepsItem.Installer.Journey.JourneyID)
							steps1.Installer.Journey.Name = types.StringPointerValue(stepsItem.Installer.Journey.Name)
						}
						steps1.Installer.Label = types.StringPointerValue(stepsItem.Installer.Label)
					}
					if stepsItem.Journey == nil {
						steps1.Journey = nil
					} else {
						steps1.Journey = &tfTypes.StepJourney{}
						steps1.Journey.ID = types.StringPointerValue(stepsItem.Journey.ID)
						steps1.Journey.JourneyID = types.StringPointerValue(stepsItem.Journey.JourneyID)
						steps1.Journey.Name = types.StringPointerValue(stepsItem.Journey.Name)
					}
					steps1.Name = types.StringValue(stepsItem.Name)
					steps1.Order = types.NumberValue(big.NewFloat(float64(stepsItem.Order)))
					steps1.Requirements = []tfTypes.StepRequirement{}
					for requirementsCount, requirementsItem := range stepsItem.Requirements {
						var requirements1 tfTypes.StepRequirement
						requirements1.Condition = types.StringValue(string(requirementsItem.Condition))
						requirements1.DefinitionID = types.StringValue(requirementsItem.DefinitionID)
						requirements1.Type = types.StringValue(string(requirementsItem.Type))
						if requirementsCount+1 > len(steps1.Requirements) {
							steps1.Requirements = append(steps1.Requirements, requirements1)
						} else {
							steps1.Requirements[requirementsCount].Condition = requirements1.Condition
							steps1.Requirements[requirementsCount].DefinitionID = requirements1.DefinitionID
							steps1.Requirements[requirementsCount].Type = requirements1.Type
						}
					}
					steps1.Type = types.StringValue(string(stepsItem.Type))
					steps1.UserIds = []types.Number{}
					for _, v := range stepsItem.UserIds {
						steps1.UserIds = append(steps1.UserIds, types.NumberValue(big.NewFloat(float64(v))))
					}
					if stepsCount+1 > len(flow1.Section.Steps) {
						flow1.Section.Steps = append(flow1.Section.Steps, steps1)
					} else {
						flow1.Section.Steps[stepsCount].AssignedTo = steps1.AssignedTo
						flow1.Section.Steps[stepsCount].AutomationConfig = steps1.AutomationConfig
						flow1.Section.Steps[stepsCount].Description = steps1.Description
						flow1.Section.Steps[stepsCount].DueDate = steps1.DueDate
						flow1.Section.Steps[stepsCount].DynamicDueDate = steps1.DynamicDueDate
						flow1.Section.Steps[stepsCount].Ecp = steps1.Ecp
						flow1.Section.Steps[stepsCount].ExecutionType = steps1.ExecutionType
						flow1.Section.Steps[stepsCount].ID = steps1.ID
						flow1.Section.Steps[stepsCount].Installer = steps1.Installer
						flow1.Section.Steps[stepsCount].Journey = steps1.Journey
						flow1.Section.Steps[stepsCount].Name = steps1.Name
						flow1.Section.Steps[stepsCount].Order = steps1.Order
						flow1.Section.Steps[stepsCount].Requirements = steps1.Requirements
						flow1.Section.Steps[stepsCount].Type = steps1.Type
						flow1.Section.Steps[stepsCount].UserIds = steps1.UserIds
					}
				}
				flow1.Section.Type = types.StringValue(string(flowItem.Section.Type))
			}
			if flowItem.Step != nil {
				flow1.Step = &tfTypes.Step{}
				flow1.Step.AssignedTo = []types.String{}
				for _, v := range flowItem.Step.AssignedTo {
					flow1.Step.AssignedTo = append(flow1.Step.AssignedTo, types.StringValue(v))
				}
				if flowItem.Step.AutomationConfig == nil {
					flow1.Step.AutomationConfig = nil
				} else {
					flow1.Step.AutomationConfig = &tfTypes.AutomationConfig{}
					flow1.Step.AutomationConfig.FlowID = types.StringValue(flowItem.Step.AutomationConfig.FlowID)
				}
				if flowItem.Step.Description == nil {
					flow1.Step.Description = nil
				} else {
					flow1.Step.Description = &tfTypes.StepDescription{}
					flow1.Step.Description.Enabled = types.BoolPointerValue(flowItem.Step.Description.Enabled)
					flow1.Step.Description.Value = types.StringPointerValue(flowItem.Step.Description.Value)
				}
				flow1.Step.DueDate = types.StringPointerValue(flowItem.Step.DueDate)
				if flowItem.Step.DynamicDueDate == nil {
					flow1.Step.DynamicDueDate = nil
				} else {
					flow1.Step.DynamicDueDate = &tfTypes.DynamicDueDate{}
					if flowItem.Step.DynamicDueDate.ActionTypeCondition != nil {
						flow1.Step.DynamicDueDate.ActionTypeCondition = types.StringValue(string(*flowItem.Step.DynamicDueDate.ActionTypeCondition))
					} else {
						flow1.Step.DynamicDueDate.ActionTypeCondition = types.StringNull()
					}
					if flowItem.Step.DynamicDueDate.NumberOfUnits != nil {
						flow1.Step.DynamicDueDate.NumberOfUnits = types.NumberValue(big.NewFloat(float64(*flowItem.Step.DynamicDueDate.NumberOfUnits)))
					} else {
						flow1.Step.DynamicDueDate.NumberOfUnits = types.NumberNull()
					}
					flow1.Step.DynamicDueDate.StepID = types.StringPointerValue(flowItem.Step.DynamicDueDate.StepID)
					if flowItem.Step.DynamicDueDate.TimePeriod != nil {
						flow1.Step.DynamicDueDate.TimePeriod = types.StringValue(string(*flowItem.Step.DynamicDueDate.TimePeriod))
					} else {
						flow1.Step.DynamicDueDate.TimePeriod = types.StringNull()
					}
				}
				if flowItem.Step.Ecp == nil {
					flow1.Step.Ecp = nil
				} else {
					flow1.Step.Ecp = &tfTypes.ECPDetails{}
					flow1.Step.Ecp.Description = types.StringPointerValue(flowItem.Step.Ecp.Description)
					flow1.Step.Ecp.Enabled = types.BoolPointerValue(flowItem.Step.Ecp.Enabled)
					if flowItem.Step.Ecp.Journey == nil {
						flow1.Step.Ecp.Journey = nil
					} else {
						flow1.Step.Ecp.Journey = &tfTypes.StepJourney{}
						flow1.Step.Ecp.Journey.ID = types.StringPointerValue(flowItem.Step.Ecp.Journey.ID)
						flow1.Step.Ecp.Journey.JourneyID = types.StringPointerValue(flowItem.Step.Ecp.Journey.JourneyID)
						flow1.Step.Ecp.Journey.Name = types.StringPointerValue(flowItem.Step.Ecp.Journey.Name)
					}
					flow1.Step.Ecp.Label = types.StringPointerValue(flowItem.Step.Ecp.Label)
				}
				if flowItem.Step.ExecutionType != nil {
					flow1.Step.ExecutionType = types.StringValue(string(*flowItem.Step.ExecutionType))
				} else {
					flow1.Step.ExecutionType = types.StringNull()
				}
				flow1.Step.ID = types.StringPointerValue(flowItem.Step.ID)
				if flowItem.Step.Installer == nil {
					flow1.Step.Installer = nil
				} else {
					flow1.Step.Installer = &tfTypes.ECPDetails{}
					flow1.Step.Installer.Description = types.StringPointerValue(flowItem.Step.Installer.Description)
					flow1.Step.Installer.Enabled = types.BoolPointerValue(flowItem.Step.Installer.Enabled)
					if flowItem.Step.Installer.Journey == nil {
						flow1.Step.Installer.Journey = nil
					} else {
						flow1.Step.Installer.Journey = &tfTypes.StepJourney{}
						flow1.Step.Installer.Journey.ID = types.StringPointerValue(flowItem.Step.Installer.Journey.ID)
						flow1.Step.Installer.Journey.JourneyID = types.StringPointerValue(flowItem.Step.Installer.Journey.JourneyID)
						flow1.Step.Installer.Journey.Name = types.StringPointerValue(flowItem.Step.Installer.Journey.Name)
					}
					flow1.Step.Installer.Label = types.StringPointerValue(flowItem.Step.Installer.Label)
				}
				if flowItem.Step.Journey == nil {
					flow1.Step.Journey = nil
				} else {
					flow1.Step.Journey = &tfTypes.StepJourney{}
					flow1.Step.Journey.ID = types.StringPointerValue(flowItem.Step.Journey.ID)
					flow1.Step.Journey.JourneyID = types.StringPointerValue(flowItem.Step.Journey.JourneyID)
					flow1.Step.Journey.Name = types.StringPointerValue(flowItem.Step.Journey.Name)
				}
				flow1.Step.Name = types.StringValue(flowItem.Step.Name)
				flow1.Step.Order = types.NumberValue(big.NewFloat(float64(flowItem.Step.Order)))
				flow1.Step.Requirements = []tfTypes.StepRequirement{}
				for requirementsCount1, requirementsItem1 := range flowItem.Step.Requirements {
					var requirements3 tfTypes.StepRequirement
					requirements3.Condition = types.StringValue(string(requirementsItem1.Condition))
					requirements3.DefinitionID = types.StringValue(requirementsItem1.DefinitionID)
					requirements3.Type = types.StringValue(string(requirementsItem1.Type))
					if requirementsCount1+1 > len(flow1.Step.Requirements) {
						flow1.Step.Requirements = append(flow1.Step.Requirements, requirements3)
					} else {
						flow1.Step.Requirements[requirementsCount1].Condition = requirements3.Condition
						flow1.Step.Requirements[requirementsCount1].DefinitionID = requirements3.DefinitionID
						flow1.Step.Requirements[requirementsCount1].Type = requirements3.Type
					}
				}
				flow1.Step.Type = types.StringValue(string(flowItem.Step.Type))
				flow1.Step.UserIds = []types.Number{}
				for _, v := range flowItem.Step.UserIds {
					flow1.Step.UserIds = append(flow1.Step.UserIds, types.NumberValue(big.NewFloat(float64(v))))
				}
			}
			if flowCount+1 > len(r.Flow) {
				r.Flow = append(r.Flow, flow1)
			} else {
				r.Flow[flowCount].Section = flow1.Section
				r.Flow[flowCount].Step = flow1.Step
			}
		}
		r.ID = types.StringPointerValue(resp.ID)
		r.LastUpdateTime = types.StringPointerValue(resp.LastUpdateTime)
		r.Name = types.StringValue(resp.Name)
		r.UpdateEntityAttributes = []tfTypes.UpdateEntityAttributes{}
		if len(r.UpdateEntityAttributes) > len(resp.UpdateEntityAttributes) {
			r.UpdateEntityAttributes = r.UpdateEntityAttributes[:len(resp.UpdateEntityAttributes)]
		}
		for updateEntityAttributesCount, updateEntityAttributesItem := range resp.UpdateEntityAttributes {
			var updateEntityAttributes1 tfTypes.UpdateEntityAttributes
			updateEntityAttributes1.Source = types.StringValue(string(updateEntityAttributesItem.Source))
			updateEntityAttributes1.Target.EntityAttribute = types.StringValue(updateEntityAttributesItem.Target.EntityAttribute)
			updateEntityAttributes1.Target.EntitySchema = types.StringValue(updateEntityAttributesItem.Target.EntitySchema)
			if updateEntityAttributesCount+1 > len(r.UpdateEntityAttributes) {
				r.UpdateEntityAttributes = append(r.UpdateEntityAttributes, updateEntityAttributes1)
			} else {
				r.UpdateEntityAttributes[updateEntityAttributesCount].Source = updateEntityAttributes1.Source
				r.UpdateEntityAttributes[updateEntityAttributesCount].Target = updateEntityAttributes1.Target
			}
		}
		r.UserIds = []types.Number{}
		for _, v := range resp.UserIds {
			r.UserIds = append(r.UserIds, types.NumberValue(big.NewFloat(float64(v))))
		}
	}
}
