terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.12.0"
    }
  }
}

variable "epilot_auth" {
  type = string
}

provider "epilot-workflow" {
  # Configuration options
  bearer_auth = var.epilot_auth
}

# resource "epilot-workflow_closing_reason" "my_closingreason" {

# }