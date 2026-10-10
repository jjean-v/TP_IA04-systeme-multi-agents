package agt

import (
	"log"
	"time"
)

type Moderateur struct {
	id           string
	count        int
	cqueue       chan ChanMessage
	lastResponse string
}

type ChanMessage struct {
	AgentId  string
	Crequest chan string
}

func NewModerateur(name string, c chan ChanMessage) *Moderateur {
	return &Moderateur{id: name, cqueue: c}
}

func (m *Moderateur) Percept() {
	for {
		s := <-m.cqueue
		switch s.AgentId {
		case "agent1":
			log.Println("Tour de Agent 1")
			s.Crequest <- m.lastResponse
			m.lastResponse = <-s.Crequest
		case "agent2":
			log.Println("Tour de Agent 2")
			s.Crequest <- m.lastResponse
			m.lastResponse = <-s.Crequest
		}
		m.count++
		log.Println("Receive: ", m.lastResponse)
		log.Println("Count: ", m.count)
		time.Sleep(3 * time.Second)
	}

}
