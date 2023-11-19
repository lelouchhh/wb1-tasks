package main

import (
	"fmt"
	"sync"
)

func main() {
	var s []int
	s = []int{1, 2, 3, 4, 5}
	var sum int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, v := range s {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			sum += n * n
			mu.Unlock()

		}(v)
	}
	wg.Wait()
	fmt.Println(sum)
}
