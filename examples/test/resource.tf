# epilot-product_product.s_product:
resource "epilot-product_product" "s_product" {
    active        = true
    internal_name = "Price testing"
    name          = "Night storage heating tariff"
    price_options = {
        dollar_relation = [
            {
                entity_id = "f6194ac9-eb96-49e3-a21a-9ee9d19420b9"
            },
        ]
    }
    type          = "product"
}

# epilot-product_tax.s_tax:
resource "epilot-product_tax" "s_tax" {
    active = true
    rate   = "0"
    region = "DE"
    type   = "VAT"
}
