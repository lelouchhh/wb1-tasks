package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdefss"
	fmt.Println(Uniq(s))
}

func Uniq(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]struct{})
	for _, v := range s {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}
