package simulation

import (
	"log"
	"math/rand"
)

type Action int64

const (
	Noop = iota
	Mark
)

type Coord [2]float64

type Rect [2]Coord

type Agent interface {
	Start()
	Percept(*Environment)
	Deliberate()
	Act(*Environment)
	ID() AgentID
}

type AgentID string

type AgentPI struct {
	id       AgentID
	rect     Rect
	decision Action
}

func NewAgentPI(id string) *AgentPI {
	return &AgentPI{AgentID(id), Rect{Coord{0, 0}, Coord{1, 1}}, Noop}
}

func (ag *AgentPI) ID() AgentID {
	return ag.id
}

func (ag *AgentPI) Start() {
	log.Printf("%s starting...\n", ag.id)
}

func (ag *AgentPI) Percept(env *Environment) {
	ag.rect = env.Rect()
}

func (ag *AgentPI) Deliberate() {
	if rand.Float64() < 0.1 {
		ag.decision = Noop
	} else {
		ag.decision = Mark
	}
}

func (ag *AgentPI) Act(env *Environment) {
	if ag.decision == Noop {
		env.Do(Noop, Coord{})
	} else {
		x := rand.Float64()*(ag.rect[1][0]-ag.rect[0][0]) + ag.rect[0][0]
		y := rand.Float64()*(ag.rect[1][1]-ag.rect[0][1]) + ag.rect[0][1]
		env.Do(Mark, Coord{x, y})
	}
}
