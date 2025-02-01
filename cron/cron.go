package cron

import (
	"app_store_reviewer/utils"
	"github.com/robfig/cron/v3"
)

func Run() {
	println("Setting up Cron")
	c := cron.New()
	_, err := c.AddFunc("@every 10s", getReviews)
	if err != nil {
		return
	}
	c.Start()
}

func getReviews() {
	println("get reviews")
	utils.GetReviewsForApp("595068606")
}
