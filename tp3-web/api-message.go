package tp3web

import (
	"encoding/json"
	"log"
)

const CONTEXT_PROMPT = `You are receiving a series of reasoning traces extracted from previous AI responses. 
	Each trace represents the internal chain of thought the model produced before giving its final answer.

	Your task is to analyze these traces and use it as a context.

	The traces will be provided as a list in the next message, in this format:
	- reasoning : "<reasoning text>"`

const URL_MODEL_LIST = "https://api.groq.com/openai/v1/models"
const URL_MODEL_QUESTIONS = "https://api.groq.com/openai/v1/chat/completions"

type MessageRequest struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string           `json:"model"`
	Messages []MessageRequest `json:"messages"`
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

func JsonResponse(body []byte, result any) any {

	err := json.Unmarshal(body, result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func PrepareRequestWithContext(question string, ia string) []byte {

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
	return buffer

}

func PrepareRequestWithoutContext(question string, ia string) []byte {
	prompt := `You are receiving a tuple, response and message, extracted from previous AI responses. 
		your goal is to summarize this into a direct message that the AI will then be able to use to have access to useful information about the previous discussion.
		I want yout to keep only the important information, who i am, what I'm looking for, ... Your answer must be pretty short few sentences.`

	// Prepare the request

	// Add actual question
	message := []MessageRequest{{"user", prompt}, {"user", question}}

	request := Request{ia, message}

	buffer, errMarshal2 := json.Marshal(request)
	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}
	return buffer
}
