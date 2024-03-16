package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to video on Slices")

	var fruitlist = []string{"apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T  \n ", fruitlist)

	fruitlist = append(fruitlist, "Mango", "banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	highscores := make([]int, 4) // type of dtaa  and size
	//make built-in function hai jo use krte h allocate and initializes krne k liye

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 253
	highscores[3] = 865
	// highscores[4] = 777

	// fmt.Println(highscores)

	highscores = append(highscores, 555, 666, 755)

	// fmt.Println(highscores)

	// fmt.Println(sort.IntsAreSorted(highscores)) //ye check krega ki aapki slice sorted hai ya nhi hai
	// sort.Ints(highscores)
	// fmt.Println(highscores)
	// fmt.Println(sort.IntsAreSorted(highscores))

	//how to remove a value from slice based on index

	var courses = []string{"reactjs", "javascirpt", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2

	//append ko hi use kiya h po p krne k lie, is cources wali list me se swift ko remove krna tha
	//to 0 se 1 tk select kr liya then 3 se last tk select kr liya.
	// means bich ka pop ho gya

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Print(courses)

}
