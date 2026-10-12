package simu

import (
	"tp4-debat/agt"
	"tp4-debat/env"
)

type SimulationMA struct {
	Env  *env.Environment
	Agts []agt.Agent
}

func (simu *SimulationMA) Run() {
	step := 0

	for _, agt := range simu.Agts {
		go func() {
			for {
				agt.Percept(*simu.Env)
				agt.Deliberate()
				agt.Act(simu.Env)
			}
		}()

	}
	step += 1
	//simu.log()   // On sauvegarde l'état courant de la simulation
	//simu.print() // On "affiche le résultat actuel"

}
