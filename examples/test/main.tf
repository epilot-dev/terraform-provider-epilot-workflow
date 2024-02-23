# epilot-product_product.s_product:
resource "epilot-product_product" "s_product" {
    acl           = {
        delete = [
            "org_66",
        ]
        edit   = [
            "org_66",
        ]
        view   = [
            "org_66",
        ]
    }
    active        = true
    created_at    = "2024-02-07T14:44:33.799Z"
    id            = "103f1186-7f7c-42d1-8891-00868a589cf3"
    name          = "Product PH - Heating"
    org           = "66"
    owners        = [
        {
            org_id  = "66"
            user_id = "77701"
        },
    ]
    price_options = {
        dollar_relation = [
            {
                entity_id = "cbcf02eb-bfca-4317-912b-7f4db997a739"
            },
        ]
    }
    schema        = "product"
    title         = "Product PH - Heating"
    type          = "product"
    updated_at    = "2024-02-07T14:44:33.799Z"
}
