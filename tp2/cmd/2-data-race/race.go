package main

import (
	"fmt"
)

/*
var locker sync.Mutex
var n = 0


func f() {
	locker.Lock()
	defer locker.Unlock()
	n++
}

func main() {

	for i := 0; i < 10000; i++ {
		go f()

	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
}
*/

var n = 0

func microService(c chan int) {
	for {
		<-c
		n++
	}
}

func inc(c chan int) {
	c <- 1
}

func main() {
	c := make(chan int)
	go microService(c)

	for i := 0; i < 10000; i++ {
		// Envoie dans le channel
		go inc(c)
	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
}
