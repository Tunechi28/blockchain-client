package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tunechi28/blockchain-client/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tunechi28/blockchain-client/docs" 
)

func main() {
	router := gin.Default()
	h := &handlers.Handler{BC: &handlers.RealClient{}}

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	router.GET("/block/number", h.GetBlockNumberHandler)
	router.GET("/block/:number", h.GetBlockByNumberHandler)
	
	router.Run(":8080")
}
