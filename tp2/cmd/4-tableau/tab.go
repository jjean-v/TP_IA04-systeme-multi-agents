package main

import (
	"fmt"
	"math"
	v2 "math/rand/v2"
	"sync"
	"time"
)

const gmax int = 4 // nb max de goroutine
const size int = 1 << 27

var tab [size]int

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

func FillGo(tab []int, v int) {
	var wg sync.WaitGroup

	for i := range gmax {
		ssize := size / gmax
		start := i * ssize
		stop := (i + 1) * ssize

		wg.Add(1)

		go func() {
			Fill(tab[start:stop], v)
			defer wg.Done()
		}()
	}

	wg.Add(1)

	go func() {
		Fill(tab[size%gmax:], v)
		defer wg.Done()
	}()

	wg.Wait()
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

func operation(c int) int {
	return int(math.Sqrt(v2.Float64() * 100))
}

func main() {

	//testForEach()

	fmt.Print("Mode Normal : ")

	t1 := time.Now()
	Fill(tab[:], 3)

	fmt.Println(time.Since(t1))

	fmt.Print("Mode Concurentielle : ")

	t2 := time.Now()
	FillGo(tab[:], 3)
	fmt.Println(time.Since(t2))

}
