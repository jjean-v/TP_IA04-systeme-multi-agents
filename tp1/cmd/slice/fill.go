package main

import "fmt"
import "math/rand"
import "slices"

func Fill(sl []int) []int {
	for i, _ := range sl {
		sl[i] = rand.Int()
	}
	return sl

}

func Moyenne(sl []int) float64 {
	var mean float64 = 0.0
	for _, v := range sl {
		mean += float64(v)
	}
	var size int = len(sl)
	mean = mean / float64(size)
	return mean
}

func ValeursCentrales(sl []int) []int {
	// Test if Slice empty
	size := len(sl)
	if size == 0 {
		fmt.Println("Slice empty")
		return sl
	}

	var tab [2]int

	SizeCmp := func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}
	slices.SortFunc(sl, SizeCmp)

	if size%2 != 0 {
		result := tab[:1]
		result[0] = sl[(size-1)/2]
		return result
	} else {
		fmt.Println("Hello")
		result := tab[:2]
		result[0] = sl[(size/2)-1]
		result[1] = sl[size/2]
		return result
	}

}

func TestValeursCentrales() {
	var list = []int{10, 3, 9, 20, 3, 2, 2}
	// creation du slice
	sl := list[:]
	fmt.Println(ValeursCentrales(sl))
	fmt.Println(sl)

}

func Plus1(sl []int) {

	for i, _ := range sl {
		sl[i] += 1
	}
}

func Compte(sl []int) {
	for _, v := range sl {
		fmt.Println(v)
	}
}

func TestFill() {
	var list [1000]int
	// creation du slice
	sl := list[:]

	fmt.Println(list)
	sl = Fill(sl)
	fmt.Println(sl)
}

func TestMoyenne() {
	var list = []int{10, 3, 2, 2}
	// creation du slice
	sl := list[:]

	var mean = Moyenne(sl)
	fmt.Println(mean)
}

func TestPlus1() {
	var list = []int{10, 3, 2, 2}
	// creation du slice
	sl := list[:]

	fmt.Println(sl)
	Plus1(sl)
	fmt.Println(sl)

}

func TestCompte() {
	var list = []int{10, 3, 2, 2}
	// creation du slice
	sl := list[:]
	Compte(sl)

}

func main() {

	TestCompte()

}
