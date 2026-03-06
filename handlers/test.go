package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Test(c *gin.Context) {
	response := Response{
		Title:       "Test1",
		Description: "Hi, my name is Vincent Samuel Paul",
	}

	c.JSON(http.StatusOK, response)
}
