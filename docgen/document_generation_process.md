# Document Generation Process Guide

This guide describes the current custom resource documentation generator in this
repository. The generator is YAML-driven: it reads resource definitions from
`generator/defs`, reads the matching Terraform example from `examples/resources`,
renders `templates/resources.md.tmpl`, and writes Markdown files under
`docs/resources`.

The current custom generator is separate from `tfplugindocs`. It gives the
project direct control over the resource documentation format and over which YAML
fields are shown in the final Markdown.

## Quick Command

Run the custom generator from the repository root:

```shell
go run ./docgen
```

The command supports these flags:

| Flag | Default | Purpose |
| --- | --- | --- |
| `-defs` | `generator/defs/defs.yaml` | Manifest that lists resource YAML files to process. |
| `-templates` | `templates` | Directory containing `resources.md.tmpl`. |
| `-out` | `docs/resources` | Output directory for generated resource documentation. |
| `-provider-prefix` | `nd` | Terraform provider prefix used to build names like `nd_local_user`. |
| `-examples-dir` | `examples/resources` | Directory containing example resource folders. |

Example with explicit paths:

```shell
go run ./docgen \
  -defs generator/defs/defs.yaml \
  -templates templates \
  -out docs/resources \
  -provider-prefix nd \
  -examples-dir examples/resources
```

## Files Involved

| File or Directory | Role |
| --- | --- |
| `docgen/doc_gen.go` | Custom Go renderer. Reads YAML, prepares template data, renders Markdown. |
| `generator/defs/defs.yaml` | Manifest of resource definition YAML files to process. |
| `generator/defs/<resource>.yaml` | Resource source data used for docs. |
| `templates/resources.md.tmpl` | Generic Go template used when a resource-specific template is not available. |
| `templates/resources/<resource>.md.tmpl` | Optional resource-specific template that overrides the generic template for one resource. |
| `examples/resources/<terraform_resource_name>/resource.tf` | Required HCL example rendered in the Example Usage section. |
| `docs/resources/<resource>.md` | Generated resource documentation output. |
| `generate.sh` | Existing full generator script. Currently still uses `tfplugindocs`, not `docgen`. |

Current POC input path for local user:

```text
generator/defs/defs.yaml
  -> generator/defs/local_user.yaml
  -> examples/resources/nd_local_user/resource.tf
  -> templates/resources/local_user.md.tmpl when present
  -> otherwise templates/resources.md.tmpl
  -> docs/resources/local_user.md
```

## End-to-End Workflow

```text
1. Start command
   go run ./docgen

2. Read manifest
   generator/defs/defs.yaml

3. For each active YAML file in the manifest
   generator/defs/<resource>.yaml

4. Parse resource YAML
   resource name, title, doc category, API info, GUI info, notes, attributes

5. Skip non-resource entries
   skip when resource.name is empty
   skip when generate_tf_resource is false

6. Prepare derived resource values
   ResourceName = provider-prefix + "_" + resource.name
   Title = resource_title, or a fallback from resource.name
   DocCategory = doc_category, or category

7. Read the required example file
   examples/resources/<ResourceName>/resource.tf

8. Group top-level attributes
   Required
   Optional
   Read-Only

9. Select and render template
   templates/resources/<resource.name>.md.tmpl when present
   otherwise templates/resources.md.tmpl

10. Write output
   docs/resources/<resource.name>.md
```

For `local_user`, the default command produces:

```text
docs/resources/local_user.md
```

using the Terraform resource name:

```text
nd_local_user
```

## Function Responsibility Map

The generator code is organized around the documentation flow.

