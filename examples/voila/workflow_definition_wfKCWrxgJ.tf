# Provider definitions
terraform {
  required_providers {
    epilot-journey = {
      source  = "epilot-dev/epilot-journey"
      version = "1.0.1"
      
    }
    epilot-product = {
      source  = "epilot-dev/epilot-product"
      version = "0.6.2"
    }
    epilot-automation = {
      source  = "epilot-dev/epilot-automation"
      version = "2.0.1"
    }
    epilot-entitymapping = {
      source  = "epilot-dev/epilot-entitymapping"
      version = "1.6.7"
    }
    epilot-designbuilder = {
      source  = "epilot-dev/epilot-designbuilder"
      version = "0.0.5"
    }
    epilot-file = {
      source  = "epilot-dev/epilot-file"
      version = "2.2.3"
    }
    epilot-emailtemplate = {
      source  = "epilot-dev/epilot-emailtemplate"
      version = "0.0.2"
    }
    epilot-schema = {
      source  = "epilot-dev/epilot-schema"
      version = "4.0.0"
    }
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.2.3"
    }
  }
}

# Variables
variable "epilot_auth" {
  type = string
}
variable "journey_api_url" {
  type = string
  default = "https://journey-config.sls.epilot.io"
}
variable "product_api_url" {
  type = string
  default = "https://product.sls.epilot.io"
}
variable "automation_api_url" {
  type = string
  default = "https://automation.sls.epilot.io"
}
variable "file_api_url" {
  type = string
  default = "https://file.sls.epilot.io"
}
variable "emailtemplate_api_url" {
  type = string
  default = "https://email-template.sls.epilot.io"
}
variable "designbuilder_api_url" {
  type = string
  default = "https://design-builder-api.sls.epilot.io"
}
variable "entitymapping_api_url" {
  type = string
  default = "https://entity-mapping.sls.epilot.io"
}
variable "schema_api_url" {
  type = string
  default = "https://entity.sls.epilot.io"
}
variable "workflow_api_url" {
  type = string
  default = "https://workflows-definition.sls.epilot.io"
}

# Providers configuration
provider "epilot-journey" {
  epilot_auth = var.epilot_auth
  server_url = var.journey_api_url
}
provider "epilot-product" {
  epilot_auth = var.epilot_auth
  server_url = var.product_api_url
}
provider "epilot-automation" {
  epilot_auth = var.epilot_auth
  server_url = var.automation_api_url
}
provider "epilot-file" {
  epilot_auth = var.epilot_auth
  server_url = var.file_api_url
}
provider "epilot-emailtemplate" {
  epilot_auth = var.epilot_auth
  server_url = var.emailtemplate_api_url
}
provider "epilot-designbuilder" {
  custom_authorizer = var.epilot_auth
  server_url = var.designbuilder_api_url
}
provider "epilot-entitymapping" {
  epilot_auth = var.epilot_auth
  server_url = var.entitymapping_api_url
}
provider "epilot-schema" {
  epilot_auth = var.epilot_auth
  server_url = var.schema_api_url
}
provider "epilot-workflow" {
  bearer_auth = var.epilot_auth
  server_url = var.workflow_api_url
}

