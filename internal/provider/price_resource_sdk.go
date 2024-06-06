// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
)

func (r *PriceResourceModel) ToSharedPriceCreate() *shared.PriceCreate {
	active := r.Active.ValueBool()
	billingDurationAmount := new(float64)
	if !r.BillingDurationAmount.IsUnknown() && !r.BillingDurationAmount.IsNull() {
		*billingDurationAmount, _ = r.BillingDurationAmount.ValueBigFloat().Float64()
	} else {
		billingDurationAmount = nil
	}
	billingDurationUnit := new(shared.PriceCreateBillingDurationUnit)
	if !r.BillingDurationUnit.IsUnknown() && !r.BillingDurationUnit.IsNull() {
		*billingDurationUnit = shared.PriceCreateBillingDurationUnit(r.BillingDurationUnit.ValueString())
	} else {
		billingDurationUnit = nil
	}
	description := r.Description.ValueString()
	isCompositePrice := new(bool)
	if !r.IsCompositePrice.IsUnknown() && !r.IsCompositePrice.IsNull() {
		*isCompositePrice = r.IsCompositePrice.ValueBool()
	} else {
		isCompositePrice = nil
	}
	isTaxInclusive := new(bool)
	if !r.IsTaxInclusive.IsUnknown() && !r.IsTaxInclusive.IsNull() {
		*isTaxInclusive = r.IsTaxInclusive.ValueBool()
	} else {
		isTaxInclusive = nil
	}
	longDescription := new(string)
	if !r.LongDescription.IsUnknown() && !r.LongDescription.IsNull() {
		*longDescription = r.LongDescription.ValueString()
	} else {
		longDescription = nil
	}
	noticeTimeAmount := new(float64)
	if !r.NoticeTimeAmount.IsUnknown() && !r.NoticeTimeAmount.IsNull() {
		*noticeTimeAmount, _ = r.NoticeTimeAmount.ValueBigFloat().Float64()
	} else {
		noticeTimeAmount = nil
	}
	noticeTimeUnit := new(shared.PriceCreateNoticeTimeUnit)
	if !r.NoticeTimeUnit.IsUnknown() && !r.NoticeTimeUnit.IsNull() {
		*noticeTimeUnit = shared.PriceCreateNoticeTimeUnit(r.NoticeTimeUnit.ValueString())
	} else {
		noticeTimeUnit = nil
	}
	priceDisplayInJourneys := new(shared.PriceCreatePriceDisplayInJourneys)
	if !r.PriceDisplayInJourneys.IsUnknown() && !r.PriceDisplayInJourneys.IsNull() {
		*priceDisplayInJourneys = shared.PriceCreatePriceDisplayInJourneys(r.PriceDisplayInJourneys.ValueString())
	} else {
		priceDisplayInJourneys = nil
	}
	pricingModel := new(shared.PriceCreatePricingModel)
	if !r.PricingModel.IsUnknown() && !r.PricingModel.IsNull() {
		*pricingModel = shared.PriceCreatePricingModel(r.PricingModel.ValueString())
	} else {
		pricingModel = nil
	}
	renewalDurationAmount := new(float64)
	if !r.RenewalDurationAmount.IsUnknown() && !r.RenewalDurationAmount.IsNull() {
		*renewalDurationAmount, _ = r.RenewalDurationAmount.ValueBigFloat().Float64()
	} else {
		renewalDurationAmount = nil
	}
	renewalDurationUnit := new(shared.PriceCreateRenewalDurationUnit)
	if !r.RenewalDurationUnit.IsUnknown() && !r.RenewalDurationUnit.IsNull() {
		*renewalDurationUnit = shared.PriceCreateRenewalDurationUnit(r.RenewalDurationUnit.ValueString())
	} else {
		renewalDurationUnit = nil
	}
	var tax interface{}
	if !r.Tax.IsUnknown() && !r.Tax.IsNull() {
		_ = json.Unmarshal([]byte(r.Tax.ValueString()), &tax)
	}
	terminationTimeAmount := new(float64)
	if !r.TerminationTimeAmount.IsUnknown() && !r.TerminationTimeAmount.IsNull() {
		*terminationTimeAmount, _ = r.TerminationTimeAmount.ValueBigFloat().Float64()
	} else {
		terminationTimeAmount = nil
	}
	terminationTimeUnit := new(shared.PriceCreateTerminationTimeUnit)
	if !r.TerminationTimeUnit.IsUnknown() && !r.TerminationTimeUnit.IsNull() {
		*terminationTimeUnit = shared.PriceCreateTerminationTimeUnit(r.TerminationTimeUnit.ValueString())
	} else {
		terminationTimeUnit = nil
	}
	var tiers []shared.PriceTier = []shared.PriceTier{}
	for _, tiersItem := range r.Tiers {
		displayMode := new(shared.PriceTierDisplayMode)
		if !tiersItem.DisplayMode.IsUnknown() && !tiersItem.DisplayMode.IsNull() {
			*displayMode = shared.PriceTierDisplayMode(tiersItem.DisplayMode.ValueString())
		} else {
			displayMode = nil
		}
		flatFeeAmount := new(float64)
		if !tiersItem.FlatFeeAmount.IsUnknown() && !tiersItem.FlatFeeAmount.IsNull() {
			*flatFeeAmount, _ = tiersItem.FlatFeeAmount.ValueBigFloat().Float64()
		} else {
			flatFeeAmount = nil
		}
		flatFeeAmountDecimal := new(string)
		if !tiersItem.FlatFeeAmountDecimal.IsUnknown() && !tiersItem.FlatFeeAmountDecimal.IsNull() {
			*flatFeeAmountDecimal = tiersItem.FlatFeeAmountDecimal.ValueString()
		} else {
			flatFeeAmountDecimal = nil
		}
		unitAmount := new(float64)
		if !tiersItem.UnitAmount.IsUnknown() && !tiersItem.UnitAmount.IsNull() {
			*unitAmount, _ = tiersItem.UnitAmount.ValueBigFloat().Float64()
		} else {
			unitAmount = nil
		}
		unitAmountDecimal := new(string)
		if !tiersItem.UnitAmountDecimal.IsUnknown() && !tiersItem.UnitAmountDecimal.IsNull() {
			*unitAmountDecimal = tiersItem.UnitAmountDecimal.ValueString()
		} else {
			unitAmountDecimal = nil
		}
		upTo := new(float64)
		if !tiersItem.UpTo.IsUnknown() && !tiersItem.UpTo.IsNull() {
			*upTo, _ = tiersItem.UpTo.ValueBigFloat().Float64()
		} else {
			upTo = nil
		}
		tiers = append(tiers, shared.PriceTier{
			DisplayMode:          displayMode,
			FlatFeeAmount:        flatFeeAmount,
			FlatFeeAmountDecimal: flatFeeAmountDecimal,
			UnitAmount:           unitAmount,
			UnitAmountDecimal:    unitAmountDecimal,
			UpTo:                 upTo,
		})
	}
	typeVar := new(shared.PriceCreateType)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeVar = shared.PriceCreateType(r.Type.ValueString())
	} else {
		typeVar = nil
	}
	var unit *shared.PriceCreateUnit
	if r.Unit != nil {
		priceCreate1 := new(shared.PriceCreate1)
		if !r.Unit.One.IsUnknown() && !r.Unit.One.IsNull() {
			*priceCreate1 = shared.PriceCreate1(r.Unit.One.ValueString())
		} else {
			priceCreate1 = nil
		}
		if priceCreate1 != nil {
			unit = &shared.PriceCreateUnit{
				PriceCreate1: priceCreate1,
			}
		}
		str := new(string)
		if !r.Unit.Str.IsUnknown() && !r.Unit.Str.IsNull() {
			*str = r.Unit.Str.ValueString()
		} else {
			str = nil
		}
		if str != nil {
			unit = &shared.PriceCreateUnit{
				Str: str,
			}
		}
	}
	unitAmount1 := new(float64)
	if !r.UnitAmount.IsUnknown() && !r.UnitAmount.IsNull() {
		*unitAmount1, _ = r.UnitAmount.ValueBigFloat().Float64()
	} else {
		unitAmount1 = nil
	}
	unitAmountCurrency := new(string)
	if !r.UnitAmountCurrency.IsUnknown() && !r.UnitAmountCurrency.IsNull() {
		*unitAmountCurrency = r.UnitAmountCurrency.ValueString()
	} else {
		unitAmountCurrency = nil
	}
	unitAmountDecimal1 := new(string)
	if !r.UnitAmountDecimal.IsUnknown() && !r.UnitAmountDecimal.IsNull() {
		*unitAmountDecimal1 = r.UnitAmountDecimal.ValueString()
	} else {
		unitAmountDecimal1 = nil
	}
	variablePrice := new(bool)
	if !r.VariablePrice.IsUnknown() && !r.VariablePrice.IsNull() {
		*variablePrice = r.VariablePrice.ValueBool()
	} else {
		variablePrice = nil
	}
	out := shared.PriceCreate{
		Active:                 active,
		BillingDurationAmount:  billingDurationAmount,
		BillingDurationUnit:    billingDurationUnit,
		Description:            description,
		IsCompositePrice:       isCompositePrice,
		IsTaxInclusive:         isTaxInclusive,
		LongDescription:        longDescription,
		NoticeTimeAmount:       noticeTimeAmount,
		NoticeTimeUnit:         noticeTimeUnit,
		PriceDisplayInJourneys: priceDisplayInJourneys,
		PricingModel:           pricingModel,
		RenewalDurationAmount:  renewalDurationAmount,
		RenewalDurationUnit:    renewalDurationUnit,
		Tax:                    tax,
		TerminationTimeAmount:  terminationTimeAmount,
		TerminationTimeUnit:    terminationTimeUnit,
		Tiers:                  tiers,
		Type:                   typeVar,
		Unit:                   unit,
		UnitAmount:             unitAmount1,
		UnitAmountCurrency:     unitAmountCurrency,
		UnitAmountDecimal:      unitAmountDecimal1,
		VariablePrice:          variablePrice,
	}
	return &out
}

