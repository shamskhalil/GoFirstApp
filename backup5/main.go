package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	cpus := runtime.NumCPU()
	fmt.Printf("No of CPUS >> %v\n", cpus)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	now := time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "User Logged in successfully!"
			time.Sleep(20 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nLogin Duration:%v\n", i, response, duration)
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "User Logged in successfully!"
			time.Sleep(20 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nLogin Duration:%v\n", i, response, duration)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "User Logged in successfully!"
			time.Sleep(20 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nLogin Duration:%v\n", i, response, duration)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "User Logged in successfully!"
			time.Sleep(20 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nLogin Duration:%v\n", i, response, duration)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "User Info: Khalil Shams"
			time.Sleep(10 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nUserInfo Duration:%v\n", i, response, duration)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "I like pepsi, coke and fanta"
			time.Sleep(100 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nUserLikes Duration:%v\n", i, response, duration)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			now := time.Now()
			response := "I like pepsi, coke and fanta"
			time.Sleep(100 * time.Millisecond)
			duration := fmt.Sprintf("%v\n", time.Since(now))
			fmt.Printf("Loop:%v\nResponse: %v\nUserLikes Duration:%v\n", i, response, duration)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Total time taken: %v\n", time.Since(now))

}
