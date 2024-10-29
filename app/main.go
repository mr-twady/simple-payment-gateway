package main

import (
    "net/http"
	"github.com/gin-gonic/gin"
)

func main (){
	router := gin.Default()
	router.GET("/health", healthCheck)
	
	router.Run("localhost:8080")
}

func healthCheck (c *gin.Context){
	c.IndentedJSON(http.StatusOK, "ok")
}
