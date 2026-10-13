package agt

import (
	"encoding/json"
	"log"
	"tp4-debat/com"
	"tp4-debat/env"
)

type Juge struct {
	jugeChan chan string
	debat    *com.MessageRequest
	prompt   string
}

func NewJuge(cjuge chan string) *Juge {
	return &Juge{jugeChan: cjuge}
}

func (j *Juge) Start(prompt string) {
	j.prompt = prompt
}

func (j *Juge) Percept(e env.Environment) {
	<-j.jugeChan // Listen for call
	log.Println("Le juge délibère")
	ancientPrompt := e.Read()
	j.debat = &com.MessageRequest{Role: "user", Content: string(ancientPrompt)}
}

func (j *Juge) Deliberate() {
	log.Println("Le juge délibère encore")

	listQuestion := []com.MessageRequest{{Role: "user", Content: j.prompt}, *j.debat}

	request := com.Request{Model: MODEL_IA_GPT_OSS_20B, Messages: listQuestion} // Prepare request

	buffer, errMarshal2 := json.Marshal(request) // Transform into []byte

	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	// Send question
	result := com.ChatWitModels(buffer) // type ReponseRequest

	if len(result.Choices) == 0 {
		log.Println("Erreur Appel API")
		log.Fatal(result.Error.Message)
	}

	var response string

	for _, element := range result.Choices {
		response = element.Message.Content
	}

	j.jugeChan <- response

}

func (j *Juge) Act(e *env.Environment) {

}
