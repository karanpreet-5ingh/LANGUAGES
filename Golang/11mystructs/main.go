package main

import "fmt"

func main() {
	fmt.Println("Struts in golan")
	// no inheritance in golang: No super no parent

	karan := User{"karan", "karan .dev@gmail.com", true, 16}
	fmt.Println(karan)
	fmt.Printf("karan details are: %+v\n", karan)
	fmt.Printf("name is %v and email is %v", karan.Name, karan.Email)

}

type User struct {
	//U in User is capital so it is like a class
	//also see Name Email and Age also has first letter as capital

	Name   string
	Email  string
	Status bool
	Age    int
}
