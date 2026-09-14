package main

import "fmt"

func compte(n int) {
	for i :=0; i < n; i++ {
		fmt.Println(i)
	}
}


func compteWithGoRoutine(n int) {
	for i :=0; i < n; i++ {

		go fmt.Println(i)
	}
}


func compteMsg(n int, msg string) {
	for i :=0; i < n; i++ {
		fmt.Println(msg)
		fmt.Println(i)	
	}
}

func compteMsgFromTo(start int, end int, msg string) {
	for i := start; i < end; i++ {
		fmt.Println(msg, i)
	}
}

func testCompteMsg() {
	go compteMsg(5, "func1")
	go compteMsg(5, "func2")
}

func testCompteMsgFromTo(start int, end int, msg string) {
	for i := 0; i <= 90; i +=10 {
		go compteMsgFromTo(i,i+10,"hey")
	}
}

func main() {
	
	//compteWithGoRoutine(5)
	testCompteMsg()
	fmt.Scanln()
}