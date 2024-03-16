package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Print("Switch and case in golang\n")
	dice_number := rand.Intn(6) + 1
	fmt.Println("values of dice is ", dice_number)

	switch dice_number {
	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("you can move 2 spot")

	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough

	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough

	case 5:
		fmt.Println("you can move 5 spot")

	case 6:
		fmt.Println("you can move 7 spot")

	default:
		fmt.Println("what was that!!!")

	}

}
