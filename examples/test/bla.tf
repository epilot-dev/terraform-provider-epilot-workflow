terraform {
  required_providers {
    # epilot-journey = {
    #   source  = "epilot-dev/epilot-journey"
    #   version = "0.3.0"
    # }
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.6.0"
    }
  }
}

variable "epilot_auth" {
  type = string
}

# provider "epilot-journey" {
#   # Configuration options
#   epilot_auth = var.epilot_auth
# }

provider "epilot-product" {
  # Configuration options
  epilot_auth = var.epilot_auth
}

# import {
#   to = epilot-product_product.s_product
#   id = "103f1186-7f7c-42d1-8891-00868a589cf3"
# # }


# resource "epilot-product_product" "s_product" {
# }



# resource "epilot-product_tax" "s_tax" {
#   # (resource arguments)
# }