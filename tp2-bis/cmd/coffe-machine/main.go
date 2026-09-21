package main

import (
	"math/rand"
	"time"
)

type Etudiant struct {
	id     int
	ccoffe chan int
}

type Machine struct {
	nb_coffe int
	creq     chan int
}

func (e *Etudiant) TakeCoffee() {
	e.ccoffe <- e.id
	time.Sleep(time.Duration(15+rand.Intn(16)) * time.Second) // n := a + rand.Intn(b-a+1) => [a:b]

}
