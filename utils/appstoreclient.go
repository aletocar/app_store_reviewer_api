package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetReviewsForApp(appId string) {
	println(appId)
	client := &http.Client{}
	url := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=1/json", appId)
	println(url)
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
	/*
		//TODO: Create
		1. Create struct to parse response
		2. Handle pagination of the request.
		3. Write the result to the file
		4. keep the last time it was checked so as to only check until the last one is found
	*/
}
