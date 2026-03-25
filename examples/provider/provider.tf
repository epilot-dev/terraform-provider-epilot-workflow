terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.28.17"
    }
  }
}

provider "epilot-workflow" {
  bearer_auth = "<YOUR_BEARER_AUTH>" # Required
  server_url  = "..."                # Optional
}