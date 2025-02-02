package main

import (
	"app_store_reviewer/cron"
	"app_store_reviewer/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {

	cron.Run()
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
	data := getReviewsForAppId(id)
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

func getReviewsForAppId(id string) []utils.AppReview {
	file, _ := os.Open("./reviews_" + id + ".json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var reviews utils.ReviewFile
	byteValue, _ := io.ReadAll(file)
	err := json.Unmarshal(byteValue, &reviews)
	if err != nil {
		return nil
	}
	return reviews.Reviews
}
