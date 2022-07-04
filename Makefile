APP_NAME ?= gogin

check-subscription_id:
ifndef SUBSCRIPTION_ID
	$(error SUBSCRIPTION_ID is not set)
endif

tfinstalldeps:
	@brew bundle --file=Brewfile

tflint:
	@tflint ./terraform -c ./terraform/.tflint.hcl

tfdoc:
	@terraform-docs markdown table --output-file README.md terraform

tfinit: 
	@terraform -chdir=terraform init --upgrade

tfplan: check-subscription_id tfinit
	@terraform -chdir=terraform plan -var="app_name=${APP_NAME}" -var="subscription_id=${SUBSCRIPTION_ID}"

tfapply: check-subscription_id tfinit
	@terraform -chdir=terraform apply -auto-approve -var="app_name=${APP_NAME}" -var="subscription_id=${SUBSCRIPTION_ID}"

tfdestroy: check-subscription_id
	@terraform -chdir=terraform destroy -auto-approve -var="subscription_id=${SUBSCRIPTION_ID}"
	@rm -rf .env

loadenv: tfapply
	@sleep 5
	@echo "KEY_VAULT_NAME=$(shell terraform -chdir=terraform output -raw keyvault_name)" > .env
	@echo "AZURE_KEYVAULT_URL=$(shell terraform -chdir=terraform output -raw keyvault_uri)" >> .env
	@echo "AZURE_CLIENT_ID=$(shell terraform -chdir=terraform output -raw client_id)" >> .env
	@echo "AZURE_TENANT_ID=$(shell terraform -chdir=terraform output -raw tenant_id)" >> .env
	@echo "AZURE_CLIENT_SECRET=$(shell terraform -chdir=terraform output -raw client_secret)" >> .env