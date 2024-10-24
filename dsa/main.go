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

func isPalindrome (word string) bool {
	i:=0
	j:=(len(word) - 1)

	for i <= j {
		if (word[i] == word[j]){
			i += 1
			j -= 1
		} else{
			return false
		}
	}
	return true


}

func main() {
	fmt.Println("Enter word to count its character's count: ")
	word := utils.ReadStringInput()

	word1 := utils.RemoveWhiteSpace(word)
	word2 := utils.ToLowerCase(word1)

	fmt.Println(freqCount(word2))

	// Palindrome Checker
	fmt.Println("Enter a word to check if palindrone: ")

	palindrome := utils.ReadStringInput()
	palindrome1 := utils.RemoveWhiteSpace(palindrome)
	palindrome2 := utils.ToLowerCase(palindrome1)
	
	fmt.Println(isPalindrome(palindrome2))
}
