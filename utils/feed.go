package utils

type Response struct {
	Feed Feed `json:"feed"`
}

type Feed struct {
	Author  Author    `json:"author"`
	Entry   []Entry   `json:"entry"`
	Updated FeedLabel `json:"updated"`
	Rights  FeedLabel `json:"rights"`
	Title   FeedLabel `json:"title"`
	Icon    FeedLabel `json:"icon"`
	Link    []Link    `json:"link"`
	Id      FeedLabel `json:"id"`
}

type Author struct {
	Name  FeedLabel `json:"name"`
	Uri   FeedLabel `json:"uri"`
	Label string    `json:"label"`
}

type Entry struct {
	Author        Author    `json:"author"`
	Updated       FeedLabel `json:"updated"`
	ImRating      FeedLabel `json:"im:rating"`
	ImVersion     FeedLabel `json:"im:version"`
	Id            FeedLabel `json:"id"`
	Title         FeedLabel `json:"title"`
	Link          Link      `json:"link"`
	ImVoteSum     FeedLabel `json:"im:voteSum"`
	ImContentType Attribute `json:"im:contentType"`
	ImVoteCount   FeedLabel `json:"im:voteCount"`
}

type FeedLabel struct {
	Label string `json:"label"`
}

type Content struct {
	Label      string    `json:"label"`
	Attributes Attribute `json:"attribute"`
}
type Link struct {
	Attributes Attribute `json:"attributes"`
}

type Attribute struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
	Type string `json:"type"`
}

type ContentTypeAttribute struct {
	Term  string `json:"term"`
	Label string `json:"label"`
}
