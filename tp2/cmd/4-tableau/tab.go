package main

import (
	"fmt"
	"math"
	v2 "math/rand/v2"
	"sync"
	"time"
)

const gmax int = 10 // nb max de goroutine
const size int = 1 << 20

func Fill(tab []int, v int) {
	for pos := range tab {
		tab[pos] = v
	}
}

// applique f sur chaque élément de tab et remplace la valeur initiale par le résultat de f
func ForEach(tab []int, f func(int) int) {
	for pos, v := range tab {
		tab[pos] = f(v)
	}
}

// copy le tableau src dans dest
func Copy(src []int, dest []int) {
	for pos, val := range src {
		dest[pos] = val
	}
}

func testFill() {
	// Initialisation
	var tab [100000]int

	// Affichage
	fmt.Println("Mode Normal")

	t1 := time.Now()
	Fill(tab[:], 4)
	fmt.Println(time.Since(t1))

	fmt.Println("Mode Concurentielle")

	t2 := time.Now()

	//go ForEach(tab[:], f)

	for i := 0; i < len(tab); i = i + len(tab)/100 {
		go Fill(tab[i:i+len(tab)/100], 4)
	}

	fmt.Println(time.Since(t2))

}

func testForEach() {
	// Initialisation
	var tab [100]int

	f := func(c int) int {
		return int(math.Sqrt(v2.Float64() * 100))
	}

	// Affichage
	fmt.Println("Mode Normal")

	t1 := time.Now()
	ForEach(tab[:], f)
	fmt.Println(time.Since(t1))

	fmt.Println("Mode Concurentielle")

	t2 := time.Now()

	//go ForEach(tab[:], f)

	for i := 0; i < len(tab); i = i + len(tab)/10 {
		//fmt.Println("départ :", i)
		//fmt.Println("fin : ", i+len(tab)/10-1)
		go ForEach(tab[i:i+len(tab)/10], f)
	}

	fmt.Println(time.Since(t2))

}

func correction(tab []int, f func(int) int) {
	var wg sync.WaitGroup

	for i := range gmax {
		ssize := size / gmax
		start := i * ssize
		stop := (i + 1) * ssize

		wg.Add(1)

		go func() {
			ForEach(tab[start:stop], f)
			wg.Done()
		}()
	}

	wg.Add(1)

	go func() {
		ForEach(tab[size%gmax:], f)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	testForEach()
	fmt.Scanln()
}
