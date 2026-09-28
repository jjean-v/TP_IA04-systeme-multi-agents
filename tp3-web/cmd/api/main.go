package main

import (
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

type Request struct {
	Model string `json:"model"`
	Messages string `json:"messages"`
}

type Message struc {
	Role string `json:"role"`
}

var GROQ_KEY = os.Getenv("GROQ_API_KEY")

func getModel() {

	req, _ := http.NewRequest("GET", "https://api.groq.com/openai/v1/models", nil)
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)
	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

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

func chatWitModels() {
	req, _ := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", nil)
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)
	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

}

func main() {
	getModel()

}
