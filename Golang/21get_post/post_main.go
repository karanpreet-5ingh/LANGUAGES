package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const myurl = "http://localhost:8000/post"

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	// fmt.Pr intln("hi")
	// fmt.Println("ByteCount is: ", byteCount)
	// fmt.Println(responseString.String())
	// fmt.Println(string(content))

}

func PerformPostJsonRequest() {

	//fake json payload: simple url body...
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platorm":"learnCodeonline.in"
		}	
	`)

	// post request
	response, err := http.Post(myurl, "application/json", requestBody)

	//handle url
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
