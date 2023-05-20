package main

import "fmt"

func main() {
	//Defer
	/*
		fmt.Println("file opened")
		defer fmt.Println("file closed")
		fmt.Println("writting to file ...")
		fmt.Println("adding more data to file and writting ...")
	*/

	//Panics
	/*
		water()
		seal()
		store()
	*/

	//pointers
	//& (Address)mpersand
	// * (Stored)tar

	x := 23
	y := 10
	var z *int = &x
	//var ax *int = &x
	//ax = &x
	//fmt.Printf("x = %v at address = %v\n", x, ax)
	//fmt.Printf("Address = %v and value Stroed at Address = %v\n", ax, *ax)
	fmt.Printf("Address of z = %p and content of z Address is %p\n", z, *z)
	fmt.Printf("Before: x = %v and y = %v\n", x, y)
	fmt.Printf("%v minus %p is %p\n", x, y, subtract(&x, &y))
	fmt.Printf("After: x = %v and y = %v\n", x, y)
	fmt.Println("And z i now %v", *z)

}

func subtract(x, y *int) int {
	*x = *x - *y
	return *x
}

func catchPanics(scope string) {
	if k := recover(); k != nil {
		switch k {
		case "In water purification":
			seal()
			store()
		case "In sealing machine":
			store()
		case "In water product storage":
			store()
		}
		fmt.Println("Panic in: ", scope)
		fmt.Printf("PANIC REASON :%v\n", k)
		fmt.Println("recovered !")
	}
}

func water() {
	defer catchPanics("In water purification")
	fmt.Println("Adding water ...")
}

func seal() {
	defer catchPanics("In sealing machine")
	panic("oops bad sealing machine stopped")
	fmt.Println("sealing packaged water ...")
}

func store() {
	defer catchPanics("In water product storage")
	fmt.Println("Storing water\n")
}
