package agt

import (
	"env"
	"tp3web"
)

type Agent interface {
	Start()
	Percept(env.Environment)
	Deliberate()
	Act(*env.Environnent)
}

type Debater struct {
	chanRequest chan string
	Discussion  *tp3web.MessageRequest
}

func (d *Debater) Percept(env env.Environnement) {
	ancientPrompt := env.Read()
	d.Discussion = tp3web.MessageRequest{"user", string(ancientPrompt)}
}

func (d *Debater) Act(env *env.Environment) {

}
