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
		case "Robert":
			log.Println("Tour de Robert")
			s.Crequest <- m.lastResponse
			m.lastResponse = <-s.Crequest
		case "Dr Stone":
			log.Println("Tour de Dr Stone")
			s.Crequest <- m.lastResponse
			m.lastResponse = <-s.Crequest
		}
		m.count++
		log.Println("Count: ", m.count)
		time.Sleep(3 * time.Second)
	}

}
