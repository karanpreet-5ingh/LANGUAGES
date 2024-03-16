package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	mydefer()
}


//world, one, Two
//0, 1, 2, 3, 4
//hello, 43210, two, one, world

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}
