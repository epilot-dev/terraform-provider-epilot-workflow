resource "epilot-product_product" "my_product" {
  code        = "...my_code..."
  description = "...my_description..."
  feature = [
    {
      tags = [
        "...",
      ]
      feature = "...my_feature..."
    },
  ]
  internal_name = "...my_internal_name..."
  name          = "Dallas Von"
  price_options = {
    dollar_relation = [
      {
        tags = [
          "...",
        ]
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
      },
    ]
  }
  type = "product"
}