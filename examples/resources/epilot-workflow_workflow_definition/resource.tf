resource "epilot-workflow_workflow_definition" "my_workflowdefinition" {
  creation_time       = "2021-04-27T12:01:13.000Z"
  definition_id       = "7hj28a"
  description         = "...my_description..."
  due_date            = "2021-04-27T12:00:00.000Z"
  enable_ecp_workflow = false
  flow = [
    {
      section = {
        id    = "452e3f02-980e-44cf-845d-81d428f0e4f5"
        name  = "Luther Greenfelder"
        order = 12.7
        steps = [
          {
            assigned_to = [
              "...",
            ]
            automation_config = {
              flow_id = "...my_flow_id..."
            }
            description = {
              enabled = false
              value   = "...my_value..."
            }
            due_date = "2021-04-27T12:00:00.000Z"
            dynamic_due_date = {
              action_type_condition = "WORKFLOW_STARTED"
              number_of_units       = 95.2
              step_id               = "...my_step_id..."
              time_period           = "weeks"
            }
            ecp = {
              description = "...my_description..."
              enabled     = true
              journey = {
                id         = "56e082d1-2997-4428-b0e6-c4cc2bdf6697"
                journey_id = "...my_journey_id..."
                name       = "Floyd Pacocha"
              }
              label = "...my_label..."
            }
            execution_type = "MANUAL"
            id             = "7a2ecac9-f13d-48a1-a6b9-dc969f95f0ea"
            installer = {
              description = "...my_description..."
              enabled     = true
              journey = {
                id         = "4100e0ec-f324-443f-89bc-c4a7284764a3"
                journey_id = "...my_journey_id..."
                name       = "Kara Hoppe"
              }
              label = "...my_label..."
            }
            journey = {
              id         = "eff3a06c-e7de-4f51-84a6-068ca5b78833"
              journey_id = "...my_journey_id..."
              name       = "Brandy Senger"
            }
            name  = "Sammy Gislason"
            order = 14.54
            requirements = [
              {
                condition     = "CLOSED"
                definition_id = "...my_definition_id..."
                type          = "STEP"
              },
            ]
            type = "STEP"
            user_ids = [
              94.55,
            ]
          },
        ]
        type = "SECTION"
      }
    },
  ]
  id               = "6301351b-2bbd-41bf-80cf-7a2f75db8f8a"
  last_update_time = "2021-04-27T12:01:13.000Z"
  name             = "Dr. Matthew Cronin"
}