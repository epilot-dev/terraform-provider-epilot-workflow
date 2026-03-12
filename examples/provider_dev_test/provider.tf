terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.22.2"
    }
  }
}

provider "epilot-workflow" {
  bearer_auth = "eyJraWQiOiJ1SERLWlFKSnBGXC9cL2FmSjhEb0UrclhISFI4dFY2a0J1Q0k4VTZ3dE5sc0k9IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIyMzk0YzgxMi04MGIxLTcwYzEtMGE0Ny1iZjFiZWVmNTg2OWIiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfWHRSalU3VkFKIiwiY3VzdG9tOml2eV9vcmdfaWQiOiI5MTE4NDUiLCJjb2duaXRvOnVzZXJuYW1lIjoiZC5ib3RAZXBpbG90LmNsb3VkIiwiY3VzdG9tOml2eV91c2VyX2lkIjoiMTAwMTA3MjgiLCJvcmlnaW5fanRpIjoiZTAyMzM4MWUtYjRjZC00YTQ3LTlhOTctYzc2NTU2MTU2MzRhIiwiYXVkIjoiNGo0ZnNpbjNncG8yZTRvMm9jYXBuNjdodGciLCJldmVudF9pZCI6ImU0Nzk3NmQ3LTYwNGEtNDFlNS04Y2U3LTE1ZmE2NDc2MWQxMyIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzczMzExODMxLCJleHAiOjE3NzMzMTU0MzEsImlhdCI6MTc3MzMxMTgzMSwianRpIjoiM2Y0NmM0ZGEtODhlYy00N2VlLWFkN2QtMDMxNjkxYzExMGM0IiwiZW1haWwiOiJkLmJvdEBlcGlsb3QuY2xvdWQifQ.S0en0E6gkYgsG90E9hOreM2RAmMJsE2xnWa1VpIYt4dbuHmHwxxIhfXqdgchjj7Fbp2JusylcTmxjJXMsgR1ryLJyRaZYK0OZli6Kp_D0hIGQrxJThf31ZQ2fBchuFX7EOpvr2Gn8z8ePJRKhEP99j29qt9uPYr9wrWk9lBusfiaJrAxOq69po57h7P-i5uRrlOURw-diB3Qg385PmxQmIW0WwJqHFClrpkmtwycSXRO14tFopo03Tm6jweqmcACSpBjxAKIlukShmEWrw3fcaE0Mlrd0UUjhd_fyo_NSZ4dm_oGj-Z3iYz44H6L7xenleuvtvpKEuxxKnPNJwwSNw"
  server_url  = "https://workflows-definition.dev.sls.epilot.io/"                # Optional
}

# import {
#   to = epilot-workflow_flow_template.my_epilot-workflow_flow_template
#   id = "wf_db8OBtw"
# }

# resource "epilot-workflow_flow_template" "my_epilot-workflow_flow_template" {}