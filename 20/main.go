package main

import (
	"fmt"
	"strings"
)

func reverseWordsInString(input string) string {
	words := strings.Fields(input)
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[i] = reverseString(word)
	}

	return strings.Join(reversedWords, " ")
}

func reverseString(input string) string {
	runes := []rune(input)
	reversedRunes := make([]rune, len(runes))

	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		reversedRunes[i], reversedRunes[j] = runes[j], runes[i]
	}

	return string(reversedRunes)
}

func main() {
	inputString := "snow dog sun — sun dog snow"
	reversedResult := reverseWordsInString(inputString)
	fmt.Println(reversedResult)
}
