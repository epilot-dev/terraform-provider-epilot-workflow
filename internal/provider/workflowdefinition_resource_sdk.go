// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
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
		var id string
		id = closingReasonsItem.ID.ValueString()

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
		actionTypeCondition := shared.ActionTypeCondition(r.DynamicDueDate.ActionTypeCondition.ValueString())
		var numberOfUnits float64
		numberOfUnits, _ = r.DynamicDueDate.NumberOfUnits.ValueBigFloat().Float64()

		phaseID := new(string)
		if !r.DynamicDueDate.PhaseID.IsUnknown() && !r.DynamicDueDate.PhaseID.IsNull() {
			*phaseID = r.DynamicDueDate.PhaseID.ValueString()
		} else {
			phaseID = nil
		}
		stepID := new(string)
		if !r.DynamicDueDate.StepID.IsUnknown() && !r.DynamicDueDate.StepID.IsNull() {
			*stepID = r.DynamicDueDate.StepID.ValueString()
		} else {
			stepID = nil
		}
		timePeriod := shared.TimeUnit(r.DynamicDueDate.TimePeriod.ValueString())
		dynamicDueDate = &shared.DynamicDueDate{
			ActionTypeCondition: actionTypeCondition,
			NumberOfUnits:       numberOfUnits,
			PhaseID:             phaseID,
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
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var flow interface{}
	_ = json.Unmarshal([]byte(r.Flow.ValueString()), &flow)
	id1 := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		*id1 = r.ID.ValueString()
	} else {
		id1 = nil
	}
	lastUpdateTime := new(string)
	if !r.LastUpdateTime.IsUnknown() && !r.LastUpdateTime.IsNull() {
		*lastUpdateTime = r.LastUpdateTime.ValueString()
	} else {
		lastUpdateTime = nil
	}
	var name string
	name = r.Name.ValueString()

	var taxonomies []string = []string{}
	for _, taxonomiesItem := range r.Taxonomies {
		taxonomies = append(taxonomies, taxonomiesItem.ValueString())
	}
	var updateEntityAttributes []shared.UpdateEntityAttributes = []shared.UpdateEntityAttributes{}
	for _, updateEntityAttributesItem := range r.UpdateEntityAttributes {
		source := shared.Source(updateEntityAttributesItem.Source.ValueString())
		var entityAttribute string
		entityAttribute = updateEntityAttributesItem.Target.EntityAttribute.ValueString()

		var entitySchema string
		entitySchema = updateEntityAttributesItem.Target.EntitySchema.ValueString()

		target := shared.Target{
			EntityAttribute: entityAttribute,
			EntitySchema:    entitySchema,
		}
		updateEntityAttributes = append(updateEntityAttributes, shared.UpdateEntityAttributes{
			Source: source,
			Target: target,
		})
	}
	var userIds []float64 = []float64{}
	for _, userIdsItem := range r.UserIds {
		userIdsTmp, _ := userIdsItem.ValueBigFloat().Float64()
		userIds = append(userIds, userIdsTmp)
	}
	out := shared.WorkflowDefinition{
		AssignedTo:             assignedTo,
		ClosingReasons:         closingReasons,
		CreationTime:           creationTime,
		Description:            description,
		DueDate:                dueDate,
		DynamicDueDate:         dynamicDueDate,
		EnableECPWorkflow:      enableECPWorkflow,
		Enabled:                enabled,
		Flow:                   flow,
		ID:                     id1,
		LastUpdateTime:         lastUpdateTime,
		Name:                   name,
		Taxonomies:             taxonomies,
		UpdateEntityAttributes: updateEntityAttributes,
		UserIds:                userIds,
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
			r.DynamicDueDate.ActionTypeCondition = types.StringValue(string(resp.DynamicDueDate.ActionTypeCondition))
			r.DynamicDueDate.NumberOfUnits = types.NumberValue(big.NewFloat(float64(resp.DynamicDueDate.NumberOfUnits)))
			r.DynamicDueDate.PhaseID = types.StringPointerValue(resp.DynamicDueDate.PhaseID)
			r.DynamicDueDate.StepID = types.StringPointerValue(resp.DynamicDueDate.StepID)
			r.DynamicDueDate.TimePeriod = types.StringValue(string(resp.DynamicDueDate.TimePeriod))
		}
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.EnableECPWorkflow = types.BoolPointerValue(resp.EnableECPWorkflow)
		flowResult, _ := json.Marshal(resp.Flow)
		r.Flow = types.StringValue(string(flowResult))
		r.ID = types.StringPointerValue(resp.ID)
		r.LastUpdateTime = types.StringPointerValue(resp.LastUpdateTime)
		r.Name = types.StringValue(resp.Name)
		r.Taxonomies = []types.String{}
		for _, v := range resp.Taxonomies {
			r.Taxonomies = append(r.Taxonomies, types.StringValue(v))
		}
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
