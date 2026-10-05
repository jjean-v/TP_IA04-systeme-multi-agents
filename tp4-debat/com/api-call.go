package com

import (
	"fmt"
)

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

type MessageResponse struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	Reasoning string `json:"reasoning"`
}

type ContentResponse struct {
	Message MessageResponse `json:"message"`
}

type ResponseRequest struct {
	Model   string            `json:"model"`
	Choices []ContentResponse `json:"choices"`
}

func GetModel() {

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

func ChatWitModels(buffer []byte) *ResponseRequest {

	body := SendRequestPost(URL_MODEL_QUESTIONS, buffer)

	// Unmarshal the response
	var response ResponseRequest

	JsonResponse(body, &response)

	return &response
}

func Reasoning(question string, ia string) *ResponseRequest {

	request := PrepareRequestWithoutContext(question, ia)

	return ChatWitModels(request)
}
