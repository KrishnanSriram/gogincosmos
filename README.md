<!-- ABOUT THE PROJECT -->

## About The Project

To build a microservice, we need a reliable framework that is nimble and stable too. Golang+Gin combination has been a delight. Azure has been one of the leading public cloud solutions. In this codebase, we'll see how to connect your microservice over to cloud DB, in this case CosmosDB
The journey:

- Build a simple GIN base golang framework
- Containarize your solutions
- Add REST controllers
- Connect REST endpoints to local MongoDB
- Create resource group + cosmos DB
- Connect your solution to cosmos DB
- create Key valut and store all DB information in KeyVault
- Connect your solution with KeyVault
- Deploy your docker instance over to Azure WebApp/ContainerApp service
- Connect your service with Application Gateway
- Don't sit back, write more services

### Built With

GoLang has been one of the most simplest of languages to learn for developers. Objective of Golang has been to keep things simple and focussed. It's not a jungle we can get lost.
Together with Gin, Golang is a force to reckon.

- [Golang](https://go.dev)
- [gin](https://github.com/gin-gonic/gin)
- [Mongo](https://www.mongodb.com)
- [CosmosDB](https://azure.microsoft.com/en-us/services/cosmos-db/)
- [KeyVault](https://azure.microsoft.com/en-us/services/key-vault/)
- [WebApp](https://azure.microsoft.com/en-us/services/app-service/web/)
- [ContainerAppService](https://azure.microsoft.com/en-us/services/container-apps/)

<!-- GETTING STARTED -->

### Prerequisites

Install Golang. I used VSCode for development. There are a ton of other IDE's out there. Use the one that you are comfortable with.

- Browse into the directory you put this git code
- go get
- go run main.go

### Installation

1. Setup your Azure account, subscription before you get started
2. Install Azure command CLI. Can be more than handy
3. Create Azure CosmosDB from terminal - check the usage section
4. Create Azure Keyvault and enter vaules for DB connections

<!-- USAGE EXAMPLES -->

## Usage
Follow either of "Step by Step Guidelines" or "Using Terraform Scripts" steps for usage.

### **Step by Step Guidelines**

- Create ResourceGroup

```
az group create --name gogincosmosrg --location "eastus"
```

- Create Cosmos DB

```
az cosmosdb create --name gogincosmos --resource-group gogincosmosrg --kind MongoDB
```

- Create Keyvault

```
az keyvault create --location "east us" --name gogincosmoskv --resource-group gogincosmosrg
```

- Assign contributors to KV. Note: ID used below is the subscription ID

```
az ad sp create-for-rbac --name "gogincosmossp" --role Contributor --scopes /subscriptions/<<SUBSCRIPTION-ID>>
```

- Finally, add permissions. ID used below is the app ID from above command

```
az keyvault set-policy --name gogincosmoskv --spn "e1ae22e5-2d33-4b22-a4e2-193665f0b352" --secret-permissions get list set delete
```

- Setup CosmosDB connection string, DB name and collection name in Azure KeyVault

- Create .env file with below details

```
KEY_VAULT_NAME=<>
AZURE_KEYVAULT_URL=<>
AZURE_CLIENT_ID=<>
AZURE_TENANT_ID=<>
AZURE_CLIENT_SECRET=<>
```
### **Using Terraform Scripts**
- Run terraform apply to create the infrastructure including service provider client, cosmos db and key vault.

```
make SUBSCRIPTION_ID=<subscription-id> tfapply
```

- Load environment file (.env)

```
make SUBSCRIPTION_ID=<subscription-id> loadenv
```

- Destroy the infrastructure.

```
make SUBSCRIPTION_ID=<subscription-id> tfdestroy
```

## Run code

```
go run main.go
```
