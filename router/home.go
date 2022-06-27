package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "Hello, check on other routes for DB connections!",
  })
}