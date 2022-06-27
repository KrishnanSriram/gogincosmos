package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
  db *mongo.Client
  productCollection *mongo.Collection
}

var MongoDBConfig *DBConfig

func (dbConfig *DBConfig)ConnectDB() error {
  connection_string := getMongoURI()
  log.Println("Connect to DB", connection_string)
  client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
  if err != nil {
      log.Fatal(err)
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
      log.Fatal(err)
  }

  //ping the database
  err = client.Ping(ctx, nil)
  if err != nil {
      log.Fatal(err)
  }
  log.Println("Connected to MongoDB")
  dbConfig.db = client
  return nil 
}

func (dbConfig *DBConfig)GetDB() *mongo.Client {
  return dbConfig.db
}

func (dbConfig *DBConfig)SetProductCollection() {
  dbConfig.productCollection = dbConfig.db.Database(getMongoDB()).Collection("apple_products")
}

func (dbConfig *DBConfig)GetProductCollection() *mongo.Collection {
  log.Println("DBConfig - GetProductCollection")
  if(dbConfig.productCollection == nil) {
    dbConfig.SetProductCollection()
  }
  return dbConfig.productCollection
}

func NewDBConfig() error {
  MongoDBConfig = &DBConfig{}
  return MongoDBConfig.ConnectDB()
}

func GetMongoDBConfig() *DBConfig {
  return MongoDBConfig
}
