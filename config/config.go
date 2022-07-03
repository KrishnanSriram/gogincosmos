package config

import "log"

type Config struct {
	EnvConfig
	DBConfig
}

// Has to be set by main
var AppConfig Config

func SetupApplicationConfig() (*Config, error) {
	envConfig := NewEnvConfig()
	kvConfig := NewKVConfig(*envConfig)
	mongo_connection_string := kvConfig.GetValueFor("MONGOURL")
	// GET all necessary value for DB connection and set it as ENV variables
	mongo_db := kvConfig.GetValueFor("MONGODB")
	mongo_collection := kvConfig.GetValueFor("PRODUCTCOLLECTION")
	envConfig.SetParam("MONGODB_CONNECTIONSTRING", mongo_connection_string)
	envConfig.SetParam("MONGO_DATABASE", mongo_db)
	envConfig.SetParam("MONGO_COLLECTION", mongo_collection)
	// connect to MongoDB and retain connection to ensure we re-use connection
	mongoDBConfig := NewMongoDBConfig(*envConfig)
	client, err := mongoDBConfig.ConnectDB()
	if err != nil {
		log.Println("Failed to connect to MonogDB")
		return nil, err
	}
	mongoDBConfig.MongoClient = client

	return newConfig(*envConfig, *mongoDBConfig), nil
}

func newConfig(envConfig EnvConfig, dbConfig DBConfig) *Config {
	return &Config{
		EnvConfig: envConfig,
		DBConfig:  dbConfig,
	}
}