# epilot-workflow_workflow_definition.workflow_definition_wfKCWrxgJ:
resource "epilot-workflow_workflow_definition" "workflow_definition_wfKCWrxgJ" {
  assigned_to = [
    "11001018",
    "11000566",
    "10002602",
    "911676_11001026",
    "911677_11001027",
    "placeholder_ecp_customers",
  ]
  closing_reasons = [
    {
      id = "qgK9sGbKoS7NxlAbNReVn"
    },
    {
      id = "1gNTrk7Kx50g7EKGPc5GR"
    },
  ]

  description         = "Brew a nice cup of coffee :) "
  enable_ecp_workflow = false
  flow = [
    {
      step = {
        assigned_to = [
          "11001018",
          "10002602",
          "911674_11001023",
          "placeholder_ecp_customers",
        ]
        automation_config = {
          flow_id = "7262dcc1-0352-4bfc-9981-a2b8797c2d32"
        }
        description = {
          enabled = true
          value   = "Ladida"
        }
        due_date = "2024-06-19T15:29:00.000Z"
        ecp = {
          enabled = false
        }
        execution_type = "AUTOMATION"

        installer = {
          enabled = false
        }
        journey = {
          id         = "49c36269-b6f2-49d6-82d2-050328899e79"
          journey_id = "f9c7bf40-2f19-11ef-ab16-6546f1cba0b0"
          name       = "public data valiation"
        }
        name         = "Choose high quality coffee beans"
        order        = 1
        requirements = []
        type         = "STEP"
        user_ids     = []
      }
    },
    {
      section = {

        name  = "Pre-steps"
        order = 2
        steps = [
          {
            assigned_to = []
            description = {
              enabled = true
              value   = "The standard ratio is about 1 to 2 tablespoons of coffee per 6 ounces of water. Adjust according to your taste preferences. For a standard 12-cup coffee maker, you’ll need about 12-24 tablespoons (3/4 to 1 1/2 cups) of coffee."
            }
            dynamic_due_date = {
              action_type_condition = null
              number_of_units       = 0
              time_period           = null
            }
            ecp = {
              enabled = false
            }

            installer = {
              enabled = false
            }
            name         = "Measure the right amounts of coffee and water "
            order        = 3
            requirements = []
            type         = "STEP"
            user_ids     = []
          },
          {
            assigned_to = []
            description = {
              enabled = true
              value   = "If using whole beans, grind them to a medium grind, similar to the texture of sand. This is ideal for drip coffee makers. Too fine a grind can lead to over-extraction and a bitter taste, while too coarse can lead to under-extraction and a weak flavor."
            }
            due_date = "2024-05-23T20:11:00.000Z"
            ecp = {
              enabled = false
            }

            installer = {
              enabled = false
            }
            name  = "Grind the coffee beans"
            order = 4
            requirements = [
              {
                condition     = "CLOSED"
                definition_id = "3bF4nboZ"
                type          = "STEP"
              },
            ]
            type     = "STEP"
            user_ids = []
          },
        ]
        type = "SECTION"
      }
    },
    {
      section = {

        name  = "Prepare the coffee"
        order = 5
        steps = [
          {
            assigned_to = []
            ecp = {
              enabled = false
            }

            installer = {
              enabled = false
            }
            name  = "Add coffee grounds into the filter"
            order = 6
            requirements = [
              {
                condition     = "CLOSED"
                definition_id = "3bF4nboZ"
                type          = "STEP"
              },
              {
                condition     = "CLOSED"
                definition_id = "56z3PpJ-"
                type          = "STEP"
              },
            ]
            type     = "STEP"
            user_ids = []
          },
          {
            assigned_to = []
            description = {
              enabled = true
              value   = "Pour the measured amount of filtered water into the coffee maker’s water reservoir"
            }
            ecp = {
              enabled = false
            }

            installer = {
              enabled = false
            }
            name  = "Add water"
            order = 7
            requirements = [
              {
                condition     = "CLOSED"
                definition_id = "lLr14EM1"
                type          = "STEP"
              },
            ]
            type     = "STEP"
            user_ids = []
          },
          {
            assigned_to = []
            dynamic_due_date = {
              action_type_condition = null
              number_of_units       = 0
              time_period           = null
            }
            ecp = {
              enabled = false
            }

            installer = {
              enabled = false
            }
            name  = "Start brewing"
            order = 8
            requirements = [
              {
                condition     = "CLOSED"
                definition_id = "3bF4nboZ"
                type          = "STEP"
              },
              {
                condition     = "CLOSED"
                definition_id = "lLr14EM1"
                type          = "STEP"
              },
              {
                condition     = "CLOSED"
                definition_id = "56z3PpJ-"
                type          = "STEP"
              },
              {
                condition     = "CLOSED"
                definition_id = "Kpo2TBM0"
                type          = "STEP"
              },
              {
                condition     = "CLOSED"
                definition_id = "hr__AP3D"
                type          = "STEP"
              },
            ]
            type     = "STEP"
            user_ids = []
          },
        ]
        type = "SECTION"
      }
    },
    {
      step = {
        assigned_to = []
        ecp = {
          enabled = false
        }

        installer = {
          enabled = false
        }
        name         = "Serve fresh"
        order        = 9
        requirements = []
        type         = "STEP"
        user_ids     = []
      }
    },
    {
      section = {

        name  = "A new phase"
        order = 10
        steps = [
          {
            assigned_to = []

            name         = "Something else"
            order        = 11
            requirements = []
            type         = "STEP"
            user_ids     = []
          },
        ]
        type = "SECTION"
      }
    },
    {
      step = {
        assigned_to = []
        automation_config = {
          flow_id = "378d2cad-91b9-4161-9fda-fc3df1908490"
        }
        ecp = {
          enabled = false
        }
        execution_type = "AUTOMATION"

        installer = {
          enabled = false
        }
        name         = "Enjoy"
        order        = 12
        requirements = []
        type         = "STEP"
        user_ids     = []
      }
    },
  ]


  name = "Brew a cup of coffee"
  update_entity_attributes = [
    {
      source = "workflow_status"
      target = {
        entity_attribute = "status"
        entity_schema    = "opportunity"
      }
    },
  ]
  user_ids = []
}
