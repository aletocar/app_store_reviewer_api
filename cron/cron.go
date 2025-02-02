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
	println("Setting up Cron")
	c := cron.New()
	_, err := c.AddFunc("@every 60s", getReviews)
	getReviews()
	if err != nil {
		return
	}
	c.Start()
}

func getReviews() {
	println("get reviews")
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
		fmt.Printf("there was an error parsing the reviews file: %s\n", err)
	}
	for _, app := range apps {
		utils.GetReviewsForApp(app)
	}

}
