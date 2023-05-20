package cash

import "fmt"

type Cash struct{}

func (c Cash) Pay(a float64) {
	fmt.Printf("Paying %.2f Naira with Cash !\n", a)
}
