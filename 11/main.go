package main

import "fmt"

func main() {
	setA := make(map[int]struct{})
	setA = map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}
	setB := make(map[int]struct{})
	setB = map[int]struct{}{
		3: {},
		4: {},
		5: {},
	}
	fmt.Println(intersection(setB, setA))
}

func intersection(a, b map[int]struct{}) map[int]struct{} {
	m := make(map[int]struct{})
	for k, _ := range b {
		if _, ok := a[k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}
