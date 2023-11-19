package main

import (
	"fmt"
	"sync"
)

type counter struct {
	val int
	mu  sync.Mutex
}

func (c *counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func (c *counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func main() {
	counter := counter{}
	numGoroutines := 10
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait() // Ожидание завершения всех горутин

	finalCount := counter.Get()
	fmt.Printf("Итоговое значение счетчика: %d\n", finalCount)
}
