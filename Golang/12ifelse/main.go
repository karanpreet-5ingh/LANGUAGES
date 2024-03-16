package main

import "fmt"

func main() {
	fmt.Println("If elese in golang")

	logincount := 23

	var result string

	if logincount < 10 {

		result = "Regular user"
	} else if logincount > 10 {
		result = "wactch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	// case 2 ye bhi valid hai
	if 9%2 == 10 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// case 2 ye bhi valid hai
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is more than 10")
	}

	//if err != nil{

	//}

}
