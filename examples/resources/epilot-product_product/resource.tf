resource "epilot-product_product" "my_product" {
  availability_files = [
    {
      dollar_relation = [
        {
          entity_id = "123e4567-e89b-12d3-a456-426614174000"
        },
      ]
    },
  ]
  code = "...my_code..."
  cross_sellable_products = {
    dollar_relation = [
      {
        dollar_relation = [
          {
            entity_id = "123e4567-e89b-12d3-a456-426614174000"
          },
        ]
      },
    ]
  }
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
  name          = "Faye Schiller"
  price_options = {
    dollar_relation = [
      {
        entity_id = "123e4567-e89b-12d3-a456-426614174000"
      },
    ]
  }
  product_downloads = {
    dollar_relation = [
      {
        dollar_relation = [
          {
            entity_id = "123e4567-e89b-12d3-a456-426614174000"
          },
        ]
      },
    ]
  }
  product_images = {
    dollar_relation = [
      {
        dollar_relation = [
          {
            entity_id = "123e4567-e89b-12d3-a456-426614174000"
          },
        ]
      },
    ]
  }
  type = "service"
}