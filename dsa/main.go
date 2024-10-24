package main

import (
	"dsa/utils"
	"fmt"
)

func freqCount(word string) map[string]int {
	counter := make(map[string]int)
	n := len(word)

	for i := 0; i < n; i++ {
		counter[string(word[i])] += 1
	}
	return counter
}



func main() {
	fmt.Println("Enter word: ")
	word := utils.ReadStringInput()

	word1 := utils.RemoveWhiteSpace(word)
	word2 := utils.ToLowerCase(word1)

	fmt.Println(freqCount(word2))
}
