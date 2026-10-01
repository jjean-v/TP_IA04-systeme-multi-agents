package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const CONTEXT_PROMPT = `You are receiving a series of reasoning traces extracted from previous AI responses. 
	Each trace represents the internal chain of thought the model produced before giving its final answer.

	Your task is to analyze these traces and use it as a context.

	The traces will be provided as a list in the next message, in this format:
	- reasoning : "<reasoning text>"`

const URL_MODEL_LIST = "https://api.groq.com/openai/v1/models"
const URL_MODEL_QUESTIONS = "https://api.groq.com/openai/v1/chat/completions"

const MODEL_IA_GPT_OSS_20B = "openai/gpt-oss-20b"
const MODEL_IA_GPT_OSS_SAFEGUARD_20B = "openai/gpt-oss-safeguard-20b"
const MODEL_IA_ALIBABA_QWEN_3_8_27B = "qwen/qwen3.8-27b"
const MODEL_IA_SDAIA_ALLAM_2_7B = "allam-2-7b"

var GROQ_KEY = os.Getenv("GROQ_API_KEY")

var wd, _ = os.Getwd()

var PATH = filepath.Join(wd, "tp3-web", "cmd", "api", "conversation", "chat1.md")

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

type MessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type MessageResponse struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	Reasoning string `json:"reasoning"`
}

type Request struct {
	Model    string           `json:"model"`
	Messages []MessageRequest `json:"messages"`
}

type ContentResponse struct {
	Message MessageResponse `json:"message"`
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

	// Add previous conversation
	context := generateContext()
	message := MessageRequest{"user", question}

	// Add actual question
	context = append(context, message)

	request := Request{ia, context}

	buffer, errMarshal2 := json.Marshal(request)
	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	body := SendRequestPost(URL_MODEL_QUESTIONS, buffer)

	// Unmarshal the response
	var response ResponseRequest

	JsonResponse(body, &response)

	for _, element := range response.Choices {
		fmt.Println(element.Message.Content)
		storeConversation("reasoning: " + element.Message.Reasoning)
	}
}

func storeConversation(message string) {

	file, err := os.OpenFile(PATH, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintln(message + "\n"))
	if err != nil {
		log.Fatal(err)
	}

}

func readLines() ([]string, error) {
	file, err := os.Open(PATH)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func generateContext() []MessageRequest {
	lines, _ := readLines()
	var message []MessageRequest

	message = append(message, MessageRequest{"user", CONTEXT_PROMPT})
	for _, element := range lines {
		if element != "" {
			message = append(message, MessageRequest{"user", element})
		}
	}
	return message
}

func main() {

	for {
		fmt.Print("Question: ")
		reader := bufio.NewReader(os.Stdin)
		question, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		chatWitModels(question, MODEL_IA_GPT_OSS_20B)
	}

	//fmt.Println(generateContext())

}
