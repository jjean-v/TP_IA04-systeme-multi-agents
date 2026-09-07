package main

import "fmt"

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

func main() {
	fmt.Println(IsPalindrome("RADAR"))
	fmt.Println(IsPalindrome("AGENT"))
	fmt.Println(Palindromes(dict[:])) // ON passe un slice pour éviter de copier le dico entier sur la pile

}
