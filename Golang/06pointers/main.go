package main

import "fmt"

func main() {

	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int
	// fmt.Println(" Value of Pointer is ", ptr)

	my_number := 23

	var ptr = &my_number // &  this symbol is use to give reference to that particular number for the pointer.

	fmt.Println("Value of actual pointer is ", ptr)  // will show the memory address where this ointer is stored.
	fmt.Println("Value of actual pointer is ", *ptr) //will show what value does this ptr hold

	*ptr = *ptr + 2

	fmt.Println("value of pointer is ", my_number)

}
