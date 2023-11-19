package main

import (
	"fmt"
)

func reverseString(input string) string {
	runes := []rune(input)
	reversedRunes := make([]rune, len(runes))

	for i, j := 0, len(runes)-1; i <= j; i, j = i+1, j-1 {
		reversedRunes[i], reversedRunes[j] = runes[j], runes[i]
	}

	return string(reversedRunes)
}

func main() {
	inputString := "ты"
	reversedResult := reverseString(inputString)
	fmt.Println(reversedResult)
}
