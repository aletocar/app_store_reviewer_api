package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/apps", getApps)
	router.GET("/apps/:id/reviews", getAppReviews)
	router.GET("/health", health)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func getAppReviews(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": id})
}
func getApps(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No apps have been found"})
}
func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
}
