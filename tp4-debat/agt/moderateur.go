package agt

import (
	"fmt"
	"log"
)

type chanMessage struct {
	AgentId  string
	Crequest chan string
}

type Moderateur struct {
	id     string
	count  int
	cqueue chan chanMessage
}

func NewModerateur(name string, c chan chanMessage) *Moderateur {
	return &Moderateur{id: name, cqueue: c}
}

func (m *Moderateur) Percept() {
	s := <-m.cqueue
	switch s.AgentId {
	case "agent 1":
		log.Println("Tour de Agent1")
		s.Crequest <- "Your turn"
		<-s.Crequest
		fmt.Println("Agent 1 a terminé")

	case "agent 2":
		log.Println("Tour de Agent2")
		s.Crequest <- "Your turn"
		<-s.Crequest
		fmt.Println("Agent 2 a terminé")
	}

}