| Function | Responsibility |
| --- | --- |
| `main` | Reads command-line flags and starts generation. |
| `run` | Coordinates the full flow: manifest -> YAML definition -> template data -> Markdown file. |
| `readManifest` | Parses `generator/defs/defs.yaml`. |
| `readDefinition` | Parses one resource YAML file, derives resource-level values, groups attributes, and reads the example HCL. |
| `groupAttributes` | Splits top-level attributes into Required, Optional, and Read-Only sections. |
| `selectTemplatePath` | Selects `templates/resources/<resource>.md.tmpl` when present, otherwise falls back to `templates/resources.md.tmpl`. |
| `renderTemplate` | Parses and executes the selected Go template. |
| `renderAttributeMarkdown` | Renders one schema attribute and its nested children. |
| `writeAttributeMarkdown` | Writes nested attribute Markdown with Registry-friendly indentation. |
| `skipAttribute` | Hides attributes with no `tf_name` or with `tf_hide: true`. |
| `attributeTypeLabel` | Adds the `Sensitive, ` prefix when `sensitive: true`. |
| `readExampleHCL` | Reads the required `resource.tf` example file. |
| `titleFromName` | Creates a readable title only when `resource_title` is missing. |
| `examplesURL` | Builds the public GitHub examples folder URL. |
| `markdownValues` | Formats YAML `valid_values` as inline Markdown code values. |
| `allowedRange` | Builds numeric range text from `min_int`, `max_int`, `min_float`, and `max_float`. |
| `rangeText` | Formats complete or one-sided range text. |
| `markdownValue` | Formats scalar default values for Markdown. |
| `valueIsEmpty` | Decides whether a default value should be rendered. |

## Manifest Behavior

The manifest is the only list of YAML files that `docgen` processes:

```yaml
---
files:
  - local_user.yaml
```

Only uncommented entries are active. In the current checkout, `local_user.yaml`
is the active POC resource. Commented entries are ignored by YAML parsing.

Each manifest entry is resolved relative to the manifest directory. With the
default manifest, `local_user.yaml` resolves to:

```text
generator/defs/local_user.yaml
```

## Resource YAML Contract

The generator reads the `resource` object from each YAML file:

```yaml
resource:
  name: local_user
  doc_category: Users and Security
  category: Infra
  generate_tf_resource: true
  description: Manages local user for Nexus Dashboard
  api_endpoint: /api/v1/infra/aaa/localUsers
  api_reference: https://developer.cisco.com/docs/nexus-dashboard/latest/list-local-users/
  ui_location: "Admin -> Users and Security -> Users -> Local"
  resource_title: Local User
  attributes:
    - model_name: loginID
      tf_name: login_id
      description: The User ID of the local user.
      type: String
      mandatory: true
```

Supported resource-level fields:

| YAML Field | Required | Used For |
| --- | --- | --- |
| `name` | Yes | Output filename and Terraform resource name suffix. |
| `generate_tf_resource` | Yes | Resource is rendered only when this is `true`. |
| `resource_title` | No | Human-readable title in API, examples, and import text. |
| `doc_category` | No | Terraform Registry frontmatter `subcategory`. |
| `category` | No | Fallback for `doc_category`. |
| `description` | Recommended | Frontmatter description, H1 body text, and summary text. |
| `resource_note` | No | Optional note rendered below the description. |
| `import_note` | No | Optional note rendered at the start of the Importing section. |
| `api_endpoint` | No | API Information endpoint line. |
| `api_reference` | No | API Information link. |
| `ui_location` | No | GUI Information location. |
| `attributes` | Recommended | Schema sections. |

### Derived Resource Values

`ResourceName` is derived by joining the provider prefix and YAML resource name:

```text
ResourceName = <provider-prefix> + "_" + resource.name
```

With defaults:

```text
provider-prefix = nd
resource.name = local_user
ResourceName = nd_local_user
```

The generated file path uses the YAML name, not the Terraform resource name:

```text
docs/resources/local_user.md
```

The template heading, page title, sidebar key, examples path, and import command
use the Terraform resource name:

```text
nd_local_user
```

`resource_title` is preferred for readable text. If it is missing, `docgen`
falls back to a title generated from `resource.name`, such as `local_user` ->
`Local User`.

`doc_category` is preferred for the frontmatter subcategory. If it is missing,
`category` is used as the fallback.

