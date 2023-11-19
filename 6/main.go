package main

import (
	"context"
	"runtime"
	"sync"
)

func StopWithChannel(stopCh chan struct{}, f func()) {
	go func() {
		select {
		case <-stopCh:
			f()
		}
	}()
}

func StopWithContext(ctx context.Context, f func()) {
	go func() {
		select {
		case <-ctx.Done():
			f() // Вызов функции для остановки горутины
		}
	}()
}

func StopWithWaitGroup(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f() // Вызов функции для остановки горутины
	}()
}

func StopWithGoexit(f func()) {
	go func() {
		select {
		case <-stopCh:
			f()
			runtime.Goexit()
		}
	}()
}

func StopWithPanic(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				f()
			}
		}()
		f()
		panic("Stop the goroutine")
	}()
}
