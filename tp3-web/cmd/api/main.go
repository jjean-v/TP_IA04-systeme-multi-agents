package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Model struct {
	Id                    string `json:"id"`
	Max_completion_tokens int    `json:"max_completions_token"`
	Name                  string `json:"name"`
}

type Data struct {
	Tab []Model `json:"data"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ContentResponse struct {
	Message Message `json:"message"`
}

type ResponseRequest struct {
	Model   string            `json:"model"`
	Choices []ContentResponse `json:"choices"`
}

var GROQ_KEY = os.Getenv("GROQ_API_KEY")

func SendRequestGet(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)
	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body

}

func SendRequestPost(url string, data []byte) []byte {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)

	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body

}

func JsonResponse(body []byte, result any) any {

	err := json.Unmarshal(body, result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func getModel() {

	body := SendRequestGet("https://api.groq.com/openai/v1/models")

	var response Data

	JsonResponse(body, &response)

	for _, element := range response.Tab {
		fmt.Println(element)
	}
}

func chatWitModels(question string) {

	// Prepare the request
	message := []Message{Message{"user", question}}

	request := Request{"openai/gpt-oss-20b", message}

	buffer, errMarshal2 := json.Marshal(request)
	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	body := SendRequestPost("https://api.groq.com/openai/v1/chat/completions", buffer)

	// Unmarshal the response
	var response ResponseRequest

	err := json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range response.Choices {
		fmt.Println(element)
	}
}

func main() {
	chatWitModels("combien font 2+2")
	getModel()

}
