package main

import "fmt"

func main() {

	fmt.Println("welcome to array in golangs")

	var fruit_list [4]string

	fruit_list[0] = "Apple"
	fruit_list[1] = "pear"
	fruit_list[3] = "Banana"

	fmt.Println("Fruit list is: ", fruit_list)
	// as you can see we have initialize 4 vaue but we add only 3 value to the list
	//due to which you can see a big space between element 1 and 3

	fmt.Println("Fruit list is: ", len(fruit_list))
	// no matter how many element you use it is always going to return 4

	var veglist = [3]string{"potatos", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(veglist))
}
