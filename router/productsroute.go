package router

import (
	"gogincosmos/model"
	"gogincosmos/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProducRouteHandler struct {}

var pservice service.ProductService

func init() {
  log.Println("ProductsRoute - Init")
  pservice = service.ProductService{}
  
}
func (rh *ProducRouteHandler)ListProducts(c *gin.Context) {
  pr := pservice.ListProducts()
  c.JSON(pr.Status, pr)
}

func (rh *ProducRouteHandler)FindProduct(c *gin.Context) {
  productId := c.Param("productId")
  log.Println("ProductRoute - Find ", productId)
  pr := pservice.FindProduct(productId)
  c.JSON(pr.Status, pr)
}

func (rh *ProducRouteHandler)AddProduct(c *gin.Context) {
  var productRequest model.ProductRequest
  if c.ShouldBind(&productRequest) == nil {
    var productResponse = pservice.AddProduct(productRequest)
    c.JSON(productResponse.Status, productResponse)
    return
  }
  c.JSON(http.StatusInternalServerError, productRequest.BindError())
}

func (rh *ProducRouteHandler)RemoveProduct(c *gin.Context) {
  pr := pservice.ListProducts()
  c.JSON(pr.Status, pr)
}