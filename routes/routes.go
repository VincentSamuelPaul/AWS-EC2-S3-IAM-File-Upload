package routes

import (
	"github.com/VincentSamuelPaul/AWS/s3storage/handlers"
	"github.com/VincentSamuelPaul/AWS/s3storage/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.Use(middleware.CORS())

	api := router.Group("/api/v1")
	{
		api.GET("/test", handlers.Test)
		api.POST("/upload", handlers.UploadFile)
		api.GET("/files", handlers.ListFiles)
	}
}
