package agt

import (
	"encoding/json"
	"log"
	"tp4-debat/com"
	"tp4-debat/env"
)

const MODEL_IA_GPT_OSS_20B = "openai/gpt-oss-20b"
const MODEL_IA_GPT_OSS_SAFEGUARD_20B = "openai/gpt-oss-safeguard-20b"
const MODEL_IA_ALIBABA_QWEN_3_8_27B = "qwen/qwen3.8-27b"
const MODEL_IA_SDAIA_ALLAM_2_7B = "allam-2-7b"

type Agent interface {
	Start(string)
	Percept(env.Environment)
	Deliberate()
	Act(*env.Environment)
}

type Debater struct {
	id            string
	prompt        string
	discussion    *com.MessageRequest
	Crequest      chan ChanMessage
	Creceive      chan string
	CSynthetiseur chan ChanMessageSynthetiseur
}

func NewAgentDebater(name string, c chan ChanMessage, cSynthetiseur chan ChanMessageSynthetiseur) *Debater {
	chanReceive := make(chan string)
	return &Debater{id: name, Crequest: c, Creceive: chanReceive, CSynthetiseur: cSynthetiseur}
}

func (d *Debater) Start(prompt string) {
	d.prompt = prompt
}

func (d *Debater) Percept(env env.Environment) {
	ancientPrompt := env.Read()
	d.discussion = &com.MessageRequest{Role: "user", Content: string(ancientPrompt)}
}

func (d *Debater) Deliberate() {

	request := ChanMessage{AgentId: d.id, Crequest: d.Creceive}
	log.Println(d.id + " envoie une demande de parler")
	d.Crequest <- request // Envoi de la demande d'action

}

func (d *Debater) Act(envReal *env.Environment) {
	message := <-d.Creceive
	log.Println("Message receive")

	listQuestion := []com.MessageRequest{{Role: "user", Content: d.prompt}, *d.discussion, {Role: "user", Content: "Actual Message: " + message}} // Add the question to the prompt

	request := com.Request{Model: MODEL_IA_GPT_OSS_20B, Messages: listQuestion} // Prepare request

	buffer, errMarshal2 := json.Marshal(request) // Transform into []byte

	if errMarshal2 != nil {
		log.Fatal(errMarshal2)
	}

	result := com.ChatWitModels(buffer) // Send question

	var reponse string
	for _, element := range result.Choices {
		reponse = element.Message.Content
	}

	if len(reponse) > 700 {
		c := make(chan string)
		log.Printf("Message trop long: %d", len(reponse))

		d.CSynthetiseur <- ChanMessageSynthetiseur{reponse, c}
		reponse = <-c
	}

	d.Creceive <- reponse

}