## Example File Behavior

The Example Usage section is sourced from a real example file:

```text
examples/resources/<ResourceName>/resource.tf
```

For `local_user`, the required file is:

```text
examples/resources/nd_local_user/resource.tf
```

The generator fails if this file is missing or cannot be read. There is no HCL
fallback generated from YAML.

The rendered documentation also includes:

- The local example source path.
- A public examples folder URL built from the Terraform resource name.

For `local_user`, the public examples URL is:

```text
https://github.com/CiscoDevNet/terraform-provider-nd/tree/main/examples/resources/nd_local_user
```

## Attribute YAML Contract

The generator intentionally renders only the documentation fields that are
needed in the schema output.

Supported attribute-level fields:

| YAML Field | Used For |
| --- | --- |
| `tf_name` | Terraform attribute name. Required for rendering the attribute. |
| `model_name` | API/model field name shown in parentheses after the Terraform name. |
| `description` | Attribute description text. |
| `type` | Attribute type shown in the schema line. |
| `mandatory` | Places top-level attribute in the Required section. |
| `optional` | Places top-level attribute in the Optional section. |
| `computed` | Places top-level attribute in the Read-Only section when it is not required or optional. |
| `payload_hide` | Places top-level attribute in the Read-Only section when it is not required or optional. |
| `tf_hide` | Excludes the attribute from docs when `true`. |
| `sensitive` | Adds `Sensitive, ` before the type. |
| `default_value` | Renders a Default Value line. |
| `valid_values` | Renders a Valid Values line. |
| `min_int` | Minimum integer range value. |
| `max_int` | Maximum integer range value. |
| `min_float` | Minimum float range value. |
| `max_float` | Maximum float range value. |
| `attributes` | Nested child attributes rendered under the parent. |

Fields such as `validator`, `example`, `use_state`, `write_only`,
`tf_requires_replace`, `ndfc_nested`, `ndfc_type`, and other generator/runtime
fields may still exist in YAML, but the current doc generator does not render
them directly.

## Schema Section Rules

Only top-level attributes are grouped into schema sections.

Required:

```yaml
mandatory: true
```

Optional:

```yaml
optional: true
```

Read-Only:

```yaml
computed: true
```

or:

```yaml
payload_hide: true
```

The grouping priority is:

```text
mandatory -> optional -> computed/payload_hide
```

That means an attribute with both `optional: true` and `computed: true` appears
in Optional, not Read-Only.

An attribute is skipped when:

```yaml
tf_hide: true
```

or when `tf_name` is empty.

Nested attributes are not placed in their own top-level schema section. They are
rendered recursively under their parent attribute.

## Attribute Rendering Format

Each attribute is rendered as:

```md
* `tf_name` (model_name) - (Type) description
```

Example:

```md
* `login_id` (loginID) - (String) The User ID of the local user.
```

Sensitive attributes render as:

```md
* `user_password` (password) - (Sensitive, String) The password of the local user.
```

Default values render below the attribute:

```md
* `remote_user_authorization` (xLaunch) - (Bool) The Remote user authorization is used for signing into Nexus Dashboard.
    * Default Value: `false`
```

Valid values come directly from the YAML `valid_values` list:

```yaml
valid_values:
  - approver
  - designer
  - fabric-admin
```

and render as:

```md
    * Valid Values: `approver`, or `designer`, or `fabric-admin`.
```

The generator does not parse valid values from `validator`.

Allowed ranges come from the existing generator number fields:

```yaml
min_int: 0
max_int: 9
```

or:

```yaml
min_float: 0.1
max_float: 99.9
```

Range rendering behavior:

| YAML Input | Rendered Text |
| --- | --- |
| `min_int: 0`, `max_int: 9` | `Allowed Range: 0 to 9` |
| `min_int: 1` | `Allowed Range: 1 or higher` |
| `max_int: 10` | `Allowed Range: up to 10` |
| `min_float: 0.1`, `max_float: 99.9` | `Allowed Range: 0.1 to 99.9` |

