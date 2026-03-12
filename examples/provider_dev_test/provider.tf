terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.22.2"
    }
  }
}

provider "epilot-workflow" {
  bearer_auth = "eyJraWQiOiJ1SERLWlFKSnBGXC9cL2FmSjhEb0UrclhISFI4dFY2a0J1Q0k4VTZ3dE5sc0k9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIyMzk0YzgxMi04MGIxLTcwYzEtMGE0Ny1iZjFiZWVmNTg2OWIiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfWHRSalU3VkFKIiwiY3VzdG9tOml2eV9vcmdfaWQiOiI5MTE4NDUiLCJjb2duaXRvOnVzZXJuYW1lIjoiZC5ib3RAZXBpbG90LmNsb3VkIiwiY3VzdG9tOml2eV91c2VyX2lkIjoiMTAwMTA3MjgiLCJvcmlnaW5fanRpIjoiNzA4NzJiMDktMjM5NS00OTkwLWExN2YtNzE5NzRmYzYwNjBmIiwiYXVkIjoiNGo0ZnNpbjNncG8yZTRvMm9jYXBuNjdodGciLCJldmVudF9pZCI6IjUyZWEyN2JlLTE5MjQtNDVjOC04MTY3LWVhZGFhMDIxM2U1MSIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzczMzIyNDcxLCJleHAiOjE3NzMzMjY4NjYsImlhdCI6MTc3MzMyMzI2NiwianRpIjoiMmUwZjI2ZjQtODlhOS00YTA5LWE5Y2YtMTNiMWJhN2JlMDdiIiwiZW1haWwiOiJkLmJvdEBlcGlsb3QuY2xvdWQifQ.pColnfmNoQSJI4T77kgGdHS5Hv69Ezt4RqZEte7JG-XOeFPKleo6w9KYoP_gjALzSvYYOpgp3pe9916UIzyPiYQK3hh3zJ7gcS1Pr0WWdJYV-SFlxgRAhCVIMKN5qoNAq12_gcV95GG2ub5PTSDO6At8nGbXAjcNoG1M0QHh0tMqHBEyxfC8fKyl-w10qKQ0QdZhDoQN7tEBDKLDCxae7_f0EZVNC2AW0deiECsXNcx1Bh41YfEiAXx-0zH0Iwkjuc4jQCip4uKgX1NSvI_uc24F5ExlfHXhtOqV7CO4bYxsO6yogUC2WMRtPNdZUuweqG0CkJjuJwDHsq7Bz4iB7Q"
  server_url  = "https://workflows-definition.dev.sls.epilot.io/"                # Optional
}

# import {
#   to = epilot-workflow_flow_template.my_epilot-workflow_flow_template
#   id = "wf_db8OBtw"
# }

# resource "epilot-workflow_flow_template" "my_epilot-workflow_flow_template" {}