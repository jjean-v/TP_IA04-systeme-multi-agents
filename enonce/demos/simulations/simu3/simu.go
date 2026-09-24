package simulation

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Simu struct {
	env         Environment
	agents      []Agent
	maxStep     int
	maxDuration time.Duration
	step        int // Stats
	start       time.Time
	syncChans   sync.Map
}

func NewSimulation(agentCount int, maxStep int, maxDuration time.Duration) (simu *Simu) {
	simu = &Simu{}
	simu.agents = make([]Agent, 0, agentCount)
	simu.maxStep = maxStep
	simu.maxDuration = maxDuration

	simu.env = *NewEnvironment([]Agent{})

	// création des agents et des channels
	for i := 0; i < agentCount; i++ {
		// création de l'agent
		id := fmt.Sprintf("Agent #%d", i)
		syncChan := make(chan int)
		ag := NewAgentPI(id, &simu.env, syncChan)

		// ajout de l'agent à la simulation
		simu.agents = append(simu.agents, ag)

		// ajout du channel de synchro
		simu.syncChans.Store(ag.ID(), syncChan)

		// ajout de l'agent à l'environnement
		ag.env.AddAgent(ag)
	}

	return simu
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
	simu.start = time.Now()

	// Boucle principale
Loop:
	for {
		simu.step += 1

		for _, agt := range simu.agents {
			go func(agt Agent) {
				time.Sleep(10 * time.Millisecond) // "cool down"
				c, _ := simu.syncChans.Load(agt.ID())
				c.(chan int) <- simu.step // /!\ utilisation d'un "Type Assertion"
			}(agt)
		}

		if (simu.maxStep > 0 && simu.step >= simu.maxStep) || time.Since(simu.start) > simu.maxDuration {
			fmt.Println()
			break Loop
		}

		time.Sleep(10 * time.Millisecond)
	}

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
