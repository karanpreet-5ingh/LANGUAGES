package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "this need to go in a file -  Bhai kya baat hai ye bhi sikh gya "

	file, err := os.Create("./gofile.txt")

	if err != nil {
		panic(err)
	}

	// so ye wali line se hum write krwaenge and write k sath sath ye length return krega ya error
	//kaise pata ki length hi return krega ? bhai documentation padh le.....
	length, err := io.WriteString(file, content)

	check_nill_error(err)

	fmt.Println("length is: ", length)
	defer file.Close()

	readfile("./gofile.txt")

}

func readfile(filename string) {
	databyte, err := os.ReadFile(filename)
	check_nill_error(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func check_nill_error(err error) {

	if err != nil {
		panic(err)
	}
}
