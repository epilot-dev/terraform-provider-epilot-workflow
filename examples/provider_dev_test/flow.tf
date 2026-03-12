# epilot-workflow_flow_template.my_epilot-workflow_flow_template:
resource "epilot-workflow_flow_template" "my_epilot-workflow_flow_template" {
    additional_triggers = []
    assigned_to         = []
    available_in_ecp    = true
    closing_reasons     = [
        {
            id               = "qgK9sGbKoS7NxlAbNReVn"
        },
        {
            id               = "1gNTrk7Kx50g7EKGPc5GR"
        },
    ]
    description         = "Testing blueprint migration heyo"
    due_date_config     = {
        duration = 1
        type     = "WORKFLOW_STARTED"
        unit     = "months"
    }
    edges               = [
        {
            from_id = "c03fb06f-3917-467a-aa04-44f9c68eceda"
            id      = "118b81be-5726-4401-9938-8e99d1ac55e5"
            to_id   = "86e6ccf4-27bd-4468-a804-a6ec25585ea6"
        },
        {
            from_id = "86e6ccf4-27bd-4468-a804-a6ec25585ea6"
            id      = "060da883-3a44-4831-a4e7-e9c84cdcda30"
            to_id   = "55f72bb7-e059-44d7-8d87-687d5979a199"
        },
        {
            condition_id = "5ba9fff1-bd40-42a8-99eb-2dca43172f3e"
            from_id      = "55f72bb7-e059-44d7-8d87-687d5979a199"
            id           = "loop-edge-6eeb812d-8431-401d-a9fc-9a8648263d3f"
            to_id        = "86e6ccf4-27bd-4468-a804-a6ec25585ea6"
        },
        {
            condition_id = "5ca3b628-4f0a-4ccf-924d-0c9d38a1c417"
            from_id      = "55f72bb7-e059-44d7-8d87-687d5979a199"
            id           = "366de8d8-940c-472c-8b2d-4ddc1c1da2ec"
            to_id        = "e828668c-a329-4a68-9697-6fcf5d8f6485"
        },
        {
            condition_id = "exit-loop-3a6875c365cb4881ba0eb526d7bf390e"
            from_id      = "55f72bb7-e059-44d7-8d87-687d5979a199"
            id           = "4ba4ab83-c1bb-44b8-93e0-23374dade230"
            to_id        = "8527b440-e0f4-4ee0-9b55-9471f64c4a11"
        },
        {
            condition_id = "none-met-63763665-4944-4d10-bb7b-7b93f20c69e2"
            from_id      = "55f72bb7-e059-44d7-8d87-687d5979a199"
            id           = "7ac702f0-f372-4e55-a4d1-7d3d2058db91"
            none_met     = true
            to_id        = "85bfc25c-01c4-42d6-b345-a0796e3a3d98"
        },
        {
            from_id = "c03fb06f-3917-467a-aa04-44f9c68eceda"
            id      = "b0e01e54-ca3c-4d06-bbdd-a3ecfe45066e"
            to_id   = "35ef3c7c-8529-4d23-9dd2-529cabdfc7cb"
        },
        {
            from_id = "35ef3c7c-8529-4d23-9dd2-529cabdfc7cb"
            id      = "4083ca2a-6d15-4ef6-933d-30151c9251a6"
            to_id   = "6bf6e79b-6d26-4d0c-ba4f-56090c500652"
        },
        {
            from_id = "6bf6e79b-6d26-4d0c-ba4f-56090c500652"
            id      = "978a875a-07d0-4ac1-a294-6e04080b5bd9"
            to_id   = "702f0cd4-e425-4939-ab4a-f9a13c8c06bc"
        },
        {
            condition_id = "none-met-e1ca4fa1-b2ed-4863-9406-c66d4ec864c8"
            from_id      = "702f0cd4-e425-4939-ab4a-f9a13c8c06bc"
            id           = "0be4f4c3-9091-4f86-a2c7-67068070d1c3"
            none_met     = true
            to_id        = "86b1b03f-495e-43ab-9da3-f9ed2d5de1d9"
        },
        {
            condition_id = "38b4d4bd-f7da-4961-a840-faf35d6834e8"
            from_id      = "702f0cd4-e425-4939-ab4a-f9a13c8c06bc"
            id           = "a4505ede-0f58-4c61-817f-ddda1ae08468"
            to_id        = ""
        },
        {
            from_id = "86b1b03f-495e-43ab-9da3-f9ed2d5de1d9"
            id      = "6ae16905-eed6-4f9d-90e5-9fc235f543e6"
            to_id   = ""
        },
        {
            from_id = "trigger"
            id      = "1571e080-b707-4f92-bef4-fe62bd14136f"
            to_id   = "9639e6df-2f17-4c9a-99e7-94aa96a3f441"
        },
        {
            from_id = "9639e6df-2f17-4c9a-99e7-94aa96a3f441"
            id      = "0a586a27-5a3a-45a6-9660-9405db7e38f5"
            to_id   = "c03fb06f-3917-467a-aa04-44f9c68eceda"
        },
    ]
    enabled             = true
    entity_sync         = [
        {
            target  = {
                entity_attribute = "virt_test_1"
                entity_schema    = "opportunity"
            }
            trigger = {
                event = "FlowStarted"
            }
            value   = {
                source = "workflow_status"
            }
        },
    ]
    name                = "Daniel Test Blueprint"
    phases              = [
        {
            assigned_to     = []
            due_date_config = {
                duration = 2
                type     = "WORKFLOW_STARTED"
                unit     = "hours"
            }
            id              = "c03fb06f-3917-467a-aa04-44f9c68eceda"
            name            = "Installation"
            taxonomies      = [
                "baf48dcc-d7bb-44ae-8d0d-284495dedcec",
            ]
        },
        {
            assigned_to = []
            id          = "6bf6e79b-6d26-4d0c-ba4f-56090c500652"
            name        = "Validation"
            taxonomies  = [
                "33158105-16da-4c0d-be37-bec5d49def90",
            ]
        },
    ]
    tasks               = [
        {
            task_base = {
                assigned_to  = []
                description  = {
                    enabled = true
                    value   = "<body><ul __lexicallisttype=\"check\"><li role=\"checkbox\" tabindex=\"-1\" aria-checked=\"false\" value=\"1\" class=\"listitem listitemUnchecked\" data-epilot-style=\"\" style=\"\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">Do this</span></li><li role=\"checkbox\" tabindex=\"-1\" aria-checked=\"false\" value=\"2\" class=\"listitem listitemUnchecked\" style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\" data-epilot-style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">Do that</span></li></ul></body>"
                }
                ecp          = {
                    description = "<body><ul><li value=\"1\" class=\"listitem\" data-epilot-style=\"\" style=\"\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">1</span></li><li value=\"2\" class=\"listitem\" style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\" data-epilot-style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">2</span></li><li value=\"3\" class=\"listitem\" style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\" data-epilot-style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">3</span></li></ul></body>"
                    enabled     = true
                    label       = "My goodness"
                }
                id           = "86e6ccf4-27bd-4468-a804-a6ec25585ea6"
                journey      = {
                    complete_task_automatically = true
                    id                          = "acae4f15-bdf0-4b01-9a6e-ad64271e3f1f"
                    journey_id                  = "ea59c7f0-ecba-11f0-b478-6d6b2bd8039e"
                    name                        = "Daniel Test NEW NEW 2026-01"
                }
                name         = "Task w/ Journey"
                phase_id     = "c03fb06f-3917-467a-aa04-44f9c68eceda"
                requirements = []
                task_type    = "MANUAL"
                taxonomies   = []
            }
        },
        {
            decision_task = {
                assigned_to  = []
                conditions   = [
                    {
                        branch_name      = "East"
                        id               = "5ba9fff1-bd40-42a8-99eb-2dca43172f3e"
                        logical_operator = "AND"
                        statements       = []
                    },
                    {
                        branch_name      = "Nowhere"
                        id               = "none-met-63763665-4944-4d10-bb7b-7b93f20c69e2"
                        logical_operator = "AND"
                        statements       = []
                    },
                    {
                        branch_name      = "West"
                        id               = "5ca3b628-4f0a-4ccf-924d-0c9d38a1c417"
                        logical_operator = "AND"
                        statements       = []
                    },
                    {
                        branch_name      = "Exit Loop"
                        id               = "exit-loop-3a6875c365cb4881ba0eb526d7bf390e"
                        logical_operator = "AND"
                        statements       = []
                    },
                ]
                id           = "55f72bb7-e059-44d7-8d87-687d5979a199"
                loop_config  = {
                    exit_branch_id = "exit-loop-3a6875c365cb4881ba0eb526d7bf390e"
                    loop_branch_id = "5ba9fff1-bd40-42a8-99eb-2dca43172f3e"
                    max_iterations = 3
                }
                name         = "Where now?"
                phase_id     = "c03fb06f-3917-467a-aa04-44f9c68eceda"
                requirements = []
                task_type    = "DECISION"
                taxonomies   = []
                trigger_mode = "manual"
            }
        },
        {
            automation_task = {
                assigned_to           = []
                automation_config     = {
                    flow_id = "35d9a5bb-773c-4d45-9677-35d3f3049547"
                }
                created_automatically = false
                id                    = "e828668c-a329-4a68-9697-6fcf5d8f6485"
                name                  = "Webhook"
                phase_id              = "c03fb06f-3917-467a-aa04-44f9c68eceda"
                requirements          = []
                schedule              = {
                    delayed_schedule = {
                        duration = 2
                        mode     = "delayed"
                        unit     = "minutes"
                    }
                }
                task_type             = "AUTOMATION"
                taxonomies            = []
                trigger_mode          = "automatic"
            }
        },
        {
            automation_task = {
                assigned_to           = []
                automation_config     = {
                    flow_id = "ed5e8e37-ae2e-42e3-8c89-6555bb38bb90"
                }
                created_automatically = false
                id                    = "8527b440-e0f4-4ee0-9b55-9471f64c4a11"
                name                  = "Close Wf"
                phase_id              = "c03fb06f-3917-467a-aa04-44f9c68eceda"
                requirements          = []
                schedule              = {
                    immediate_schedule = {
                        mode = "immediate"
                    }
                }
                task_type             = "AUTOMATION"
                taxonomies            = []
                trigger_mode          = "automatic"
            }
        },
        {
            automation_task = {
                assigned_to           = []
                automation_config     = {
                    flow_id = "b698e432-59ac-49d4-ad9f-6d78117a5534"
                }
                created_automatically = false
                id                    = "85bfc25c-01c4-42d6-b345-a0796e3a3d98"
                name                  = "Start another Workflow"
                phase_id              = "c03fb06f-3917-467a-aa04-44f9c68eceda"
                requirements          = []
                task_type             = "AUTOMATION"
                taxonomies            = []
                trigger_mode          = "manual"
            }
        },
        {
            automation_task = {
                assigned_to           = []
                automation_config     = {
                    flow_id = "b0ddb290-87b1-4414-9c24-496ce9c1c0e1"
                }
                created_automatically = false
                id                    = "35ef3c7c-8529-4d23-9dd2-529cabdfc7cb"
                name                  = "Send Email"
                requirements          = []
                schedule              = {
                    immediate_schedule = {
                        mode = "immediate"
                    }
                }
                task_type             = "AUTOMATION"
                taxonomies            = []
                trigger_mode          = "automatic"
            }
        },
        {
            decision_task = {
                assigned_to  = []
                conditions   = [
                    {
                        branch_name      = "None met"
                        id               = "none-met-e1ca4fa1-b2ed-4863-9406-c66d4ec864c8"
                        logical_operator = "AND"
                        statements       = []
                    },
                    {
                        branch_name      = "Nowhere"
                        id               = "38b4d4bd-f7da-4961-a840-faf35d6834e8"
                        logical_operator = "AND"
                        statements       = [
                            {
                                id       = "87f3f88a-e411-406e-ad01-6f9ed2abcdd9"
                                operator = "equals"
                                source   = {
                                    attribute            = "modem_installationsort"
                                    attribute_repeatable = false
                                    attribute_type       = "string"
                                    id                   = "82d4e9a5-8a6f-4a39-b36c-fab18ea96baf"
                                    origin               = "trigger"
                                    schema               = "opportunity"
                                }
                                values   = [
                                    "nowhere",
                                ]
                            },
                        ]
                    },
                ]
                id           = "702f0cd4-e425-4939-ab4a-f9a13c8c06bc"
                name         = "Auto Condition"
                phase_id     = "6bf6e79b-6d26-4d0c-ba4f-56090c500652"
                requirements = []
                schedule     = {
                    delayed_schedule = {
                        duration = 1
                        mode     = "delayed"
                        unit     = "minutes"
                    }
                }
                task_type    = "DECISION"
                taxonomies   = []
                trigger_mode = "automatic"
            }
        },
        {
            automation_task = {
                assigned_to           = []
                automation_config     = {
                    flow_id = "60ee0b95-400b-4124-8bd7-82ff8c3a860a"
                }
                created_automatically = false
                id                    = "86b1b03f-495e-43ab-9da3-f9ed2d5de1d9"
                name                  = "Map Data"
                phase_id              = "6bf6e79b-6d26-4d0c-ba4f-56090c500652"
                requirements          = []
                schedule              = {
                    immediate_schedule = {
                        mode = "immediate"
                    }
                }
                task_type             = "AUTOMATION"
                taxonomies            = []
                trigger_mode          = "automatic"
            }
        },
        {
            task_base = {
                assigned_to  = []
                description  = {
                    enabled = true
                    value   = "<body><ul><li value=\"1\" class=\"listitem\" data-epilot-style=\"\" style=\"\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">hello</span></li><li value=\"2\" class=\"listitem\" style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\" data-epilot-style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">mister</span></li></ul></body>"
                }
                ecp          = {
                    description = "<body><ul><li value=\"1\" class=\"listitem\" data-epilot-style=\"\" style=\"\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">goggins</span></li><li value=\"2\" class=\"listitem\" style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\" data-epilot-style=\"font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\"><span style=\"white-space: pre-wrap;font-family: ProximaNova, 'Open Sans', Helvetica, Arial, sans-serif;\">represents</span></li></ul></body>"
                    enabled     = true
                }
                id           = "9639e6df-2f17-4c9a-99e7-94aa96a3f441"
                name         = "Test STABLE bug"
                requirements = []
                task_type    = "MANUAL"
                taxonomies   = []
            }
        },
    ]
    taxonomies          = [
        "pv-anlage:angebot",
        "pv-anlage:angebot-unterschrieben",
    ]
    trigger             = {
        manual_trigger = {
            entity_schema = "opportunity"
            id            = "82d4e9a5-8a6f-4a39-b36c-fab18ea96baf"
            type          = "manual"
        }
    }
    version             = "v3"
}
