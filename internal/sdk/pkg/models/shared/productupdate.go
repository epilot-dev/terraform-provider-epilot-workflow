// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/utils"
)

// ProductUpdateCrossSellableProducts - Stores references to products that can be cross sold with the current product.
type ProductUpdateCrossSellableProducts struct {
	DollarRelation []BaseRelation `json:"$relation,omitempty"`
}

func (o *ProductUpdateCrossSellableProducts) GetDollarRelation() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

type ProductUpdateFeature struct {
	// An arbitrary set of tags attached to a feature
	Tags    []string `json:"_tags,omitempty"`
	Feature *string  `json:"feature,omitempty"`
}

func (o *ProductUpdateFeature) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ProductUpdateFeature) GetFeature() *string {
	if o == nil {
		return nil
	}
	return o.Feature
}

// ProductUpdateProductDownloads - Stores references to a set of files downloadable from the product.
// e.g: tech specifications, quality control sheets, privacy policy agreements
type ProductUpdateProductDownloads struct {
	DollarRelation []BaseRelation `json:"$relation,omitempty"`
}

func (o *ProductUpdateProductDownloads) GetDollarRelation() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// ProductUpdateProductImages - Stores references to a set of file images of the product
type ProductUpdateProductImages struct {
	DollarRelation []BaseRelation `json:"$relation,omitempty"`
}

func (o *ProductUpdateProductImages) GetDollarRelation() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// ProductUpdateType - The type of Product:
//
// | type | description |
// |----| ----|
// | `product` | Represents a physical good |
// | `service` | Represents a service or virtual product |
type ProductUpdateType string

const (
	ProductUpdateTypeProduct ProductUpdateType = "product"
	ProductUpdateTypeService ProductUpdateType = "service"
)

func (e ProductUpdateType) ToPointer() *ProductUpdateType {
	return &e
}

func (e *ProductUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "product":
		fallthrough
	case "service":
		*e = ProductUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductUpdateType: %v", v)
	}
}

type ProductUpdate struct {
	// Stores references to the availability files that define where this product is available.
	// These files are used when interacting with products via epilot Journeys, thought the AvailabilityCheck block.
	//
	AvailabilityFiles []BaseRelation `json:"availability_files,omitempty"`
	// The product code
	Code *string `json:"code,omitempty"`
	// Stores references to products that can be cross sold with the current product.
	CrossSellableProducts *ProductUpdateCrossSellableProducts `json:"cross_sellable_products,omitempty"`
	// A description of the product. Multi-line supported.
	Description *string                `json:"description,omitempty"`
	Feature     []ProductUpdateFeature `json:"feature,omitempty"`
	// Not visible to customers, only in internal tables
	InternalName *string `json:"internal_name,omitempty"`
	// The description for the product
	Name         *string       `json:"name,omitempty"`
	PriceOptions *BaseRelation `json:"price_options,omitempty"`
	// Stores references to a set of files downloadable from the product.
	// e.g: tech specifications, quality control sheets, privacy policy agreements
	//
	ProductDownloads *ProductUpdateProductDownloads `json:"product_downloads,omitempty"`
	// Stores references to a set of file images of the product
	ProductImages *ProductUpdateProductImages `json:"product_images,omitempty"`
	// The type of Product:
	//
	// | type | description |
	// |----| ----|
	// | `product` | Represents a physical good |
	// | `service` | Represents a service or virtual product |
	//
	Type *ProductUpdateType `default:"product" json:"type"`
}

func (p ProductUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductUpdate) GetAvailabilityFiles() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.AvailabilityFiles
}

func (o *ProductUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ProductUpdate) GetCrossSellableProducts() *ProductUpdateCrossSellableProducts {
	if o == nil {
		return nil
	}
	return o.CrossSellableProducts
}

func (o *ProductUpdate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ProductUpdate) GetFeature() []ProductUpdateFeature {
	if o == nil {
		return nil
	}
	return o.Feature
}

func (o *ProductUpdate) GetInternalName() *string {
	if o == nil {
		return nil
	}
	return o.InternalName
}

func (o *ProductUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ProductUpdate) GetPriceOptions() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.PriceOptions
}

func (o *ProductUpdate) GetProductDownloads() *ProductUpdateProductDownloads {
	if o == nil {
		return nil
	}
	return o.ProductDownloads
}

func (o *ProductUpdate) GetProductImages() *ProductUpdateProductImages {
	if o == nil {
		return nil
	}
	return o.ProductImages
}

func (o *ProductUpdate) GetType() *ProductUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
