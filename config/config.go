package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	BucketName = "go-api-uploads-vincent"
	Region     = "eu-north-1"
)

var S3Client *s3.Client

func InitS3() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(Region))

	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	S3Client = s3.NewFromConfig(cfg)
	log.Println("S3 client initialized successfully")
}