The actual generated Markdown wraps the range value in inline code formatting.

## Nested Attribute Rendering

Nested attributes are rendered under the parent attribute using four spaces per
level. This indentation is important for Terraform Registry doc preview.

YAML:

```yaml
- model_name: domains
  tf_name: security_domains
  description: The security domains of the local user.
  type: Map
  mandatory: true
  attributes:
    - model_name: roles
      tf_name: roles
      description: The list of Nexus Dashboard Roles of the local user.
      type: List:String
      mandatory: true
      valid_values:
        - approver
        - designer
```

Rendered Markdown:

```md
* `security_domains` (domains) - (Map) The security domains of the local user.
    * `roles` (roles) - (List:String) The list of Nexus Dashboard Roles of the local user.
        * Valid Values: `approver`, or `designer`.
```

## Template Behavior

The generic fallback template is:

```text
templates/resources.md.tmpl
```

### Resource-Specific Template Support

Resource-specific templates are supported.

For each resource, `docgen` checks for this template first:

```text
templates/resources/<resource.name>.md.tmpl
```

For `local_user`, the resource-specific template path is:

```text
templates/resources/local_user.md.tmpl
```

If that file exists, it is used for `local_user` only. If it does not exist,
`docgen` falls back to:

```text
templates/resources.md.tmpl
```

The selection logic is:

```text
if templates/resources/<resource.name>.md.tmpl exists:
  use that resource-specific template
else:
  use templates/resources.md.tmpl
```

This allows one-off customization for a resource while keeping the generic
template as the default for every other resource.

Both generic and resource-specific templates receive the same `TemplateData`
contract and can call the same template functions, including `renderAttribute`.

It receives this data from `docgen`:

```go
type TemplateData struct {
    Resource           ResourceDef
    ExampleHCL         string
    ExampleSource      string
    ExamplesURL        string
    RequiredAttributes []AttributeDef
    OptionalAttributes []AttributeDef
    ReadOnlyAttributes []AttributeDef
}
```

The template currently renders these sections:

1. Frontmatter
2. Resource heading
3. Description
4. Optional resource note
5. API Information
6. GUI Information
7. Example Usage
8. Schema
9. Importing

### Frontmatter

Frontmatter values come from YAML-derived resource data:

```md
---
subcategory: "{{ .Resource.DocCategory }}"
layout: "nd"
page_title: "ND: {{ .Resource.ResourceName }}"
sidebar_current: "docs-nd-resource-{{ .Resource.ResourceName }}"
description: |-
  {{ .Resource.Description }}
---
```

### Resource Note

`resource_note` is rendered only when content is present:

```yaml
resource_note: This resource has a special note.
```

### API Information

When `api_reference` exists, the template renders a link:

```md
* Local User Management [API Information](...)
```

When `api_reference` is missing, the template renders a fallback message:

```md
* Local User Management API information is not defined in the YAML.
```

When `api_endpoint` exists, the template renders:

```md
* API Endpoint: `/api/v1/infra/aaa/localUsers`
```

When `api_endpoint` is missing, the template renders:

```md
* API Endpoint: not defined in the YAML.
```

### GUI Information

When `ui_location` exists, the template renders:

```md
* Location: `Admin -> Users and Security -> Users -> Local`
```

When `ui_location` is missing, the template renders:

```md
* Location: not defined in the YAML.
```

### Example Usage

The template prints:

- Generic explanatory text using `.Resource.Title`.
- The local example source path.
- The example HCL from `resource.tf`.
- The public examples folder URL.

### Schema

The template renders:

- `### Required ###`
- `### Optional ###`
- `### Read-Only ###`

Each section uses the `renderAttribute` template function. If a section has no
attributes, it renders:

```md
No required attributes.
```

or the matching optional/read-only message.

### Importing

`import_note` is rendered only when content is present.

The current import command is generic and uses `{id}`:

```shell
terraform import nd_local_user.example {id}
```

