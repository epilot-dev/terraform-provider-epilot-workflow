resource "epilot-workflow_flow_template" "my_flowtemplate" {
  assigned_to = [
    "..."
  ]
  available_in_ecp = true
  closing_reasons = [
    {
      id = "...my_id..."
    }
  ]
  created_at  = "2021-04-27T12:01:13.000Z"
  description = "...my_description..."
  due_date    = "2021-04-27T12:00:00.000Z"
  due_date_config = {
    duration = 0.19
    phase_id = "...my_phase_id..."
    task_id  = "...my_task_id..."
    type     = "PHASE_FINISHED"
    unit     = "hours"
  }
  edges = [
    {
      condition_id = "...my_condition_id..."
      from_id      = "...my_from_id..."
      id           = "...my_id..."
      none_met     = true
      to_id        = "...my_to_id..."
    }
  ]
  enabled = true
  entity_sync = [
    {
      target = {
        entity_attribute = "title"
        entity_schema    = "opportunity"
      }
      trigger = {
        event = "FlowDueDateChanged"
        filter = {
          phase_template_id = "...my_phase_template_id..."
          task_template_id  = "...my_task_template_id..."
        }
      }
      value = {
        source = "phase_status"
        value  = "...my_value..."
      }
    }
  ]
  id     = "...my_id..."
  name   = "...my_name..."
  org_id = "...my_org_id..."
  phases = [
    {
      assigned_to = [
        "..."
      ]
      due_date = "2021-04-27T12:00:00.000Z"
      due_date_config = {
        duration = 9.34
        phase_id = "...my_phase_id..."
        task_id  = "...my_task_id..."
        type     = "TASK_FINISHED"
        unit     = "days"
      }
      id   = "...my_id..."
      name = "...my_name..."
      taxonomies = [
        "..."
      ]
    }
  ]
  single_closing_reason_selection = true
  tasks = [
    {
      ai_agent_task = {
        agent_config = {
          additional_properties = "{ \"see\": \"documentation\" }"
          agent_id              = "...my_agent_id..."
        }
        assigned_to = [
          "..."
        ]
        description = {
          enabled = false
          value   = "...my_value..."
        }
        due_date = "2021-04-27T12:00:00.000Z"
        due_date_config = {
          duration = 3.61
          phase_id = "...my_phase_id..."
          task_id  = "...my_task_id..."
          type     = "PHASE_FINISHED"
          unit     = "months"
        }
        ecp = {
          description = "...my_description..."
          enabled     = false
          journey = {
            complete_task_automatically = true
            id                          = "...my_id..."
            journey_id                  = "...my_journey_id..."
            name                        = "...my_name..."
          }
          label = "...my_label..."
        }
        id = "...my_id..."
        installer = {
          description = "...my_description..."
          enabled     = true
          journey = {
            complete_task_automatically = false
            id                          = "...my_id..."
            journey_id                  = "...my_journey_id..."
            name                        = "...my_name..."
          }
          label = "...my_label..."
        }
        journey = {
          complete_task_automatically = false
          id                          = "...my_id..."
          journey_id                  = "...my_journey_id..."
          name                        = "...my_name..."
        }
        name     = "...my_name..."
        phase_id = "...my_phase_id..."
        requirements = [
          {
            phase_id = "...my_phase_id..."
            task_id  = "...my_task_id..."
            when     = "TASK_FINISHED"
          }
        ]
        task_type = "MANUAL"
        taxonomies = [
          "..."
        ]
      }
    }
  ]
  taxonomies = [
    "..."
  ]
  trigger = {
    journey_submission_trigger = {
      automation_id = "...my_automation_id..."
      id            = "...my_id..."
      journey_id    = "...my_journey_id..."
      journey_name  = "...my_journey_name..."
      type          = "journey_submission"
    }
  }
  updated_at = "2021-04-27T12:01:13.000Z"
  version    = "...my_version..."
}