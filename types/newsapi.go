package types

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Articles is a single article of news
type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// News is the list of articles
type News struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults"`
	Articles     []Articles `json:"articles"`
}

// GetTopHeadlines gets the top headlines from TOI
func GetTopHeadlines(source string) (*Articles, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?sources=%s&apiKey=%s", source, os.Getenv("NEWS_API_TOKEN"))
	var news News

	r, err := http.Get(url)
	if err != nil {
		log.Printf("Error in requesting URLs")
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		log.Printf("Error in decoding")
		return nil, err
	}

	return &news.Articles[0], nil
}
