package types

import (
	"encoding/json"
	"log"
	"net/http"
)

// WordOfTheDay is the word of the day JSON
type WordOfTheDay struct {
	Word       string       `json:"word"`
	Definitons []Definition `json:"definitions"`
	Examples   []Example    `json:"examples"`
}

// Definition includes word definitions
type Definition struct {
	Text         string `json:"text"`
	PartOfSpeech string `json:"partOfSpeech"`
}

// Example are sentence examples of the word
type Example struct {
	Text string `json:"text"`
}

// GetWordOfTheDay calls the wordnik api to get the word of the day
func GetWordOfTheDay(url string) (*WordOfTheDay, error) {
	var word WordOfTheDay

	r, err := http.Get(url)
	if err != nil {
		log.Printf("Error in requesting URLs")
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&word); err != nil {
		log.Printf("Error in decoding")
		return nil, err
	}

	return &word, nil
}
