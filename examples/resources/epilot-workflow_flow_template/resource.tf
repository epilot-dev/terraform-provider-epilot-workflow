resource "epilot-workflow_flow_template" "my_flowtemplate" {
  assigned_to = [
    "..."
  ]
  available_in_ecp = true
  closing_reasons = [
    {
      id = "x739cew"
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
  enabled          = true
  id               = "...my_id..."
  is_flow_migrated = true
  name             = "...my_name..."
  org_id           = "...my_org_id..."
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
  tasks = [
    {
      automation_task = {
        assigned_to = [
          "..."
        ]
        automation_config = {
          flow_id = "...my_flow_id..."
        }
        description = {
          enabled = false
          value   = "...my_value..."
        }
        due_date = "2021-04-27T12:00:00.000Z"
        due_date_config = {
          duration = 8.65
          phase_id = "...my_phase_id..."
          task_id  = "...my_task_id..."
          type     = "PHASE_FINISHED"
          unit     = "days"
        }
        ecp = {
          description = "...my_description..."
          enabled     = true
          journey = {
            id         = "...my_id..."
            journey_id = "...my_journey_id..."
            name       = "...my_name..."
          }
          label = "...my_label..."
        }
        id = "...my_id..."
        installer = {
          description = "...my_description..."
          enabled     = false
          journey = {
            id         = "...my_id..."
            journey_id = "...my_journey_id..."
            name       = "...my_name..."
          }
          label = "...my_label..."
        }
        journey = {
          id         = "...my_id..."
          journey_id = "...my_journey_id..."
          name       = "...my_name..."
        }
        name     = "...my_name..."
        phase_id = "...my_phase_id..."
        requirements = [
          {
            phase_id = "...my_phase_id..."
            task_id  = "...my_task_id..."
            when     = "PHASE_FINISHED"
          }
        ]
        schedule = {
          delayed_schedule = {
            duration = 9.44
            mode     = "delayed"
            unit     = "months"
          }
          immediate_schedule = {
            mode = "immediate"
          }
          relative_schedule = {
            direction = "before"
            duration  = 3.2
            mode      = "relative"
            reference = {
              attribute = "...my_attribute..."
              id        = "...my_id..."
              origin    = "flow_started"
              schema    = "...my_schema..."
            }
            unit = "days"
          }
        }
        task_type = "AUTOMATION"
        taxonomies = [
          "..."
        ]
        trigger_mode = "automatic"
      }
      decision_task = {
        assigned_to = [
          "..."
        ]
        conditions = [
          {
            branch_name      = "Branch 1"
            id               = "...my_id..."
            logical_operator = "OR"
            statements = [
              {
                id       = "...my_id..."
                operator = "is_empty"
                source = {
                  attribute            = "...my_attribute..."
                  attribute_operation  = "deleted"
                  attribute_repeatable = false
                  attribute_type       = "boolean"
                  id                   = "...my_id..."
                  origin               = "action"
                  origin_type          = "journey_block"
                  schema               = "...my_schema..."
                }
                values = [
                  "..."
                ]
              }
            ]
          }
        ]
        description = {
          enabled = false
          value   = "...my_value..."
        }
        due_date = "2021-04-27T12:00:00.000Z"
        due_date_config = {
          duration = 8.34
          phase_id = "...my_phase_id..."
          task_id  = "...my_task_id..."
          type     = "WORKFLOW_STARTED"
          unit     = "days"
        }
        ecp = {
          description = "...my_description..."
          enabled     = true
          journey = {
            id         = "...my_id..."
            journey_id = "...my_journey_id..."
            name       = "...my_name..."
          }
          label = "...my_label..."
        }
        id = "...my_id..."
        installer = {
          description = "...my_description..."
          enabled     = true
          journey = {
            id         = "...my_id..."
            journey_id = "...my_journey_id..."
            name       = "...my_name..."
          }
          label = "...my_label..."
        }
        journey = {
          id         = "...my_id..."
          journey_id = "...my_journey_id..."
          name       = "...my_name..."
        }
        name     = "...my_name..."
        phase_id = "...my_phase_id..."
        requirements = [
          {
            phase_id = "...my_phase_id..."
            task_id  = "...my_task_id..."
            when     = "PHASE_FINISHED"
          }
        ]
        schedule = {
          delayed_schedule = {
            duration = 0.87
            mode     = "delayed"
            unit     = "days"
          }
          relative_schedule = {
            direction = "after"
            duration  = 7.2
            mode      = "relative"
            reference = {
              attribute = "...my_attribute..."
              id        = "...my_id..."
              origin    = "flow_started"
              schema    = "...my_schema..."
            }
            unit = "weeks"
          }
        }
        task_type = "MANUAL"
        taxonomies = [
          "..."
        ]
      }
      task_base = {
        assigned_to = [
          "..."
        ]
        description = {
          enabled = true
          value   = "...my_value..."
        }
        due_date = "2021-04-27T12:00:00.000Z"
        due_date_config = {
          duration = 9.32
          phase_id = "...my_phase_id..."
          task_id  = "...my_task_id..."
          type     = "TASK_FINISHED"
          unit     = "hours"
        }
        ecp = {
          description = "...my_description..."
          enabled     = true
          journey = {
            id         = "...my_id..."
            journey_id = "...my_journey_id..."
            name       = "...my_name..."
          }
          label = "...my_label..."
        }
        id = "...my_id..."
        installer = {
          description = "...my_description..."
          enabled     = true
          journey = {
            id         = "...my_id..."
            journey_id = "...my_journey_id..."
            name       = "...my_name..."
          }
          label = "...my_label..."
        }
        journey = {
          id         = "...my_id..."
          journey_id = "...my_journey_id..."
          name       = "...my_name..."
        }
        name     = "...my_name..."
        phase_id = "...my_phase_id..."
        requirements = [
          {
            phase_id = "...my_phase_id..."
            task_id  = "...my_task_id..."
            when     = "PHASE_FINISHED"
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
    automation_trigger = {
      automation_id = "...my_automation_id..."
      id            = "...my_id..."
      type          = "automation"
    }
    journey_submission_trigger = {
      automation_id = "...my_automation_id..."
      id            = "...my_id..."
      journey_id    = "...my_journey_id..."
      type          = "journey_submission"
    }
    manual_trigger = {
      entity_schema = "...my_entity_schema..."
      id            = "...my_id..."
      type          = "manual"
    }
  }
  update_entity_attributes = [
    {
      source = "current_step"
      target = {
        entity_attribute = "my_status"
        entity_schema    = "opportunity"
      }
    }
  ]
  updated_at = "2021-04-27T12:01:13.000Z"
}