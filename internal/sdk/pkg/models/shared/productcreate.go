// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/pkg/utils"
)

// ProductCreateCrossSellableProducts - Stores references to products that can be cross sold with the current product.
type ProductCreateCrossSellableProducts struct {
	DollarRelation []BaseRelation `json:"$relation,omitempty"`
}

func (o *ProductCreateCrossSellableProducts) GetDollarRelation() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

type ProductCreateFeature struct {
	// An arbitrary set of tags attached to a feature
	Tags    []string `json:"_tags,omitempty"`
	Feature *string  `json:"feature,omitempty"`
}

func (o *ProductCreateFeature) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ProductCreateFeature) GetFeature() *string {
	if o == nil {
		return nil
	}
	return o.Feature
}

// ProductCreateProductDownloads - Stores references to a set of files downloadable from the product.
// e.g: tech specifications, quality control sheets, privacy policy agreements
type ProductCreateProductDownloads struct {
	DollarRelation []BaseRelation `json:"$relation,omitempty"`
}

func (o *ProductCreateProductDownloads) GetDollarRelation() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// ProductCreateProductImages - Stores references to a set of file images of the product
type ProductCreateProductImages struct {
	DollarRelation []BaseRelation `json:"$relation,omitempty"`
}

func (o *ProductCreateProductImages) GetDollarRelation() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.DollarRelation
}

// ProductCreateType - The type of Product:
//
// | type | description |
// |----| ----|
// | `product` | Represents a physical good |
// | `service` | Represents a service or virtual product |
type ProductCreateType string

const (
	ProductCreateTypeProduct ProductCreateType = "product"
	ProductCreateTypeService ProductCreateType = "service"
)

func (e ProductCreateType) ToPointer() *ProductCreateType {
	return &e
}

func (e *ProductCreateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "product":
		fallthrough
	case "service":
		*e = ProductCreateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProductCreateType: %v", v)
	}
}

type ProductCreate struct {
	// Stores references to the availability files that define where this product is available.
	// These files are used when interacting with products via epilot Journeys, thought the AvailabilityCheck block.
	//
	AvailabilityFiles []BaseRelation `json:"availability_files,omitempty"`
	// The product code
	Code *string `json:"code,omitempty"`
	// Stores references to products that can be cross sold with the current product.
	CrossSellableProducts *ProductCreateCrossSellableProducts `json:"cross_sellable_products,omitempty"`
	// A description of the product. Multi-line supported.
	Description *string                `json:"description,omitempty"`
	Feature     []ProductCreateFeature `json:"feature,omitempty"`
	// Not visible to customers, only in internal tables
	InternalName *string `json:"internal_name,omitempty"`
	// The description for the product
	Name         string        `json:"name"`
	PriceOptions *BaseRelation `json:"price_options,omitempty"`
	// Stores references to a set of files downloadable from the product.
	// e.g: tech specifications, quality control sheets, privacy policy agreements
	//
	ProductDownloads *ProductCreateProductDownloads `json:"product_downloads,omitempty"`
	// Stores references to a set of file images of the product
	ProductImages *ProductCreateProductImages `json:"product_images,omitempty"`
	// The type of Product:
	//
	// | type | description |
	// |----| ----|
	// | `product` | Represents a physical good |
	// | `service` | Represents a service or virtual product |
	//
	Type *ProductCreateType `default:"product" json:"type"`
}

func (p ProductCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProductCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProductCreate) GetAvailabilityFiles() []BaseRelation {
	if o == nil {
		return nil
	}
	return o.AvailabilityFiles
}

func (o *ProductCreate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ProductCreate) GetCrossSellableProducts() *ProductCreateCrossSellableProducts {
	if o == nil {
		return nil
	}
	return o.CrossSellableProducts
}

func (o *ProductCreate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ProductCreate) GetFeature() []ProductCreateFeature {
	if o == nil {
		return nil
	}
	return o.Feature
}

func (o *ProductCreate) GetInternalName() *string {
	if o == nil {
		return nil
	}
	return o.InternalName
}

func (o *ProductCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ProductCreate) GetPriceOptions() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.PriceOptions
}

func (o *ProductCreate) GetProductDownloads() *ProductCreateProductDownloads {
	if o == nil {
		return nil
	}
	return o.ProductDownloads
}

func (o *ProductCreate) GetProductImages() *ProductCreateProductImages {
	if o == nil {
		return nil
	}
	return o.ProductImages
}

func (o *ProductCreate) GetType() *ProductCreateType {
	if o == nil {
		return nil
	}
	return o.Type
}