func (r *PriceResourceModel) RefreshFromSharedPrice(resp *shared.Price) {
	if resp != nil {
		r.ID = types.StringValue(resp.ID)
		r.Active = types.BoolValue(resp.Active)
		if resp.BillingDurationAmount != nil {
			r.BillingDurationAmount = types.NumberValue(big.NewFloat(float64(*resp.BillingDurationAmount)))
		} else {
			r.BillingDurationAmount = types.NumberNull()
		}
		if resp.BillingDurationUnit != nil {
			r.BillingDurationUnit = types.StringValue(string(*resp.BillingDurationUnit))
		} else {
			r.BillingDurationUnit = types.StringNull()
		}
		r.Description = types.StringValue(resp.Description)
		r.IsCompositePrice = types.BoolPointerValue(resp.IsCompositePrice)
		r.IsTaxInclusive = types.BoolPointerValue(resp.IsTaxInclusive)
		r.LongDescription = types.StringPointerValue(resp.LongDescription)
		if resp.NoticeTimeAmount != nil {
			r.NoticeTimeAmount = types.NumberValue(big.NewFloat(float64(*resp.NoticeTimeAmount)))
		} else {
			r.NoticeTimeAmount = types.NumberNull()
		}
		if resp.NoticeTimeUnit != nil {
			r.NoticeTimeUnit = types.StringValue(string(*resp.NoticeTimeUnit))
		} else {
			r.NoticeTimeUnit = types.StringNull()
		}
		if resp.PriceDisplayInJourneys != nil {
			r.PriceDisplayInJourneys = types.StringValue(string(*resp.PriceDisplayInJourneys))
		} else {
			r.PriceDisplayInJourneys = types.StringNull()
		}
		if resp.PricingModel != nil {
			r.PricingModel = types.StringValue(string(*resp.PricingModel))
		} else {
			r.PricingModel = types.StringNull()
		}
		if resp.RenewalDurationAmount != nil {
			r.RenewalDurationAmount = types.NumberValue(big.NewFloat(float64(*resp.RenewalDurationAmount)))
		} else {
			r.RenewalDurationAmount = types.NumberNull()
		}
		if resp.RenewalDurationUnit != nil {
			r.RenewalDurationUnit = types.StringValue(string(*resp.RenewalDurationUnit))
		} else {
			r.RenewalDurationUnit = types.StringNull()
		}
		if resp.Tax == nil {
			r.Tax = types.StringNull()
		} else {
			taxResult, _ := json.Marshal(resp.Tax)
			r.Tax = types.StringValue(string(taxResult))
		}
		if resp.TerminationTimeAmount != nil {
			r.TerminationTimeAmount = types.NumberValue(big.NewFloat(float64(*resp.TerminationTimeAmount)))
		} else {
			r.TerminationTimeAmount = types.NumberNull()
		}
		if resp.TerminationTimeUnit != nil {
			r.TerminationTimeUnit = types.StringValue(string(*resp.TerminationTimeUnit))
		} else {
			r.TerminationTimeUnit = types.StringNull()
		}
		r.Tiers = []tfTypes.PriceTier{}
		if len(r.Tiers) > len(resp.Tiers) {
			r.Tiers = r.Tiers[:len(resp.Tiers)]
		}
		for tiersCount, tiersItem := range resp.Tiers {
			var tiers1 tfTypes.PriceTier
			if tiersItem.DisplayMode != nil {
				tiers1.DisplayMode = types.StringValue(string(*tiersItem.DisplayMode))
			} else {
				tiers1.DisplayMode = types.StringNull()
			}
			if tiersItem.FlatFeeAmount != nil {
				tiers1.FlatFeeAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.FlatFeeAmount)))
			} else {
				tiers1.FlatFeeAmount = types.NumberNull()
			}
			tiers1.FlatFeeAmountDecimal = types.StringPointerValue(tiersItem.FlatFeeAmountDecimal)
			if tiersItem.UnitAmount != nil {
				tiers1.UnitAmount = types.NumberValue(big.NewFloat(float64(*tiersItem.UnitAmount)))
			} else {
				tiers1.UnitAmount = types.NumberNull()
			}
			tiers1.UnitAmountDecimal = types.StringPointerValue(tiersItem.UnitAmountDecimal)
			if tiersItem.UpTo != nil {
				tiers1.UpTo = types.NumberValue(big.NewFloat(float64(*tiersItem.UpTo)))
			} else {
				tiers1.UpTo = types.NumberNull()
			}
			if tiersCount+1 > len(r.Tiers) {
				r.Tiers = append(r.Tiers, tiers1)
			} else {
				r.Tiers[tiersCount].DisplayMode = tiers1.DisplayMode
				r.Tiers[tiersCount].FlatFeeAmount = tiers1.FlatFeeAmount
				r.Tiers[tiersCount].FlatFeeAmountDecimal = tiers1.FlatFeeAmountDecimal
				r.Tiers[tiersCount].UnitAmount = tiers1.UnitAmount
				r.Tiers[tiersCount].UnitAmountDecimal = tiers1.UnitAmountDecimal
				r.Tiers[tiersCount].UpTo = tiers1.UpTo
			}
		}
		if resp.Type != nil {
			r.Type = types.StringValue(string(*resp.Type))
		} else {
			r.Type = types.StringNull()
		}
		if resp.Unit == nil {
			r.Unit = nil
		} else {
			r.Unit = &tfTypes.PriceCreateUnit{}
			if resp.Unit.Str != nil {
				r.Unit.Str = types.StringPointerValue(resp.Unit.Str)
			}
			if resp.Unit.One != nil {
				if resp.Unit.One != nil {
					r.Unit.One = types.StringValue(string(*resp.Unit.One))
				} else {
					r.Unit.One = types.StringNull()
				}
			}
		}
		if resp.UnitAmount != nil {
			r.UnitAmount = types.NumberValue(big.NewFloat(float64(*resp.UnitAmount)))
		} else {
			r.UnitAmount = types.NumberNull()
		}
		r.UnitAmountCurrency = types.StringPointerValue(resp.UnitAmountCurrency)
		r.UnitAmountDecimal = types.StringPointerValue(resp.UnitAmountDecimal)
		r.VariablePrice = types.BoolPointerValue(resp.VariablePrice)
	}
}
