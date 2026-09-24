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
	env      *Environment
	syncChan chan int
}

func NewAgentPI(id string, env *Environment, syncChan chan int) *AgentPI {
	return &AgentPI{AgentID(id), Rect{Coord{0, 0}, Coord{1, 1}}, Noop, env, syncChan}
}

func (ag *AgentPI) ID() AgentID {
	return ag.id
}

func (ag *AgentPI) Start() {
	log.Printf("%s starting...\n", ag.id)

	go func() {
		env := ag.env
		var step int
		for {
			step = <-ag.syncChan

			ag.Percept(env)
			ag.Deliberate()
			ag.Act(env)

			ag.syncChan <- step
		}
	}()
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
