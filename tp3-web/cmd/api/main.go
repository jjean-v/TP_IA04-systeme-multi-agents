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

const URL_MODEL_LIST = "https://api.groq.com/openai/v1/models"
const URL_MODEL_QUESTIONS = "https://api.groq.com/openai/v1/chat/completions"

const MODEL_IA_GPT_OSS_20B = "openai/gpt-oss-20b"
const MODEL_IA_GPT_OSS_SAFEGUARD_20B = "openai/gpt-oss-safeguard-20b"
const MODEL_IA_META_PROMPT_GUARD_2_86M = "meta-llama/llama-prompt-guard-2-86m"
const MODEL_IA_META_PROMPT_GUARD_2_22M = "meta-llama/llama-prompt-guard-2-22m"
const MODEL_IA_SDAIA_ALLAM_2_7B = "allam-2-7b"

var GROQ_KEY = os.Getenv("GROQ_API_KEY")

type Model struct {
	Id                    string   `json:"id"`
	Max_completion_tokens int      `json:"max_completions_token"`
	Name                  string   `json:"name"`
	Owned_by              string   `json:"owned_by"`
	Input_modalities      []string `json:"input_modalities"`
	Output_modalities     []string `json:"output_modalities"`
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

	body := SendRequestGet(URL_MODEL_LIST)

	var response Data

	JsonResponse(body, &response)

	for _, element := range response.Tab {
		fmt.Println("======================= Model =======================")
		fmt.Println("Id: ", element.Id)
		fmt.Println("Name: ", element.Name)
		fmt.Println("Owned by: ", element.Owned_by)
		fmt.Println("Input: ", element.Input_modalities)
		fmt.Println("Input: ", element.Output_modalities)

	}
}

func chatWitModels(question string, ia string) {

	// Prepare the request
	message := []Message{Message{"user", question}}

	request := Request{ia, message}

	buffer, errMarshal2 := json.Marshal(request)
	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	body := SendRequestPost(URL_MODEL_QUESTIONS, buffer)

	// Unmarshal the response
	var response ResponseRequest

	err := json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
	}

	for _, element := range response.Choices {
		fmt.Println(element.Message.Content)
	}
}

func main() {
	chatWitModels("What s the date today ?", "canopylabs/orpheus-v1-english")
	//getModel()

}
