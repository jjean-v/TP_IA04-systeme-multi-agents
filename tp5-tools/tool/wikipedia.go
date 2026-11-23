package tool

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type WikiMessage struct {
	Filters []Filter `json:"filters"`
	Fields  []string `json:"fields"`
}

type Filter struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type WikiResponse struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	DateModified string `json:"date_modified"`
	Abstract     string `json:"abstract"`
}

func NewWikiMessage() WikiMessage {
	var field = [4]string{"name", "url", "date_modified", "abstract"}
	var filter = Filter{"is_part_of.identifier", "enwiki"}
	var message = WikiMessage{}
	message.Filters = append(message.Filters, filter)
	message.Fields = field[:]

	return message
}

func Chat(article string) (WikiResponse, error) {
	url := os.Getenv("WIKI_URL") + article
	request := NewWikiMessage()
	body, _ := json.Marshal(request)
	log.Println(string(body))
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return WikiResponse{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	if key := os.Getenv("WIKI_TOKEN"); key != "" {
		httpReq.Header.Set("Authorization", "Bearer "+key)
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return WikiResponse{}, err
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		log.Println("ICI")
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return WikiResponse{}, fmt.Errorf("HTTP %d : %s", resp.StatusCode, msg)
	}
	log.Println("ET LA")

	var wr []WikiResponse
	if err := json.NewDecoder(resp.Body).Decode(&wr); err != nil {
		return WikiResponse{}, err
	}
	if len(wr) == 0 {
		return WikiResponse{}, errors.New("réponse vide")
	}
	return wr[0], nil
}

/*
}
curl -s -X POST \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  https://api.enterprise.wikimedia.com/v2/articles/NASA \
  -d '{
    "filters": [{"field": "is_part_of.identifier", "value": "enwiki"}],
    "fields": ["name", "url", "date_modified", "abstract"]
  }'
*/
