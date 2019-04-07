// Report whether 2 strings are anagrams of each other (contain the same letters in different order)

package main

import (
	"fmt"
	"os"
	"strings"
)

func anagrams(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	for i := range word1 {
		if strings.Count(word2, string(word1[i])) != strings.Count(word1, string(word1[i])) {
			return false
		}
	}
	return true
}

func main() {
	arg := os.Args[1:]
	if len(arg) < 2 {
		fmt.Println("NEED 2 WORDS!")
		return
	}
	word1 := arg[0]
	word2 := arg[1]
	fmt.Printf("Checking if %s and %s are anagrams: ", word1, word2)
	fmt.Println(anagrams(word1, word2))

}
