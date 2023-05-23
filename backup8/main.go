package main

import (
	"fmt"
	"log"
)

func main() {
	n := 5
	slicech := makeNumSlice(n)
	//chanSlices := make(chan []int)
	chSquareArr := []<-chan int{}
	for i := 0; i < len(slicech); i++ {
		val := emitSquareValue(slicech)
		chSquareArr = append(chSquareArr, val)
	}
	sum := fanIn(chSquareArr)
	fmt.Printf("The sum squares of %v is %v\n", n, <-sum)
	/*
		squaresch := makeSquaresSlice(slicech)
		sum := 0
		for _, v := range <-squaresch {
			sum += v
		}
		fmt.Printf("The sum squares of %v is %v\n", n, sum)
	*/
}

func makeNumSlice(n int) <-chan []int {
	numch := make(chan []int)
	numSlice := []int{}
	go func() {
		for i := 0; i < n; i++ {
			numSlice = append(numSlice, i)
		}
		log.Printf("makeNumSlice:%v\n", numSlice)
		numch <- numSlice
	}()
	return numch
}

func emitSquareValue(numSlice <-chan []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range <-numSlice {
			ch <- v * v
		}
	}()
	return ch
}

func fanIn(input []<-chan int) <-chan int {
	sumch := make(chan int)
	sum := 0
	go func() {
		for _, ch := range input {
			sum += <-ch
		}
		sumch <- sum
	}()
	return sumch
}

func makeSquaresSlice(numSlice <-chan []int) <-chan []int {
	ch := make(chan []int)
	squareSlice := []int{}
	go func() {
		for _, v := range <-numSlice {
			squareSlice = append(squareSlice, v*v)
		}
		log.Printf("makeSquaresSlice:%v", squareSlice)
		ch <- squareSlice
	}()
	return ch
}
