# Basic example

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 5.14 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_parameter_group"></a> [parameter\_group](#module\_parameter\_group) | ../.. | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_name"></a> [name](#input\_name) | (Required) The name of the ElastiCache parameter group. | `string` | n/a | yes |
| <a name="input_family"></a> [family](#input\_family) | (Required) The family of the ElastiCache parameter group. | `string` | n/a | yes |
| <a name="input_parameters"></a> [parameters](#input\_parameters) | (Required) A list of ElastiCache parameters to apply. | <pre>list(object({<br/>    name  = string<br/>    value = string<br/>  }))</pre> | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_elasticache_parameter_group_name"></a> [elasticache\_parameter\_group\_name](#output\_elasticache\_parameter\_group\_name) | n/a |
| <a name="output_elasticache_parameter_group_family"></a> [elasticache\_parameter\_group\_family](#output\_elasticache\_parameter\_group\_family) | n/a |
| <a name="output_elasticache_parameter_group_parameters"></a> [elasticache\_parameter\_group\_parameters](#output\_elasticache\_parameter\_group\_parameters) | n/a |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
