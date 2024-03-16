package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myurl = "http://localhost:8000/post"

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "karan")
	data.Add("lastname", "singh")
	data.Add("email", "karan@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
