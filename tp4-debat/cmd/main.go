package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tp4-debat/agt"
	"tp4-debat/env"
)

func main() {
	debater1 := agt.NewAgentDebater()
	env1 := env.NewEnvironment()

	for {
		fmt.Print("Question: ")
		reader := bufio.NewReader(os.Stdin)
		question, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		debater1.Percept(*env1)
		//debater1.Deliberate()
		debater1.Act(env1, question)
	}
}
