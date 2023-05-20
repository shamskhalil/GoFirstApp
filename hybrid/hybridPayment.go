package hybrid

import (
	"fmt"

	"github.com/shamskhalil/GoFirstApp/iface"
)

type Hybrid struct {
}

type HybridPayment interface {
	iface.Payment
	payHalf(float64)
}

func (h Hybrid) Pay(a float64) {
	fmt.Printf("Paying half, amount: %.2f Naira with Card !\n", a)
}

func (h Hybrid) PayHalf(a float64) {
	fmt.Printf("Paying half, amount: %.2f Naira with Cash !\n", a)
}
