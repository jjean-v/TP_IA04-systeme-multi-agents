package agt

import (
	"log"
	"time"
	"tp4-debat/env"
)

type Moderateur struct {
	id           string
	count        int
	cqueue       chan ChanMessage
	cjuge        chan string
	lastResponse string
}

type ChanMessage struct {
	AgentId  string
	Crequest chan string
}

func NewModerateur(name string, c chan ChanMessage, cjuge chan string, nbtour int) *Moderateur {
	return &Moderateur{id: name, cqueue: c, cjuge: cjuge, count: nbtour}
}

func (m *Moderateur) Act(e *env.Environment) {
	for m.count > 0 {
		s := <-m.cqueue
		switch s.AgentId {
		case "Robert":
			log.Println("Tour de Robert")
			s.Crequest <- m.lastResponse
			m.lastResponse = <-s.Crequest
			e.Write("Robert : " + m.lastResponse)
		case "Dr Stone":
			log.Println("Tour de Dr Stone")
			s.Crequest <- m.lastResponse
			m.lastResponse = <-s.Crequest
			e.Write("Dr Stone : " + m.lastResponse)

		}
		m.count--
		time.Sleep(5 * time.Second)
	}
	log.Printf("Count : %d \n", m.count)
	log.Println("Appel de juge")
	m.cjuge <- "Go"
	reponse := <-m.cjuge
	log.Println(reponse)
	e.Write("Juge :" + reponse)

}
