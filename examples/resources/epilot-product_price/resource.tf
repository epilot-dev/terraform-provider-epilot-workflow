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
  tax                       = "{ \"see\": \"documentation\" }"
  termination_time_amount   = 20
  termination_time_unit     = "weeks"
  type                      = "recurring"
  unit_amount               = 84.62
  unit_amount_currency      = "EUR"
  unit_amount_decimal       = "...my_unit_amount_decimal..."
  variable_price            = false
}