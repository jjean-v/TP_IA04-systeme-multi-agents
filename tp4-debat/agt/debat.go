package agt

import (
	"encoding/json"
	"fmt"
	"log"
	"tp4-debat/com"
	"tp4-debat/env"
)

const MODEL_IA_GPT_OSS_20B = "openai/gpt-oss-20b"
const MODEL_IA_GPT_OSS_SAFEGUARD_20B = "openai/gpt-oss-safeguard-20b"
const MODEL_IA_ALIBABA_QWEN_3_8_27B = "qwen/qwen3.8-27b"
const MODEL_IA_SDAIA_ALLAM_2_7B = "allam-2-7b"

type Agent interface {
	Start()
	Percept(env.Environment)
	Deliberate()
	Act(*env.Environment, string)
}

type Debater struct {
	//chanRequest chan string
	discussion *com.MessageRequest
}

func NewAgentDebater() *Debater {
	return &Debater{}
}

func (d *Debater) Percept(env env.Environment) {
	ancientPrompt := env.Read()
	d.discussion = &com.MessageRequest{"user", string(ancientPrompt)}
}

func (d *Debater) Act(envReal *env.Environment, question string) {
	storeType := env.Question
	envReal.Write(question, storeType)
	// Add the question to the prompt
	listQuestion := []com.MessageRequest{*d.discussion, {"user", "Actual question: " + question}}

	// Prepare request
	request := com.Request{MODEL_IA_GPT_OSS_20B, listQuestion}
	// Transform into []byte
	buffer, errMarshal2 := json.Marshal(request)
	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	// Send question
	result := com.ChatWitModels(buffer) // type ReponseRequest
	storeType = env.Answer
	for _, element := range result.Choices {
		envReal.Write(element.Message.Content, storeType)
		fmt.Println(element.Message.Content)
	}

}
