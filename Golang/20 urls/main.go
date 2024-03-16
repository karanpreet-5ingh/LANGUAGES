package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {

	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	query_parameters := result.Query() // extract all information form the given url

	fmt.Printf("The type of query params are: %T \n", query_parameters)

	fmt.Println(query_parameters["coursename"])

	for _, val := range query_parameters {
		fmt.Println("param is: ", val)
	}

	// maybe you have all the information in cunks and part and you might want it to convert in a full url then this is how you go!!
	//&url is important to pass reference of the name not copy...
	// &url.URL is the predefined thing that you need to keep in mind and the parameters should be the very use which are provided by language itself...

	part_of_url := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Karan ",
	}

	another_url := part_of_url.String()
	fmt.Println(another_url)

}
