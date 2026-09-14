package main

import "fmt"

func compte(n int) {
	for i :=0; i < n; i++ {
		fmt.Println(i)
	}
}


func compteWithGoRoutine(n int) {
	for i :=0; i < n; i++ {

		go Print(i)

		
	}
}

func Print(i int) {
	fmt.Println(i)
}

func main() {
	compteWithGoRoutine(10)
	fmt.Scanln()
}