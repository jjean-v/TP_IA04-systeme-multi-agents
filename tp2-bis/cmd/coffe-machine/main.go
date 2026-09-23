package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Etudiant struct {
	id int
}

type Machine struct {
	nb_coffe int
	creq     chan int
	wg       *sync.WaitGroup
}

func (e *Etudiant) TakeCoffee(c chan int) {
	c <- e.id

}

func (m *Machine) ServeCoffe() {
	var id int
	go func() {
		for i := 0; i < 5; i++ {
			id = <-m.creq
			fmt.Printf("Preparation du cafe de l'étudiant : %d\n", id)
			time.Sleep(time.Duration(3+rand.Intn(4)) * time.Second) // n := a + rand.Intn(b-a+1) => [a:b]
			fmt.Printf("Cafe terminé de l'étudiant : %d\n", id)
			m.nb_coffe++
			m.wg.Done()
		}
	}()
}

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	m := Machine{creq: c, wg: &wg}

	t := time.Now()
	for i := 1; i < 6; i++ {
		wg.Add(1)
		e := Etudiant{i}
		go e.TakeCoffee(c)
	}

	m.ServeCoffe()

	wg.Wait()

	fmt.Println("total time : ", time.Since(t))
}
