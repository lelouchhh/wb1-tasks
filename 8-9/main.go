package main

import (
	"fmt"
	"sync"
)

func main() {
	//8
	n := 0
	i := 1
	fmt.Println(n | (1 << i))
	globalChan := make(chan int)
	//9
	var s []int
	s = []int{1, 2, 3, 4, 5}
	wg := &sync.WaitGroup{}
	for _, v := range s {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			globalChan <- v
		}(v)
	}
	go func() {
		for v := range globalChan {
			fmt.Println(v * v)
		}
	}()
	wg.Wait()
}
