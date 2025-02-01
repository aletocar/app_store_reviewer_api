package cron

import (
	"github.com/robfig/cron/v3"
)

func Run() {
	c := cron.New()
	_, err := c.AddFunc("@every 10s", getReviews)
	if err != nil {
		return
	}
	c.Start()
}

func getReviews() {
	println("get reviews")
}
