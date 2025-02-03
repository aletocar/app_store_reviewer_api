package main

import (
	"app_store_reviewer/cron"
	"app_store_reviewer/utils"
	"encoding/json"
	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "HEAD", "OPTIONS"},
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"Origin"},
	}))
	router.GET("/apps", getApps)
	router.GET("/apps/:id/reviews", getAppReviews)
	router.GET("/health", health)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

// returns the reviews for an app id, if no reviews exist, returns an empty array
// Due to the fact that not all apps have reviews for the last 48 hours, I added a Query Param to select how many days are queried.
func getAppReviews(c *gin.Context) {
	id := c.Param("id")
	days := c.Query("days")
	if days == "" {
		days = "2" //default to 48 hours
	}
	val, _ := strconv.Atoi(days)
	data := getReviewsForAppId(id, val)
	if len(data) == 0 {
		c.IndentedJSON(http.StatusOK, []string{})
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

// returns a list of apps that have available reviews
func getApps(c *gin.Context) {
	file, _ := os.Open("./applications.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var apps []string
	byteValue, _ := io.ReadAll(file)
	err := json.Unmarshal(byteValue, &apps)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, apps)
}

// default health check endpoint for the api.
func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "OK"})
}

// opens the reviews file for a given id and filters the ones for the last days as provided in the params
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

// filters the reviews for the given days.
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
