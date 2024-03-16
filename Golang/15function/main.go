package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is", result)

	fmt.Println("Result is : ", result)

	pro_res, my_message := proadder(2, 5, 6, 7)
	fmt.Println("pro result is ", pro_res)
	fmt.Println("pro Message is ", my_message)
}

// yahan ek important cheez hai ki jab h um kuch return krte hai to hme use fuction me mention krna pdeaga ki Kis type ki chhez return krna hai jaise yahan pr int kiya hai!
// like yahanse ek int return hog aisliye mention kiya
func adder(valOne int, valtwo int) int {
	return valOne + valtwo
}

// here we are trying to tell that all the values are going to be the type
func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi pro result"
}

func greeter() {
	fmt.Print("Namastey from golang")
}
