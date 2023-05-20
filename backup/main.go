package main

import "fmt"

func main() {
	//Data types
	// int float string boolean complex
	//int uint, int8, int16, int32, int64, rune(int32) //
	//float float32, float64
	//complex complex64, complex128

	//operators - + * / %, >< >= <= == != && | ~ ++ --
	//if -else
	/*
				if x > y {
					fmt.Println("X is greater than y")
				} else if x == y {
					fmt.Println("X is equal to y")
				} else {
					fmt.Println("X is Less than y")
				}

			x := 45
			y := 45

				switch {
				case x >= y:
					fmt.Println("X is greater than y")
					fallthrough
				case x == y:
					fmt.Println("X is equal to y")
				default:
					fmt.Println("X is Less than y")
				}

		//loop for
		i := 0
		for i < 5 {
			fmt.Printf("Loop %v\n", i)
			i++
		}
	*/
	/*arrays
	num := [3]int{10, 45, 8}

		num[0] = 10
		num[1] = 45
		num[2] = 8

				fmt.Printf("Before change >> %v\n", num)
				change(num)
				fmt.Printf("After change >> %v\n", num)

				for i := 0; i < len(num); i++ {
					fmt.Printf("num[%v]=%v\n", i, num[i])
				}

			for _, v := range num {
				fmt.Printf("%v\n", v)
			}
	*/
	//slices
	//Dynamic arrays
	//var numbers []int;
	//numbers := []int{1, 2, 34}
	//var numbers = make([]int, 3, 4)
	//fmt.Printf("Size of slice %v and Silce = %v\n", len(numbers), numbers)
	//fmt.Println(numbers[0])
	//numbers = append(numbers, 20, 30, 40, 50)
	//fmt.Printf("Size of slice %v and Silce = %v\n", len(numbers), numbers)

	//fmt.Printf("Before change >>> %v\n", numbers)
	//change(numbers)
	//fmt.Printf("After change >>> %v\n", numbers)
	//fmt.Println(numbers[2:])
	//numbers = numbers[:3]
	//numbers = numbers[0:1]
	//fmt.Printf("len of slice = %v And capacity = %v\n", len(numbers), cap(numbers))
	//numbers = append(numbers, 10, 20, 30, 40, 50, 60, 70, 80, 90)
	//fmt.Printf("len of slice = %v And capacity = %v\n", len(numbers), cap(numbers))

	//maps
	//var state map[string]string
	//state := make(map[string]string)
	/*
		state := map[string]string{
			"Katsina": "Katsina",
		}

		state["Kaduna"] = "Kaduna"
		state["Oyo"] = "Ibadan"
		/*
			for _, v := range state {
				fmt.Printf("State Capitals: %v\n", v)
			}
			//fmt.Println(state)
	*/
	//fmt.Println(state["Zamfara"])
	//check for record existance

	/*
			if val, ok := state["Katsina"]; !ok {
				fmt.Println("Record Does Not Exists !")
			} else {
				fmt.Println("Value Exists!")
				fmt.Println(val)
			}

		delete(state, "Zamfara")

		delete(state, "Kaduna")
		change(state)
		for k, v := range state {
			fmt.Printf("State: %v and State Capitals: %v\n", k, v)
		}
		newMap := cloneMap(state)
		fmt.Println("Conetents of new map =============")
		newMap["FCT"] = "Abuja"
		for k, v := range newMap {
			fmt.Printf("State: %v and State Capitals: %v\n", k, v)
		}
	*/

	//Functions
	//Variadic

	//fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//fmt.Println(sum(1, 2, 3))

	//Multi Return
	/*
		ans, a, b, err := div(3, 0)
		if err != "" {
			fmt.Println(err)
			return
		}
		fmt.Printf("The result of dividing %v and %v is %v\n", a, b, ans)
	*/
	/*
		//Named Function
		var x int = 4
		var y int = 5
		j, msg := product(x, y)
		fmt.Printf("Product of %v and %v is %v and the message returned is: %v\n", x, y, j, msg)
	*/
	/*
		name := "Shamsu"
		func() {
			lname := "Khalil"
			func() {
				lname = "New Khalil"
				fmt.Printf("Hello Anonymous! %v, %v\n", name, lname)
			}()
		}()
	*/
	//Closures (Lambda functions)
	/*
		f := sayHello("shamsu")

		f()
		//fmt.Printf("%v\n)
	*/

	//recursion

	//sixFactorial := fact(3)
	//fmt.Println(sixFactorial)

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func sayHello(s string) func() {
	greet := "Hello Mr"
	return func() {
		greeting := fmt.Sprintf("%v, %v\n", greet, s)
		fmt.Println(greeting)
	}
}

func product(a, b int) (result int, msg string) {
	result = a * b
	msg = "Hello there!"
	return
}

func div(a float32, b float32) (float32, float32, float32, string) {
	if b == 0 {
		return 0, a, b, "Division by zero not allowed!"
	}
	return a / b, a, b, ""
}

func sum(a ...int) int {
	ans := 0
	for _, v := range a {
		ans += v
	}
	return ans
}

func change(m map[string]string) {
	m["Katsina"] = "Bakori"
	m["Oyo"] = "Obgbomosho"
}

func cloneMap(m map[string]string) map[string]string {
	retMap := make(map[string]string)
	for k, v := range m {
		retMap[k] = v
	}
	return retMap
}
