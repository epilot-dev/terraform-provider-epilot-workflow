terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.20.1"
    }
  }
}


provider "epilot-workflow" {
  # Configuration options
  bearer_auth= "..."
}

