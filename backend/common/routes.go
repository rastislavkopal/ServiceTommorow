package common

import (
	"net/http"

	"backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//GinRouter -> Gin Router
type GinRouter struct {
	Gin *gin.Engine
}

//NewGinRouter all the routes are defined here
func NewGinRouter() GinRouter {

	httpRouter := gin.Default()

	httpRouter.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Polls Up and Running..."})
	})

	// add Swagger handler
	docs.SwaggerInfo.Title = "ServiceTommorrow API docs"
	docs.SwaggerInfo.Description = "RESTful API docs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost.yet"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	httpRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return GinRouter{Gin: httpRouter}
}
