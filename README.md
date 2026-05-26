# epilot-workflow

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/epilot-dev/terraform-provider-epilot-workflow.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/epilot-dev/terraform-provider-epilot-workflow/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<no value>
<!-- Start SDK <no value> -->
To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.8.2"
    }
  }
}

provider "epilot-workflow" {
  # Configuration options go below
}
```

<!-- End SDK <no value> -->

<no value>
<!-- Start SDK <no value> -->
### Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

### Example

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

<!-- End SDK <no value> -->

<no value>
<!-- Start SDK <no value> -->

<!-- End SDK <no value> -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    epilot-workflow = {
      source  = "epilot-dev/epilot-workflow"
      version = "0.18.9"
    }
  }
}

provider "epilot-workflow" {
  bearer_auth = "<YOUR_BEARER_AUTH>" # Required
  server_url  = "..."                # Optional
}
```
<!-- End Installation [installation] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-epilot-workflow`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/epilot-dev/epilot-workflow" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Start Authentication [security] -->
## Authentication

This provider supports authentication configuration via provider configuration.

Available configuration:

| Provider Attribute | Description |
|---|---|
| `bearer_auth` | HTTP Bearer. |
<!-- End Authentication [security] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Managed Resources

* [epilot-workflow_closing_reason](docs/resources/closing_reason.md)
* [epilot-workflow_flow_template](docs/resources/flow_template.md)
* [epilot-workflow_workflow_definition](docs/resources/workflow_definition.md)

### Data Sources

* [epilot-workflow_closing_reason](docs/data-sources/closing_reason.md)
* [epilot-workflow_flow_template](docs/data-sources/flow_template.md)
* [epilot-workflow_workflow_definition](docs/data-sources/workflow_definition.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Summary [summary] -->
## Summary

Workflows Definitions: The Workflows Definitions API enables you to create, manage, and configure reusable workflow templates
within your organization. Workflow definitions serve as blueprints that define the structure and behavior
of business processes, which can then be instantiated as workflow executions.

## Core Concepts

### Workflow Definition (V1)
A workflow definition is a template that describes a business process. It consists of:
- **Sections**: Logical groupings (phases) that organize related steps together
- **Steps**: Individual tasks or actions that need to be completed within a section
- **Closing Reasons**: Predefined reasons that can be selected when closing/completing a workflow

### Flow Template (V2)
The modern workflow model that provides advanced capabilities:
- **Phases**: Named stages that group related tasks and track progress through the workflow
- **Tasks**: Individual units of work that can be manual, automated, AI-powered, or decision points
- **Edges**: Connections between tasks that define the flow sequence and support conditional branching
- **Triggers**: Define how a workflow is started (manual, automation, journey submission)

### Task Types (V2)
- **MANUAL**: Tasks assigned to users that require human action to complete
- **AUTOMATION**: Tasks that execute automated actions when reached
- **DECISION**: Conditional branching points that evaluate conditions to determine the next path
- **AI_AGENT**: Tasks that invoke AI agents to perform intelligent actions

## API Versions
- **V1 endpoints** (`/v1/workflows/...`): Legacy linear workflow model with sections and steps
- **V2 endpoints** (`/v2/flows/...`): Modern graph-based flow model with phases, tasks, and edges
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [epilot-workflow](#epilot-workflow)
  * [Installation](#installation)
  * [Testing the provider locally](#testing-the-provider-locally)
  * [Authentication](#authentication)
  * [Available Resources and Data Sources](#available-resources-and-data-sources)
  * [Core Concepts](#core-concepts)
  * [API Versions](#api-versions)

<!-- End Table of Contents [toc] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/epilot-dev/epilot-workflow" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

Note: To use the dev_overrides, please ensure you run `go build` in this folder. You must have a binary available for terraform to find.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
