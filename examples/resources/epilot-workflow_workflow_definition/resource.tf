resource "epilot-workflow_workflow_definition" "my_workflowdefinition" {
  creation_time       = "2021-04-27T12:01:13.000Z"
  definition_id       = "7hj28a"
  description         = "...my_description..."
  due_date            = "2021-04-27T12:00:00.000Z"
  enable_ecp_workflow = false
  flow = [
    {
      section = {
        id    = "2a31ed4e-b201-4187-9740-15a348b64814"
        name  = "Meredith Bruen"
        order = 1.95
        steps = [
          {
            assigned_to = [
              "...",
            ]
            automation_config = {
              flow_id = "...my_flow_id..."
            }
            description = {
              enabled = true
              value   = "...my_value..."
            }
            due_date = "2021-04-27T12:00:00.000Z"
            dynamic_due_date = {
              action_type_condition = "STEP_CLOSED"
              number_of_units       = 73.8
              step_id               = "...my_step_id..."
              time_period           = "weeks"
            }
            ecp = {
              description = "...my_description..."
              enabled     = false
              journey = {
                id         = "64361b27-76d7-498c-b5f3-e120ea0dc9e1"
                journey_id = "...my_journey_id..."
                name       = "Lorraine Witting"
              }
              label = "...my_label..."
            }
            execution_type = "AUTOMATION"
            id             = "3f96aa0e-3408-45f7-9b11-76521d792fea"
            installer = {
              description = "...my_description..."
              enabled     = true
              journey = {
                id         = "830af35a-c0f0-4516-986c-6d56674fb26d"
                journey_id = "...my_journey_id..."
                name       = "Paulette O'Hara"
              }
              label = "...my_label..."
            }
            journey = {
              id         = "9b2b84e8-8293-4d18-a93f-16fc9936c940"
              journey_id = "...my_journey_id..."
              name       = "Dr. Joel Hermann"
            }
            name  = "Lorene Denesik"
            order = 4.83
            requirements = [
              {
                condition     = "CLOSED"
                definition_id = "...my_definition_id..."
                type          = "STEP"
              },
            ]
            type = "SECTION"
            user_ids = [
              9.27,
            ]
          },
        ]
        type = "STEP"
      }
    },
  ]
  id               = "70408a19-3e72-4165-b76b-6381b5dd263e"
  last_update_time = "2021-04-27T12:01:13.000Z"
  name             = "Nina Bednar"
}