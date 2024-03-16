package main

import "fmt"

func main() {
	fmt.Println("hello")

	days := []string{"Sunday", "Tuesday", "wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//Typical for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//iterating 2 item at the same time
	for index, day := range days {
		fmt.Printf("index is%v and value is %v \n", index, day)
	}

	// similar to While loop
	rough_value := 1

	for rough_value < 10 {

		// if rough_value == 5 {
		// 	rough_value++
		// 	break
		// }

		if rough_value == 5 {
			rough_value++
			continue
		}

		// this is goto statement which is going to use in the next step
		if rough_value == 2 {
			goto lco
		}

		fmt.Println("Value is: ", rough_value)
		rough_value++

	}

	// this is goto statement
lco:
	fmt.Println("jumping at lco")
}
