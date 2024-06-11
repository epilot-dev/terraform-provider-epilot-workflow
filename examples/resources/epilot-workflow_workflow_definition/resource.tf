terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "1.0.1"
    }

  }
}

provider "epilot-workflow" {
  # Configuration options
  bearer_auth = var.bearer_auth
}

variable "bearer_auth" {
  type =  string
}

resource "epilot-workflow_workflow_definition" "my_workflowdefinition" {

}