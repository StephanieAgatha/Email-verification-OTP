package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Books(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Books Endpoint!"})
}
