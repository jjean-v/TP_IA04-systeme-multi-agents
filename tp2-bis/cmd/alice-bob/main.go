package main

import (
	"abr"
	"fmt"
	"sync"
)

func main() {
	c := make(chan string)
	var wg1 sync.WaitGroup
	wg1.Add(2)

	Alice := *abr.NewPersonne("Alice", &wg1)
	Bob := *abr.NewPersonne("Bob", &wg1)

	Maison := abr.House{}
	go Maison.StartEnvironnement(c)

	Alice.Start(c)
	Bob.Start(c)

	wg1.Wait()
	fmt.Println("Alice et Bob ont fini de se préparer")

	Maison.SetAlarm()

	wg1.Add(2)
	Alice.Partir()
	Bob.Partir()

	wg1.Wait()
	fmt.Println("Fermeture et verrouillage de la porte")

	fmt.Scanln()

}
