package main


import "fmt"

func main() {

	var username string = "Karan"
	fmt.Println(username)
	fmt.Printf("Variable is of type:  %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type:  %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type:  %T \n", smallVal)

	var smallFloat float32 = 255.692726
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type:  %T \n", smallFloat)

	//default value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("vaiable is of type %T \n ", anotherVariable)

	// implicit type

	var website = "learncodeonline.in"
	fmt.Println(website)

	//website = 3   # hey i will not allow you to do this because you've already assing string to this variable

	// no var sytle

	numberOfUser := 30000.0 // see ese bhi variable define kr skte hai h bass yad rahe ki ese variable sirf 
	fmt.Println(numberOfUser)
	fmt.Printf("vaiable is of type %T \n ", numberOfUser)

}
