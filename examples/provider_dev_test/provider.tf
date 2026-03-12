terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.22.2"
    }
  }
}

provider "epilot-workflow" {
  bearer_auth = ""
  server_url  = "https://workflows-definition.dev.sls.epilot.io/"                # Optional
}
