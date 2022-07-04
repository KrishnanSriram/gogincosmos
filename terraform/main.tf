terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.12.0"
    }
    azuread = {
      source  = "hashicorp/azuread"
      version = "2.25.0"
    }
    random = {
      source = "hashicorp/random"
      version = "3.3.2"
    }
  }
  required_version = ">= 0.14.9"
}

provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
}

provider "azuread" {
}

resource "azurerm_resource_group" "rg" {
  name     = format("%s-%s-rg", var.app_name, random_integer.ri.result)
  location = var.region
}

resource "random_integer" "ri" {
  min = 10000
  max = 99999
}

resource "azurerm_cosmosdb_account" "cosmosdb" {
  name                = format("%s-%s-cosmos-db", var.app_name, random_integer.ri.result)
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  offer_type          = "Standard"
  kind                = "MongoDB"

  enable_automatic_failover = true

  capabilities {
    name = "EnableAggregationPipeline"
  }

  capabilities {
    name = "mongoEnableDocLevelTTL"
  }

  capabilities {
    name = "MongoDBv3.4"
  }

  capabilities {
    name = "EnableMongo"
  }

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 300
    max_staleness_prefix    = 100000
  }

  geo_location {
    location          = "eastus"
    failover_priority = 0
  }
}

# Manages a Mongo Database within a Cosmos DB Account.
resource "azurerm_cosmosdb_mongo_database" "db" {
  name                = format("%s-cosmos-mongo-db", var.app_name)
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmosdb.name
}

# Manages a Mongo Collection within a Cosmos DB Account.
resource "azurerm_cosmosdb_mongo_collection" "collection" {
  name                = format("%s-cosmos-mongo-db-collection", var.app_name)
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmosdb.name
  database_name       = azurerm_cosmosdb_mongo_database.db.name

  default_ttl_seconds = "777"
  shard_key           = "_id"
  throughput          = 400

  index {
    keys   = ["_id"]
    unique = true
  }
}

# Manages an application registration within Azure Active Directory.
resource "azuread_application" "auth" {
  display_name = var.app_name
}

# Manages a service principal associated with an application within Azure Active Directory.
resource "azuread_service_principal" "auth" {
  application_id = azuread_application.auth.application_id
}

# Manages a password credential associated with a service principal within Azure Active Directory. 
resource "azuread_service_principal_password" "auth" {
  service_principal_id = azuread_service_principal.auth.id
}

data "azurerm_subscription" "primary" {}

resource "azurerm_role_assignment" "auth" {
  scope                = data.azurerm_subscription.primary.id
  role_definition_name = "Contributor"
  principal_id         = azuread_service_principal.auth.id
}

data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "kv" {
  name                        = format("%s-%s-kv", var.app_name, random_integer.ri.result)
  location                    = azurerm_resource_group.rg.location
  resource_group_name         = azurerm_resource_group.rg.name
  enabled_for_disk_encryption = true
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  soft_delete_retention_days  = 7
  purge_protection_enabled    = false

  sku_name = "standard"

  access_policy {
    object_id = azuread_service_principal.auth.object_id
    tenant_id = data.azurerm_client_config.current.tenant_id

    key_permissions = [
      "Get",
    ]

    secret_permissions = [
      "Get", "List", "Set", "Delete", "Purge"
    ]

    storage_permissions = [
      "Get",
    ]
  }
  access_policy {
    object_id = data.azurerm_client_config.current.object_id
    tenant_id = data.azurerm_client_config.current.tenant_id

    key_permissions = [
      "Get",
    ]

    secret_permissions = [
      "Get", "List", "Set", "Delete", "Purge"
    ]

    storage_permissions = [
      "Get",
    ]
  }
} 

resource "azurerm_key_vault_secret" "url" {
  name         = "MONGOURL"
  value        = tostring("${azurerm_cosmosdb_account.cosmosdb.connection_strings[1]}")
  key_vault_id = azurerm_key_vault.kv.id
}

resource "azurerm_key_vault_secret" "db" {
  name         = "MONGODB"
  value        = azurerm_cosmosdb_mongo_database.db.name
  key_vault_id = azurerm_key_vault.kv.id
}

resource "azurerm_key_vault_secret" "collection" {
  name         = "PRODUCTCOLLECTION"
  value        = azurerm_cosmosdb_mongo_collection.collection.name
  key_vault_id = azurerm_key_vault.kv.id
}