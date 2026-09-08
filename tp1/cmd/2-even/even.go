package main

import "fmt"

func pair() {
	for i := 0; i <= 1000; i+=2 {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Les nombres pairs de 0 à 1000 sont:")
	pair()

}
