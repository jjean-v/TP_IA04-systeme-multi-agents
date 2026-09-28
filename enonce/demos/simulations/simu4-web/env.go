package simulation

import (
	"fmt"
	"sync"
)

type Environment struct {
	sync.RWMutex
	ags        []Agent
	agentCount int
	in         uint64
	out        uint64
	noopCount  uint64
}

func NewEnvironment(ags []Agent) (env *Environment) {
	return &Environment{ags: ags, agentCount: len(ags)}
}

func (env *Environment) AddAgent(agt Agent) {
	env.ags = append(env.ags, agt)
	env.agentCount++
}

func (env *Environment) Do(a Action, c Coord) (err error) {
	env.Lock()
	defer env.Unlock()

	switch a {
	case Mark:
		if c[0] < 0 || c[0] > 1 || c[1] < 0 || c[1] > 1 {
			return fmt.Errorf("bad coordinates (%f,%f)", c[0], c[1])
		}

		if c[0]*c[0]+c[1]*c[1] <= 1 {
			env.in++
		} else {
			env.out++
		}
		return nil

	case Noop:
		env.noopCount++
		return nil
	}

	return fmt.Errorf("bad action number %d", a)
}

func (env *Environment) PI() float64 {
	env.RLock()
	defer env.RUnlock()

	return 4 * float64(env.in) / (float64(env.out) + float64(env.in))
}

func (env *Environment) Rect() Rect {
	return Rect{Coord{0, 0}, Coord{1, 1}}
}
