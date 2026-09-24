package main

import (
	simulation "simu"
	"time"
)

func main() {
	s := simulation.NewSimulation(100, -1, 600*time.Second)
	go simulation.StartAPI(s)
	s.Run()
}
