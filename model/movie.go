package model

type Movie struct {
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	ReleaseDate  string `json:"release_date"`
	Url          string `json:"url"`
	CommentCount int64  `json:"comment_count"`
}

type MovieResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Results  []Movie `json:"results"`
}
