package router

import (
	"log"

	"github.com/gin-gonic/gin"
)
var router *gin.Engine

func init() {
  log.Println("Router - Init")
}
func addHomeRoutes(router *gin.Engine) {
  router.GET("/", IndexHandler)
}

func addProductRoutes(router *gin.Engine) {
  productsRouteHandler := ProducRouteHandler{}
  router.GET("/products", productsRouteHandler.ListProducts)
  router.GET("/products/:productId", productsRouteHandler.FindProduct)
  router.POST("/products", productsRouteHandler.AddProduct)
  router.DELETE("/products/:productId", productsRouteHandler.RemoveProduct)
}

func AllRoutes() *gin.Engine {
  router := gin.Default()  
  addHomeRoutes(router)
  addProductRoutes(router)
  return router
}