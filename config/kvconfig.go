package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type AzKeyVault interface {
	GetSecret(secretName string) (string, error)
}

type AzKeyVaultImpl struct {
	KVLink          string
	AZSecretsClient *azsecrets.Client
}

var AzKeyVaultInstance AzKeyVaultImpl

func createNewKVClient(keyVaultUrl string) (*azsecrets.Client, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	client, err := azsecrets.NewClient(keyVaultUrl, cred, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (kv AzKeyVaultImpl) GetSecret(secretName string) (string, error) {
	getResp, err := kv.AZSecretsClient.GetSecret(context.TODO(), secretName, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
		return "", err
	}

	log.Printf("secretValue: %s\n", *getResp.Value)
	return *getResp.Value, nil
}

func NewAzKVConfig() (*AzKeyVaultImpl, error) {
	keyVaultName := os.Getenv("KEY_VAULT_NAME")
	keyVaultUrl := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)
	secretsClient, err := createNewKVClient(keyVaultUrl)
	if err != nil {
		return nil, err
	}
	return &AzKeyVaultImpl{
		KVLink:          keyVaultUrl,
		AZSecretsClient: secretsClient,
	}, nil
}
