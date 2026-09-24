package simulation

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Simulation struct {
	env         Environment // pb de copie avec les locks...
	agents      []Agent
	maxStep     int
	maxDuration time.Duration
	step        int // Stats
	start       time.Time
	syncChans   sync.Map
}

func NewSimulation(agentCount int, maxStep int, maxDuration time.Duration) (simu *Simulation) {
	simu = &Simulation{}
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

func (simu *Simulation) Run() {
	log.Printf("Démarrage de la simulation [step: %d, π: %f]", simu.step, simu.env.PI())

	// Démarrage du micro-service de Log
	go simu.Log()
	// Démarrage du micro-service d'affichage
	go simu.Print()

	// Démarrage des agents
	for i := range simu.agents {
		simu.agents[i].Start() // et pas agt.Start() !
	}

	// On sauvegarde la date du début de la simulation
	simu.start = time.Now()

	// Lancement de l'orchestration de tous les agents
	// simu.step += 1 // plus de sens
	for _, agt := range simu.agents {
		go func(agt Agent) {
			step := 0
			for {
				step++
				c, _ := simu.syncChans.Load(agt.ID())
				c.(chan int) <- step             // /!\ utilisation d'un "Type Assertion"
				time.Sleep(1 * time.Millisecond) // "cool down", 1000 tps max...
				<-c.(chan int)
			}
		}(agt)
	}

	time.Sleep(simu.maxDuration)

	log.Printf("Fin de la simulation [step: %d, in: %d, out: %d, π: %f]", simu.step, simu.env.in, simu.env.out, simu.env.PI())
}

func (simu *Simulation) Print() {
	for {
		fmt.Printf("\rπ = %.30f", simu.env.PI())
		time.Sleep(time.Second / 60) // 60 fps !
	}
}

func (simu *Simulation) Log() {
	// Not implemented
}
