output "cosmosdb_connection_string" {
  value = azurerm_cosmosdb_account.cosmosdb.connection_strings[0]
  description = "azure cosmos db connection strings"
  sensitive = true
}

output "client_secret" {
 value = azuread_service_principal_password.auth.value
 description = "client secret"
 sensitive = true
}

output "client_id" {
 value = azuread_application.auth.application_id
 description = "client id"
}

output "tenant_id" {
 value = azuread_service_principal.auth.application_tenant_id
 description = "tenant id"
}

output "keyvault_name" {
 value = azurerm_key_vault.kv.name
 description = "keyvault name"
}

output "keyvault_uri" {
 value = azurerm_key_vault.kv.vault_uri
 description = "keyvault uri"
}
