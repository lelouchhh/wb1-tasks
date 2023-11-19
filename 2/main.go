package main

import (
	"fmt"
	"sync"
)

func main() {
	var s []int
	s = []int{1, 2, 3, 4}
	ch := make(chan int, len(s))
	wg := sync.WaitGroup{}
	for _, v := range s {
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- v * v
		}()
	}
	wg.Wait()
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
