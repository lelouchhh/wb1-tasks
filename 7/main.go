package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[string]int)
	var mu sync.Mutex

	var wg sync.WaitGroup

	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			value := id * id

			mu.Lock()
			data[key] = value
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	mu.Lock()
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
	mu.Unlock()
}