The Terraform 1.5 import block is also generic:

```hcl
import {
  id = "{id}"
  to = nd_local_user.example
}
```

If a resource needs a different import identifier format, add that information
to `import_note` or extend the YAML/template contract in a future change.

## Current Intentional Choices

The current implementation keeps the POC strict and simple:

- The generic template is used for all resources unless a matching
  resource-specific template exists.
- `templates/resources/local_user.md.tmpl` is optional. When present, it only
  affects `local_user`; when absent, `local_user` uses the generic template.
- Example HCL must come from `examples/resources/<ResourceName>/resource.tf`.
- Missing example files fail generation.
- There is no generated HCL fallback from YAML.
- `valid_values` must be provided directly in YAML.
- Validator strings are not parsed for documentation.
- Numeric ranges use `min_int`, `max_int`, `min_float`, and `max_float`.
- Attribute-level output is intentionally limited to name, model name, type,
  sensitive marker, description, default value, valid values, allowed range, and
  nested attributes.
- Raw YAML debug sections are not rendered.
- Parsed resource debug sections are not rendered.
- Multiline description formatting workarounds are not implemented.

## Adding Documentation for Another Resource

1. Add or update the resource YAML under `generator/defs`.

2. Ensure the resource has at least:

   ```yaml
   resource:
     name: example_resource
     generate_tf_resource: true
     description: Manages example resource for Nexus Dashboard
     resource_title: Example Resource
     doc_category: Example Category
     attributes:
       - model_name: name
         tf_name: name
         description: The name of the example resource.
         type: String
         mandatory: true
   ```

3. Add the file to `generator/defs/defs.yaml`:

   ```yaml
   files:
     - example_resource.yaml
   ```

4. Add a real example file:

   ```text
   examples/resources/nd_example_resource/resource.tf
   ```

5. Run:

   ```shell
   go run ./docgen
   ```

6. Inspect:

   ```text
   docs/resources/example_resource.md
   ```

7. Preview the generated Markdown with the Terraform Registry doc preview tool
   when Registry formatting needs to be checked.

## Troubleshooting

### Resource Was Not Generated

Check:

- The YAML file is listed and uncommented in `generator/defs/defs.yaml`.
- `resource.name` is set.
- `generate_tf_resource: true` is set.

### Missing Example Error

The example file must exist at:

```text
examples/resources/<ResourceName>/resource.tf
```

For `local_user`:

```text
examples/resources/nd_local_user/resource.tf
```

### Valid Values Are Missing

Add a `valid_values` list to the attribute:

```yaml
valid_values:
  - value-a
  - value-b
```

The generator does not parse values from `validator`.

### Allowed Range Is Missing

Add one or more existing numeric range fields:

```yaml
min_int: 0
max_int: 9
```

or:

```yaml
min_float: 0.1
max_float: 99.9
```

### Attribute Is Missing

Check whether:

- `tf_name` is empty.
- `tf_hide: true` is set.
- The attribute is nested under another attribute and therefore renders under
  that parent instead of in a top-level section.

### Attribute Appears in Optional Instead of Read-Only

If an attribute has both `optional: true` and `computed: true`, it appears in the
Optional section. The grouping priority is:

```text
mandatory -> optional -> computed/payload_hide
```

### API or GUI Fallback Text Appears

Add the missing YAML field:

```yaml
api_reference: https://developer.cisco.com/docs/nexus-dashboard/latest/list-local-users/
api_endpoint: /api/v1/infra/aaa/localUsers
ui_location: "Admin -> Users and Security -> Users -> Local"
```

## Maintenance Notes

When generated documentation is wrong, update the source of truth first:

- `generator/defs/<resource>.yaml`
- `templates/resources.md.tmpl`
- `docgen/doc_gen.go`
- `examples/resources/<ResourceName>/resource.tf`

Avoid hand-editing `docs/resources/<resource>.md` for permanent fixes because
the generated file will be overwritten the next time `docgen` runs.
