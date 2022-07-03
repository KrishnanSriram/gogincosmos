package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	ConnectionString string
	Database         string
	MongoClient      *mongo.Client
}

func (dc *DBConfig) ConnectDB() (*mongo.Client, error) {
	log.Println("Connect to DB", dc.ConnectionString)
	client, err := mongo.NewClient(options.Client().ApplyURI(dc.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Connected to MongoDB")

	return client, nil
}

func (dc *DBConfig) GetCollectionInDB(db *mongo.Client, db_name string, collection_name string) *mongo.Collection {
	log.Println("DBConfig - GetProductCollection")
	return db.Database(db_name).Collection(collection_name)
}

func (dc *DBConfig) GetCollection(db *mongo.Client, collection_name string) *mongo.Collection {
	return dc.GetCollectionInDB(db, dc.Database, collection_name)
}

func NewMongoDBConfig(envConfig EnvConfig) *DBConfig {
	return &DBConfig{
		ConnectionString: envConfig.GetParam("MONGODB_CONNECTIONSTRING"),
		Database:         envConfig.GetParam("MONGO_DATABASE"),
	}
}
