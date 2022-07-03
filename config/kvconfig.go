package config

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type KVConfig struct {
	KVLink string
}

func createNewSecretClient(clientUrl string) *azsecrets.Client {
	//Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	//Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(clientUrl, cred, nil)
	if err != nil {
		log.Fatalf("failed to connect to client: %v", err)
	}

	return client
}

func (kc *KVConfig) GetValueFor(param string) string {
	client := createNewSecretClient(kc.KVLink)
	resp, err := client.GetSecret(context.TODO(), param, nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
	}

	log.Printf("secretValue: %s\n", *resp.Value)
	return *resp.Value
}

func NewKVConfig(envConfig EnvConfig) *KVConfig {
	return &KVConfig{
		KVLink: envConfig.GetParam("AZURE_KEYVAULT_URL"),
	}
}
