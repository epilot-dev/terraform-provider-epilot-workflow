terraform {
  required_providers {
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.2.0"
    }
  }
}

variable "epilot_auth" {
  type = string
}

provider "epilot-product" {
  # Configuration options
  epilot_auth = var.epilot_auth
}
