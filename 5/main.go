package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(1000)
		}
	}()
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	time.Sleep(time.Minute)
}
