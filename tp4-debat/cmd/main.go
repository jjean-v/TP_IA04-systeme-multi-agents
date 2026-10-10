package main

import (
	"fmt"
	"log"
	"tp4-debat/agt"
	"tp4-debat/env"
)

func main() {

	chanModo := make(chan agt.ChanMessage)

	moderateur := agt.NewModerateur("modo1", chanModo)
	debater1 := agt.NewAgentDebater("agent1", chanModo)
	debater2 := agt.NewAgentDebater("agent2", chanModo)

	env1 := env.NewEnvironment()

	go func() {
		log.Println("Lancement du modérateur")
		moderateur.Percept()

	}()

	go func() {
		for {
			debater1.Percept(*env1)
			debater1.Deliberate()
			debater1.Act(env1)
		}
	}()

	go func() {
		for {
			debater2.Percept(*env1)
			debater2.Deliberate()
			debater2.Act(env1)
		}
	}()

	fmt.Scanln()
	/*
		for {
			go func() {
				moderateur.Percept()
			}()
			go func() {
				debater1.Percept(*env1)
				debater1.Deliberate()
				debater1.Act(env1)
			}()
			go func() {
				debater2.Percept(*env1)
				debater2.Deliberate()
				debater2.Act(env1)
			}()

	*/
	/*
		fmt.Print("Question: ")
		reader := bufio.NewReader(os.Stdin)
		question, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		debater1.Percept(*env1)
		//debater1.Deliberate()
		debater1.Act(env1, question)
	*/

}
