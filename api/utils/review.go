package utils

import "time"

// Structs for the objects we will save in the json file.
// review content, author, score, and time the review was submitted.

type AppReview struct {
	Author  string    `json:"author"`
	Rating  string    `json:"rating"`
	Content string    `json:"content"`
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Updated time.Time `json:"updated"`
}

type ReviewFile struct {
	LastUpdated time.Time   `json:"lastUpdated"`
	Reviews     []AppReview `json:"reviews"`
}
