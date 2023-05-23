package main

import (
	"sync"
	"time"
)

func main() {
	var lock sync.RWMutex
	var x int = 0
	for i := 0; i < 5; i++ {
		go changeX(&x, &lock)
		go printX(&x, &lock)
	}
	time.Sleep(100 * time.Millisecond)
}

func printX(a *int, l *sync.RWMutex) {
	l.RLock()
	println(*a)
	l.RUnlock()
}

func changeX(a *int, l *sync.RWMutex) {
	l.Lock()
	*a++
	l.Unlock()
}
