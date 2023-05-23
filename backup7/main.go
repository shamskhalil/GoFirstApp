package main

import (
	"fmt"
	"math/rand"
	"time"
)

func talk(name string) <-chan string {
	i := 0
	ch := make(chan string)
	go func(i int) {
		for {
			i++
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch <- fmt.Sprintf("%v: %v\n", name, i)
		}
	}(i)
	return ch
}

func main() {
	rand.Seed(time.Now().UnixMilli())
	shams := talk("Shams")
	aminu := talk("Aminu")
	//go talk("Shams", ch)
	timeout := time.After(2000 * time.Millisecond)
	for {
		select {
		case d := <-shams:
			fmt.Print(d)
		case e := <-aminu:
			fmt.Print(e)
		case <-timeout:
			fmt.Printf("Time up, please drop your pens!\n")
			return
		}

	}

}
