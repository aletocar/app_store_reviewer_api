package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func getReviewsForAppAndPage(appId string, page int, lastPage int) ([]AppReview, int) {
	url := fmt.Sprintf("https://itunes.apple.com/us/rss/customerreviews/id=%s/sortBy=mostRecent/page=%s/json", appId, strconv.Itoa(page))
	client := &http.Client{}
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
	lastPage = getLastPage(result.Feed.Link)
	return ParseEntriesToReviews(result.Feed), lastPage
}

func getLastPage(links []Link) int {
	for _, link := range links {
		if link.Attributes.Rel == "last" {
			startIndex := len("https://itunes.apple.com/us/rss/customerreviews/page=")
			lastString := string(link.Attributes.Href[startIndex:len(link.Attributes.Href)])
			lastPage := strings.SplitAfter(lastString, "/")
			lastPageInt, _ := strconv.Atoi(lastPage[0][0 : len(lastPage[0])-1])
			return lastPageInt
		}
	}
	return 0
}

func GetReviewsForApp(appId string) {
	lastPage := 1
	var reviews []AppReview
	for i := 1; i <= lastPage; i++ {
		var newReviews []AppReview
		newReviews, lastPage = getReviewsForAppAndPage(appId, lastPage, lastPage)
		reviews = append(reviews, newReviews...)
	}
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
		review.Author = element.Author.Name.Label
		review.Rating = element.ImRating.Label
		review.Content = element.Content.Label
		review.Id = element.Id.Label
		review.Title = element.Title.Label
		date, err := time.Parse(time.RFC3339Nano, element.Updated.Label)
		if err != nil {
			fmt.Printf("Cannot parse date, %s\n", err)
		}
		//Saving all reviews in UTC to handle everything uniformly
		review.Updated = date.UTC()
		reviews = append(reviews, review)
	}
	return reviews
}

/*
	//TODO: 2. Handle pagination of the request.
*/
