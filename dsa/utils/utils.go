package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadStringInput() string {
	var word string
	reader := bufio.NewReader(os.Stdin)
	word, _ = reader.ReadString('\n') // Read the full input line including spaces
	word = strings.TrimSpace(word)
	return word
}

func ToLowerCase(word string) string {
	return strings.ToLower(word)
}

func RemoveWhiteSpace(word string) string{
	noWhitespace := strings.ReplaceAll(word, " ", "")
	return noWhitespace
}

// func RemovePunctuation(word string) string {

// }