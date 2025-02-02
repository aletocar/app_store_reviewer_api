package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

	reviews := ParseEntriesToReviews(result.Feed)
	lastUpdatedReview := reviews[0].Updated
	fmt.Printf("Last updated review: %s\n", lastUpdatedReview)
	WriteFileWithReviews(reviews, appId, lastUpdatedReview)
}

func WriteFileWithReviews(reviews []AppReview, appId string, lastUpdated time.Time) {
	var fileContent ReviewFile
	fileContent.Reviews = reviews
	fileContent.LastUpdated = lastUpdated
	var jsonText, _ = json.Marshal(fileContent)
	var fileName = "./reviews_" + appId + ".json"
	err := os.WriteFile(fileName, jsonText, 0644)
	if err != nil {
		fmt.Printf("Cannot write to file %s, %s\n", fileName, err)
	}

}
func ParseEntriesToReviews(feed Feed) []AppReview {
	var reviews []AppReview
	for _, element := range feed.Entry {
		var review AppReview
		review.Author = element.Author.Label
		review.Rating = element.ImRating.Label
		review.Content = element.Content.Label
		review.Id = element.Id.Label
		review.Title = element.Title.Label
		//2024-09-03T09:29:44-07:00
		date, err := time.Parse(time.RFC3339Nano, element.Updated.Label)
		if err != nil {
			fmt.Printf("Cannot parse date, %s\n", err)
		}
		//Saving all reviews in UTC to better manage sorting afterwards.
		review.Updated = date.UTC()
		reviews = append(reviews, review)
	}
	return reviews
}

/*

	//TODO: 2. Handle pagination of the request.
	//TODO: 3. Write the result to the file
	//TODO: 4. keep the last time it was checked so as to only check until the last one is found
*/
