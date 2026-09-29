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

func SendRequest(method string, url string) []byte {
	req, _ := http.NewRequest(method, url, nil)
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

func getModel() {

	body := SendRequest("GET", "https://api.groq.com/openai/v1/models")

	var response Data

	errMarshal := json.Unmarshal(body, &response)

	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	//fmt.Println(response)
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

	req, errCreateRequest := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(buffer))
	if errCreateRequest != nil {
		fmt.Println("Error creating request:", errCreateRequest)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)

	// Send the request
	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the response
	var response ResponseRequest

	errMarshal3 := json.Unmarshal(body, &response)

	if errMarshal3 != nil {
		log.Fatal(errMarshal3)
	}

	for _, element := range response.Choices {
		fmt.Println(element)
	}
}

func main() {
	//chatWitModels("combien font 2+2")
	getModel()

}
