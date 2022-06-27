package config

import "os"

func getMongoURI() string {
  return os.Getenv("MONGODB_CONNECTION_STRING")
}

func getMongoDB() string {
  return os.Getenv("MONGODB_DATABASE")
}

func getMongoProductCollection() string {
  return os.Getenv("MONGODB_COLLECTION")
}