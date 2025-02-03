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

// makes the Get request to the itunes api and returns the list of reviews for that page, as well as the lastPage number.
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

// Get the last page of available reviews in the itunes store so that we can check if we finished paging through them.
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

// Pages through all reviews available in the itunes api for an app id and writes the result to a file
func GetReviewsForApp(appId string) {
	lastPage := 1
	var reviews []AppReview
	for i := 1; i <= lastPage; i++ {
		var newReviews []AppReview
		newReviews, lastPage = getReviewsForAppAndPage(appId, lastPage, lastPage)
		reviews = append(reviews, newReviews...)
	}
	//Saving the last updated review as a separate field. This method could be optimized by filtering only the reviews that have been updated after the last updated review.
	lastUpdatedReview := reviews[0].Updated
	WriteFileWithReviews(reviews, appId, lastUpdatedReview)
}

// Write the file for the app reviews.
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

// Converts the entries returned in the feed to the appreview object we will save to the file.
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
