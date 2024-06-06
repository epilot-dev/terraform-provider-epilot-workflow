// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ProductResourceModel) ToSharedProductCreate() *shared.ProductCreate {
	active := r.Active.ValueBool()
	code := new(string)
	if !r.Code.IsUnknown() && !r.Code.IsNull() {
		*code = r.Code.ValueString()
	} else {
		code = nil
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	var feature []interface{} = []interface{}{}
	for _, featureItem := range r.Feature {
		var featureTmp interface{}
		_ = json.Unmarshal([]byte(featureItem.ValueString()), &featureTmp)
		feature = append(feature, featureTmp)
	}
	internalName := new(string)
	if !r.InternalName.IsUnknown() && !r.InternalName.IsNull() {
		*internalName = r.InternalName.ValueString()
	} else {
		internalName = nil
	}
	name := r.Name.ValueString()
	var priceOptions *shared.BaseRelation
	if r.PriceOptions != nil {
		var dollarRelation []shared.DollarRelation = []shared.DollarRelation{}
		for _, dollarRelationItem := range r.PriceOptions.DollarRelation {
			var tags []string = []string{}
			for _, tagsItem := range dollarRelationItem.Tags {
				tags = append(tags, tagsItem.ValueString())
			}
			entityID := new(string)
			if !dollarRelationItem.EntityID.IsUnknown() && !dollarRelationItem.EntityID.IsNull() {
				*entityID = dollarRelationItem.EntityID.ValueString()
			} else {
				entityID = nil
			}
			dollarRelation = append(dollarRelation, shared.DollarRelation{
				Tags:     tags,
				EntityID: entityID,
			})
		}
		priceOptions = &shared.BaseRelation{
			DollarRelation: dollarRelation,
		}
	}
	var productDownloads interface{}
	if !r.ProductDownloads.IsUnknown() && !r.ProductDownloads.IsNull() {
		_ = json.Unmarshal([]byte(r.ProductDownloads.ValueString()), &productDownloads)
	}
	var productImages interface{}
	if !r.ProductImages.IsUnknown() && !r.ProductImages.IsNull() {
		_ = json.Unmarshal([]byte(r.ProductImages.ValueString()), &productImages)
	}
	typeVar := new(shared.ProductCreateType)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = shared.ProductCreateType(r.Type.ValueString())
	} else {
		typeVar = nil
	}
	out := shared.ProductCreate{
		Active:           active,
		Code:             code,
		Description:      description,
		Feature:          feature,
		InternalName:     internalName,
		Name:             name,
		PriceOptions:     priceOptions,
		ProductDownloads: productDownloads,
		ProductImages:    productImages,
		Type:             typeVar,
	}
	return &out
}

func (r *ProductResourceModel) RefreshFromSharedProduct(resp *shared.Product) {
	if resp != nil {
		r.ID = types.StringValue(resp.ID)
		r.Active = types.BoolValue(resp.Active)
		r.Code = types.StringPointerValue(resp.Code)
		r.Description = types.StringPointerValue(resp.Description)
		r.Feature = nil
		for _, featureItem := range resp.Feature {
			var feature1 types.String
			feature1Result, _ := json.Marshal(featureItem)
			feature1 = types.StringValue(string(feature1Result))
			r.Feature = append(r.Feature, feature1)
		}
		r.InternalName = types.StringPointerValue(resp.InternalName)
		r.Name = types.StringValue(resp.Name)
		if resp.PriceOptions == nil {
			r.PriceOptions = nil
		} else {
			r.PriceOptions = &tfTypes.BaseRelation{}
			r.PriceOptions.DollarRelation = []tfTypes.DollarRelation{}
			if len(r.PriceOptions.DollarRelation) > len(resp.PriceOptions.DollarRelation) {
				r.PriceOptions.DollarRelation = r.PriceOptions.DollarRelation[:len(resp.PriceOptions.DollarRelation)]
			}
			for dollarRelationCount, dollarRelationItem := range resp.PriceOptions.DollarRelation {
				var dollarRelation1 tfTypes.DollarRelation
				dollarRelation1.Tags = []types.String{}
				for _, v := range dollarRelationItem.Tags {
					dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
				}
				dollarRelation1.EntityID = types.StringPointerValue(dollarRelationItem.EntityID)
				if dollarRelationCount+1 > len(r.PriceOptions.DollarRelation) {
					r.PriceOptions.DollarRelation = append(r.PriceOptions.DollarRelation, dollarRelation1)
				} else {
					r.PriceOptions.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
					r.PriceOptions.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
				}
			}
		}
		if resp.ProductDownloads == nil {
			r.ProductDownloads = types.StringNull()
		} else {
			productDownloadsResult, _ := json.Marshal(resp.ProductDownloads)
			r.ProductDownloads = types.StringValue(string(productDownloadsResult))
		}
		if resp.ProductImages == nil {
			r.ProductImages = types.StringNull()
		} else {
			productImagesResult, _ := json.Marshal(resp.ProductImages)
			r.ProductImages = types.StringValue(string(productImagesResult))
		}
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
	}
}
