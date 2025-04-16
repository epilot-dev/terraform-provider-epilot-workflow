resource "epilot-workflow_workflow_definition" "my_workflowdefinition" {
  assigned_to = [
    "..."
  ]
  closing_reasons = [
    {
      id = "x739cew"
    }
  ]
  creation_time = "2021-04-27T12:01:13.000Z"
  description   = "...my_description..."
  due_date      = "2021-04-27T12:00:00.000Z"
  dynamic_due_date = {
    action_type_condition = "PHASE_FINISHED"
    number_of_units       = 0.36
    phase_id              = "...my_phase_id..."
    step_id               = "...my_step_id..."
    time_period           = "days"
  }
  enable_ecp_workflow = false
  enabled             = false
  flow                = "{ \"see\": \"documentation\" }"
  id                  = "...my_id..."
  last_update_time    = "2021-04-27T12:01:13.000Z"
  name                = "...my_name..."
  taxonomies = [
    "..."
  ]
  update_entity_attributes = [
    {
      source = "current_step"
      target = {
        entity_attribute = "my_status"
        entity_schema    = "opportunity"
      }
    }
  ]
  user_ids = [
    5.93
  ]
}