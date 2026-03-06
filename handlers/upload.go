package handlers

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"

	"github.com/VincentSamuelPaul/AWS/s3storage/config"
)

var bucket = config.BucketName

func getPresignedURL(filename string) (string, error) {
	presignClient := s3.NewPresignClient(config.S3Client)

	req, err := presignClient.PresignGetObject(context.TODO(),
		&s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &filename,
		},
		s3.WithPresignExpires(15*time.Minute), // URL expires in 15 mins
	)
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))

	_, err = config.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filename,
		Body:   src,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Upload failed: %v", err)})
		return
	}

	url, err := getPresignedURL(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File uploaded but failed to generate URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "File uploaded successfully",
		"filename":   filename,
		"url":        url,
		"expires_in": "15 minutes",
	})
}

func ListFiles(c *gin.Context) {
	result, err := config.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to list files: %v", err)})
		return
	}

	files := []gin.H{}
	for _, obj := range result.Contents {
		url, err := getPresignedURL(*obj.Key)
		if err != nil {
			continue
		}
		files = append(files, gin.H{
			"filename":   *obj.Key,
			"size":       *obj.Size,
			"url":        url,
			"expires_in": "15 minutes",
		})
	}

	c.JSON(http.StatusOK, gin.H{"files": files})
}
