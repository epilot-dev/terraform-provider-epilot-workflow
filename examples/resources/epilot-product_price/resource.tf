resource "epilot-product_price" "my_price" {
  active                    = true
  billing_duration_amount   = 66.76
  billing_duration_unit     = "weeks"
  description               = "...my_description..."
  is_composite_price        = false
  is_tax_inclusive          = true
  long_description          = "...my_long_description..."
  notice_time_amount        = 2.66
  notice_time_unit          = "months"
  price_display_in_journeys = "show_as_starting_price"
  pricing_model             = "per_unit"
  renewal_duration_amount   = 66.53
  renewal_duration_unit     = "months"
  tax = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
      },
    ]
  }
  termination_time_amount = 44.65
  termination_time_unit   = "years"
  tiers = [
    {
      display_mode            = "hidden"
      flat_fee_amount         = 40.89
      flat_fee_amount_decimal = "...my_flat_fee_amount_decimal..."
      unit_amount             = 62.81
      unit_amount_decimal     = "...my_unit_amount_decimal..."
      up_to                   = 50.16
    },
  ]
  type = "recurring"
  unit = {
    str = "..."
  }
  unit_amount          = 93.37
  unit_amount_currency = "EUR"
  unit_amount_decimal  = "...my_unit_amount_decimal..."
  variable_price       = false
}