resource "epilot-product_tax" "my_tax" {
  active      = true
  description = "...my_description..."
  rate        = "...my_rate..."
  region      = "CH"
  type        = "VAT"
}