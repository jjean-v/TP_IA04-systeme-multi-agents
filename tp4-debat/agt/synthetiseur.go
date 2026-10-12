package agt

import (
	"encoding/json"
	"log"
	"tp4-debat/com"
	"tp4-debat/env"
)

type ChanMessageSynthetiseur struct {
	message  string
	Crequest chan string
}

type Synthetiseur struct {
	id               string
	synthetiseurChan chan ChanMessageSynthetiseur
	prompt           string
	request          ChanMessageSynthetiseur
	resume           string
}

func NewSynthetiseur(name string, c chan ChanMessageSynthetiseur) *Synthetiseur {
	return &Synthetiseur{id: name, synthetiseurChan: c}
}

func (s *Synthetiseur) Start(prompt string) {
	s.prompt = prompt
}

func (s *Synthetiseur) Percept(e env.Environment) {
	s.request = <-s.synthetiseurChan
}

func (s *Synthetiseur) Deliberate() {
	listQuestion := []com.MessageRequest{{Role: "user", Content: s.prompt}, {Role: "user", Content: s.request.message}}

	request := com.Request{Model: MODEL_IA_GPT_OSS_20B, Messages: listQuestion} // Prepare request

	buffer, errMarshal2 := json.Marshal(request) // Transform into []byte

	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	// Send question
	result := com.ChatWitModels(buffer) // type ReponseRequest

	for _, element := range result.Choices {
		s.resume = element.Message.Content
	}

}

func (s *Synthetiseur) Act(e *env.Environment) {
	s.request.Crequest <- s.resume

}
