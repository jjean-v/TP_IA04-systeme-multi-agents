package simulation

import (
	"fmt"
	"log"
	"time"
)

type Simu struct {
	env         Environment
	agents      []Agent
	maxStep     int
	maxDuration time.Duration
	step        int // Stats
}

func NewSimulation(agentCount int, maxStep int, maxDuration time.Duration) *Simu {
	ags := make([]Agent, 0, agentCount)

	for i := 0; i < agentCount; i++ {
		id := fmt.Sprintf("Agent #%d", i)
		ag := NewAgentPI(id)
		ags = append(ags, ag)
	}
	env := *NewEnvironment(ags)

	return &Simu{env, ags, maxStep, maxDuration, 0}
}

func (simu *Simu) Run() {
	log.Printf("Démarrage de la simulation [step: %d, π: %f]", simu.step, simu.env.PI())

	// Démarrage des agents
	for _, agt := range simu.agents {
		agt.Start()
	}

	// On sauvegarde la date du début de la simulation
	start := time.Now()

	// Boucle de simulation
Loop:
	for {
		simu.step++
		for _, agt := range simu.agents {
			agt.Percept(&simu.env)
			agt.Deliberate()
			agt.Act(&simu.env)

			if (simu.maxStep > 0 && simu.step >= simu.maxStep) || time.Since(start) > simu.maxDuration {
				fmt.Println()
				break Loop
			}
		}

		simu.Log()                   // On sauvegarde l'état courant de la simulation
		simu.Print()                 // On "affiche le résultat actuel"
		time.Sleep(time.Second / 60) // 60 fps !
	}

	log.Printf("Fin de la simulation [step: %d, in: %d, out: %d, π: %f]", simu.step, simu.env.in, simu.env.out, simu.env.PI())
}

func (simu *Simu) Print() {
	fmt.Printf("\rπ = %.30f", simu.env.PI())
}

func (simu *Simu) Log() {
	// not implemented
}
