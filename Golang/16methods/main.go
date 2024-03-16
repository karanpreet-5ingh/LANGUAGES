// kuch nhi h function jo structs me jaenge unhe method bolte hai!!

package main

import "fmt"

func main() {
	fmt.Println("Struts in golan")
	// no inheritance in golang: No super no parent

	karan := User{"karan", "karan.dev@gmail.com", true, 16}
	fmt.Println(karan)
	fmt.Printf("karan details are: %+v\n", karan)
	fmt.Printf("name is %v and email is %v", karan.Name, karan.Email)

	karan.GetStatus()
	karan.Change_email() // see mail is change when this is execute but it is not change at the real place to check this we print the next line
	fmt.Printf("name is %v and email is %v", karan.Name, karan.Email)

}

// structure define kiya h
type User struct {
	//U in User is capital so it is like a class
	//also see Name Email and Age also has first letter as capital
	Name   string
	Email  string
	Status bool
	Age    int
}

// ye methods h
func (u User) GetStatus() {
	fmt.Println("Is user active ", u.Status) //status ek boolen value return krega jo true ya false batega
}

func (u User) Change_email() {
	u.Email = "changed@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)

}


// see when i use the above method Change_email at that time actuall email which i define in struct is not change iat only pass along it's copy
//now next thing is that if you want to change it hae you should pass the pointer reference


