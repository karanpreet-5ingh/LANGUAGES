package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	language := make(map[string]string)

	language["js"] = "javascript"
	language["RB"] = "Ruby"
	language["py"] = "python"

	fmt.Println("List of all languages: ", language)

	//Removing any thing from the maps

	delete(language, "RB")

	fmt.Println("list of all languages: ", language)

	//looop are intresting in  golang

	for key, value := range language {

		fmt.Printf("For key %v, value is %v \n ", key, value)
	}
}
