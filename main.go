package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	data := getReviewsFromFile()
	if len(data) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": id})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}
func getApps(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No apps have been found"})
}
func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
}

func getReviewsFromFile() []map[string]interface{} {
	file, _ := os.Open("./dbfile.json")
	defer file.Close()
	decoder := json.NewDecoder(file)

	decoder.Token()

	reviews := []map[string]interface{}{}
	data := map[string]interface{}{}

	for decoder.More() {
		decoder.Decode(&data)
		reviews = append(reviews, data)
	}
	return reviews
}
