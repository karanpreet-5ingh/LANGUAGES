package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length is:", response.ContentLength)

	// so bytes aati response me to use string me convert krna pdta hai use krne k liye mai ye krna padega

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body) //here this contrnt is in the bytes format

	byteCount, _ := responseString.Write(content) //see yahan hnme strings library ka use kr k ek buuilder create kiya and us builder ko hmm

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String(), "fajskfl;j")
	fmt.Println(string(content))

}
