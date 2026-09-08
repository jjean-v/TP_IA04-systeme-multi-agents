package main

import "fmt"
import "sort"

var dict = [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}

func IsPalindrome(word string) bool {
	rune_right := len(word) - 1
	for rune_left := 0; rune_left < len(word)/2; rune_left++ {
		if word[rune_left] != word[rune_right] {
			return false
		}
		rune_right--
	}

	return true
}

func Palindromes(words []string) (l []string) {
	for _, word := range words {
		if IsPalindrome(word) {
			l = append(l, word)
		}
	}
	return l
}

func Footprint(s string) (footprint string) {
	n := []rune(s)
	sort.Slice(n, func(i int, j int) bool { return n[i] < n[j] })
	footprint = string(n)
	return footprint
}

func Anagrams(words []string) (anagrams map[string][]string) {
	var empreinte string 
	anagrams = make(map[string][]string)
	
	for _,word := range words {
		empreinte = Footprint(word)
		if _, exist :=anagrams[empreinte]; exist{
			anagrams[empreinte] = append(anagrams[empreinte],word)
		} else {
			var sl []string
			sl = append(sl,word)
			anagrams[empreinte] = sl
		}
		
	}
		

	return anagrams
}

func main() {
	fmt.Println("_______ Function IsPalindrome _______")
	fmt.Println(IsPalindrome("RADAR"))
	fmt.Println(IsPalindrome("AGENT"))

	fmt.Println("_______ Function Palindromes _______")
	fmt.Println(Palindromes(dict[:])) // ON passe un slice pour éviter de copier le dico entier sur la pile

	fmt.Println("_______ Function Footprint _______")
	fmt.Println(Footprint("AGENT"))

	fmt.Println("_______ Function Anagrams _______")
	fmt.Println(Anagrams(dict[:]))

}
