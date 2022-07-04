<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.14.9 |
| <a name="requirement_azuread"></a> [azuread](#requirement\_azuread) | 2.25.0 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | 3.12.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_azuread"></a> [azuread](#provider\_azuread) | 2.25.0 |
| <a name="provider_azurerm"></a> [azurerm](#provider\_azurerm) | 3.12.0 |
| <a name="provider_random"></a> [random](#provider\_random) | 3.3.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azuread_application.auth](https://registry.terraform.io/providers/hashicorp/azuread/2.25.0/docs/resources/application) | resource |
| [azuread_service_principal.auth](https://registry.terraform.io/providers/hashicorp/azuread/2.25.0/docs/resources/service_principal) | resource |
| [azuread_service_principal_password.auth](https://registry.terraform.io/providers/hashicorp/azuread/2.25.0/docs/resources/service_principal_password) | resource |
| [azurerm_cosmosdb_account.cosmosdb](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/cosmosdb_account) | resource |
| [azurerm_cosmosdb_mongo_collection.collection](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/cosmosdb_mongo_collection) | resource |
| [azurerm_cosmosdb_mongo_database.db](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/cosmosdb_mongo_database) | resource |
| [azurerm_key_vault.kv](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/key_vault) | resource |
| [azurerm_key_vault_secret.collection](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/key_vault_secret) | resource |
| [azurerm_key_vault_secret.db](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/key_vault_secret) | resource |
| [azurerm_key_vault_secret.url](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/key_vault_secret) | resource |
| [azurerm_resource_group.rg](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/resource_group) | resource |
| [azurerm_role_assignment.auth](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/resources/role_assignment) | resource |
| [random_integer.ri](https://registry.terraform.io/providers/hashicorp/random/latest/docs/resources/integer) | resource |
| [azurerm_client_config.current](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/data-sources/client_config) | data source |
| [azurerm_subscription.primary](https://registry.terraform.io/providers/hashicorp/azurerm/3.12.0/docs/data-sources/subscription) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_app_name"></a> [app\_name](#input\_app\_name) | name of the application. | `string` | `"gogin"` | no |
| <a name="input_region"></a> [region](#input\_region) | Azure region name. | `string` | `"eastus"` | no |
| <a name="input_subscription_id"></a> [subscription\_id](#input\_subscription\_id) | azure subscription id. | `any` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_client_id"></a> [client\_id](#output\_client\_id) | client id |
| <a name="output_client_secret"></a> [client\_secret](#output\_client\_secret) | client secret |
| <a name="output_cosmosdb_connection_string"></a> [cosmosdb\_connection\_string](#output\_cosmosdb\_connection\_string) | azure cosmos db connection strings |
| <a name="output_keyvault_name"></a> [keyvault\_name](#output\_keyvault\_name) | keyvault name |
| <a name="output_keyvault_uri"></a> [keyvault\_uri](#output\_keyvault\_uri) | keyvault uri |
| <a name="output_tenant_id"></a> [tenant\_id](#output\_tenant\_id) | tenant id |
<!-- END_TF_DOCS -->