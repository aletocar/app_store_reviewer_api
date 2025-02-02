package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetReviewsForApp(appId string) {
	client := &http.Client{}
	url := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=1/json", appId)
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Cannot unmarshall json, %s\n", err)
	}

	/*

		//TODO: 2. Handle pagination of the request.
		//TODO: 3. Write the result to the file
		//TODO: 4. keep the last time it was checked so as to only check until the last one is found
	*/
}
