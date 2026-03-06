package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/VincentSamuelPaul/AWS/s3storage/config"
	"github.com/VincentSamuelPaul/AWS/s3storage/routes"
)

func main() {

	config.InitS3()

	router := gin.Default()
	routes.Setup(router)

	log.Println("Go server starting on port: 8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
