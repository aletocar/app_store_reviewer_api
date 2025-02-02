package main

import (
	"app_store_reviewer/cron"
	"app_store_reviewer/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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
	days := c.Query("days")
	if days == "" {
		days = "2" //default to 48 hours
	}
	val, _ := strconv.Atoi(days)
	data := getReviewsForAppId(id, val)
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

func getReviewsForAppId(id string, days int) []utils.AppReview {
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

	return filterReviewsForTheLastDays(reviews.Reviews, days)
}

func filterReviewsForTheLastDays(reviews []utils.AppReview, days int) []utils.AppReview {
	var filteredReviews []utils.AppReview
	originalDate := time.Now().AddDate(0, 0, -days)
	for _, review := range reviews {
		if review.Updated.After(originalDate) {
			filteredReviews = append(filteredReviews, review)
		}
	}
	return filteredReviews
}
