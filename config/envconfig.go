package config

import (
	"log"
	"os"
)

func getMongoURI() string {
	return os.Getenv("MONGODB_CONNECTION_STRING")
}

func getMongoDB() string {
	return os.Getenv("MONGODB_DATABASE")
}

func getMongoProductCollection() string {
	return os.Getenv("MONGODB_COLLECTION")
}

func setEnvVariable(kvConfig AzKeyVault, secretName, envVariable string) {
	value, err := kvConfig.GetSecret(secretName)
	if err != nil {
		log.Fatal("Failed to set up Env variables needed for application!")
	}
	os.Setenv(envVariable, value)
}

func setEnvVariablesFromConfig(kvConfig AzKeyVault) {
	setEnvVariable(kvConfig, "MONGOURL", "MONGODB_CONNECTION_STRING")
	setEnvVariable(kvConfig, "MONGODB", "MONGODB_DATABASE")
	setEnvVariable(kvConfig, "PRODUCTCOLLECTION", "MONGODB_COLLECTION")
}

func SetUpEnvForMongoDBConnections() error {
	kvConfig, err := NewAzKVConfig()
	if err != nil {
		return err
	}
	setEnvVariablesFromConfig(kvConfig)
	return nil
}
