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

# import {
#   to = epilot-workflow_flow_template.my_epilot-workflow_flow_template
#   id = "wf_db8OBtw"
# }

# resource "epilot-workflow_flow_template" "my_epilot-workflow_flow_template" {}