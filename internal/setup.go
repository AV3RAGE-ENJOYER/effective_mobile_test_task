package internal

import (
	docs "github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/docs"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal/handlers"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/api"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupGinRouter(api api.ExternalApiClient, db repository.DatabaseRepository) *gin.Engine {
	serv := gin.New()

	serv.Use(gin.Recovery(), gin.Logger())

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := serv.Group("/api/v1")
	{
		v1.GET("/library", handlers.GetLibraryInfoHandler(db))
		v1.GET("/info", handlers.GetInfoHandler(db))
		v1.GET("/text", handlers.GetTextHandler(db))
		v1.POST("/add", handlers.AddSongHandler(api, db))
		v1.PUT("/edit", handlers.EditSongHandler(db))
		v1.DELETE("/delete", handlers.DeleteSongHandler(db))
	}
	serv.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return serv
}
