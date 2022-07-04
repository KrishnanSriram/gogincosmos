package service

import (
	"context"
	"gogincosmos/config"
	"gogincosmos/model"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type ProductService struct{}

func init() {
	log.Println("ProductService - Init")
}

func (service *ProductService) AddProduct(data model.ProductRequest) *model.ProductResponse {
	product := data.ToProduct()

	productCollection := config.AppConfig.DBConfig.GetCollection(config.AppConfig.DBConfig.MongoClient, os.Getenv("MONGO_COLLECTION"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := productCollection.InsertOne(ctx, product)

	if err != nil {
		return &model.ProductResponse{Data: err, Status: http.StatusInternalServerError, Message: "Failed to get product list!"}
	}

	return &model.ProductResponse{Data: result, Status: http.StatusOK, Message: "New product added successfully!"}
}

func (service *ProductService) ListProducts() *model.ProductResponse {
	var products []model.Product

	productCollection := config.AppConfig.DBConfig.GetCollection(config.AppConfig.DBConfig.MongoClient, os.Getenv("MONGO_COLLECTION"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	results, err := productCollection.Find(ctx, bson.M{})
	defer cancel()

	if err != nil {
		return &model.ProductResponse{Data: err, Status: http.StatusInternalServerError, Message: "Failed to get product list!"}
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var product model.Product
		if err = results.Decode(&product); err != nil {
			return &model.ProductResponse{Data: err, Status: http.StatusInternalServerError, Message: "Failed to get product list!"}
		}

		products = append(products, product)
	}

	productResponse := model.ProductResponse{Data: products, Status: http.StatusOK, Message: "Here's a list of all products in the store! Pagination is not available at the moment!"}
	return &productResponse
}

func (service *ProductService) FindProduct(id string) *model.ProductResponse {
	var product model.Product

	productCollection := config.AppConfig.DBConfig.GetCollection(config.AppConfig.DBConfig.MongoClient, os.Getenv("MONGO_COLLECTION"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := productCollection.FindOne(ctx, bson.M{"id": id}).Decode(&product)
	defer cancel()

	if err != nil {
		return &model.ProductResponse{Data: err, Status: http.StatusInternalServerError, Message: "Failed to get product list!"}
	}
	if product == (model.Product{}) {
		return &model.ProductResponse{Data: nil, Status: http.StatusNotFound, Message: "Product you are looking for is not found. Try again later!"}
	}
	//reading from the db in an optimal way
	return &model.ProductResponse{Data: product, Status: http.StatusOK, Message: "Found matching product"}
}

func (service *ProductService) RemoveProduct(id string) *model.ProductResponse {
	productCollection := config.AppConfig.DBConfig.GetCollection(config.AppConfig.DBConfig.MongoClient, os.Getenv("MONGO_COLLECTION"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := productCollection.DeleteOne(ctx, bson.M{"id": id})
	defer cancel()

	if err != nil {
		return &model.ProductResponse{Data: err, Status: http.StatusInternalServerError, Message: "Failed to get product list!"}
	}

	if result.DeletedCount > 1 {
		return &model.ProductResponse{Data: nil, Status: http.StatusOK, Message: "Product found and removed from store!"}
	}
	//reading from the db in an optimal way
	return &model.ProductResponse{Data: nil, Status: http.StatusNotFound, Message: "Product you are looking for is not found. Try again later!"}
}
