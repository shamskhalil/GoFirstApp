package main

import (
	"fmt"
	"sync"
)

/*
	func addition(a, b int) int {
		return a + b
	}
*/
/*
func additionWeb(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(args[0].Int() + args[1].Int())
}
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Line 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Line 2")
		wg.Done()
	}()

	fmt.Println("Line 3")
	fmt.Println("End of Program")

	wg.Wait()
	/*
		fmt.Println("web assembly initialized ....")
		js.Global().Set("webAdd", js.FuncOf(additionWeb))

		time.Sleep(1 * time.Hour)
		/*

			myCard := card.NewCard(100)

			myCard.GetInfo()
			myCard.Pay(23.50)
			myCard.GetInfo()
			newCard := myCard.ReplaceCard(20.50)
			newCard.GetInfo()
			newCard.Pay(10.50)
			newCard.GetInfo()

			/*
				var cardPayment Payment = Card{balance: 100}
				cardPayment.pay(23.5)

				var cashPayment Payment = Cash{}
				cashPayment.pay(43.5)

				var chequePayment Payment = Cheque{}
				chequePayment.pay(15.5)

				var hybridPayment HybridPayment = Hybrid{}
				hybridPayment.pay(10.3)
				hybridPayment.payHalf(10.3)
	*/

}
