package main

import (
	"fmt"
	"time"
)

func Print(countdown int) {
	fmt.Println(countdown)
}

func newYearSleep() {
	for i := 5; i > 0; i-- {
		go Print(i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("...Bonne année")
}

func newYearAfter() {

}

func main() {
	newYearSleep()
	fmt.Scanln()
}
