package cron

import (
	"app_store_reviewer/utils"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"os"
)

func Run() {
	c := cron.New()
	_, err := c.AddFunc("@every 60s", getReviews)
	//run an initial get of the reviews.
	getReviews()
	if err != nil {
		return
	}
	c.Start()
}

// it gets the apps and executes the get reviews for each
func getReviews() {
	apps := getAppList()
	for _, app := range apps {
		utils.GetReviewsForApp(app)
	}
}

// Reads the applications.json file and returns the list of app ids as a string slice.
func getAppList() []string {
	//The app list is saved in the applications.json file
	file, _ := os.Open("./applications.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error when closing applications.json file", err)
		}
	}(file)

	var apps []string
	byteValue, _ := io.ReadAll(file)
	err := json.Unmarshal(byteValue, &apps)
	if err != nil {
		fmt.Printf("there was an error parsing the reviews file: %s\n", err)
		return []string{}
	}
	return apps
}
