package simulation

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Simu struct {
	env         Environment // pb de copie avec les locks...
	agents      []Agent
	maxStep     int
	maxDuration time.Duration
	step        int // Stats
	start       time.Time
}

func NewSimulation(agentCount int, maxStep int, maxDuration time.Duration) *Simu {
	ags := make([]Agent, 0, agentCount)

	for i := 0; i < agentCount; i++ {
		id := fmt.Sprintf("Agent #%d", i)
		ag := NewAgentPI(id)
		ags = append(ags, ag)
	}

	return &Simu{env: *NewEnvironment(ags), agents: ags, maxStep: maxStep, maxDuration: maxDuration}
}

func (simu *Simu) Run() {
	log.Printf("Démarrage de la simulation [step: %d, π: %f]", simu.step, simu.env.PI())

	// Démarrage du micro-service de Log
	go simu.Log()
	// Démarrage du micro-service d'affichage
	go simu.Print()

	// Démarrage des agents
	for _, agt := range simu.agents {
		agt.Start()
	}

	// On sauvegarde la date du début de la simulation
	start := time.Now()

	// Boucle de simulation
Loop:
	for {
		var wg sync.WaitGroup
		simu.step++

		for _, agt := range simu.agents {
			wg.Add(1)

			go func(agt Agent) { // workers
				defer wg.Done()
				agt.Percept(&simu.env)
				agt.Deliberate()
				agt.Act(&simu.env)
			}(agt)
		}
		wg.Wait()

		if (simu.maxStep > 0 && simu.step >= simu.maxStep) || time.Since(start) > simu.maxDuration {
			fmt.Println()
			break Loop
		}
	}
	time.Sleep(time.Second / 120) // 120 tps !

	log.Printf("Fin de la simulation [step: %d, in: %d, out: %d, π: %f]", simu.step, simu.env.in, simu.env.out, simu.env.PI())
}

func (simu *Simu) Print() {
	for {
		fmt.Printf("\rπ = %.30f", simu.env.PI())
		time.Sleep(time.Second / 60) // 60 fps !
	}
}

func (simu *Simu) Log() {
	// Not implemented
}
