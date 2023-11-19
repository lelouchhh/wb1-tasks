package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	dataChannel := make(chan int)

	var wg sync.WaitGroup
	stopSignal := make(chan struct{})

	N := 100
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, stopSignal)
	}

	go func() {
		defer close(dataChannel)
		for i := 1; i <= 10; i++ {
			dataChannel <- i
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("stopped signal had appeared")
		close(stopSignal) // Закрыть канал для уведомления воркеров о завершении
	}()
	wg.Wait()
}

func worker(id int, ch <-chan int, wg *sync.WaitGroup, stop <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: Finished\n", id)
				return
			}
			fmt.Printf("Worker %d: Received %d\n", id, data)
		case <-stop:
			fmt.Printf("Worker %d: Stopped\n", id)
			return
		}
	}
}
