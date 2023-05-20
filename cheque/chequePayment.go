package cheque

import "fmt"

type Cheque struct{}

func (c Cheque) Pay(a float64) {
	fmt.Printf("Paying %.2f Naira with Cheque !\n", a)
}
