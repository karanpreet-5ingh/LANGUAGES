package main

import (
	"encoding/json"
	"fmt"
)

// declaring a structure
type course struct {
	Name     string `json:"coursename"` // golang knows thta you are going to use the se structurs in jdon alot so they have provided you the alias kind of thing
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this is how we hide passwords in our api
	Tags     []string `json:"tags,omitempty"` //if eny field is nil just omit that
}

func main() {
	fmt.Println("welcome to json video")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	//slice of typpe courses
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 299, "learncodeonline.in", "abc123", []string{"full stack", "js"}},
		{"Angular Bootcamp", 299, "learncodeonline.in", "abc123", nil},
	}

	//package this data as json data
	//marshal is the way to implement the json
	// interface is a word being borrowed and its the alternate version of struct
	// finaljson, err := json.Marshal(lcoCourses)

	finaljson, err := json.MarshalIndent(lcoCourses, "", "\t") // /t help to increase the readabiliity of code
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finaljson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
{
	"coursename": "Mern Bootcamp",
	"Price": 299,
	"website": "learncodeonline.in",
	"tags": ["full stack","js"]
}
`)

	var lcoCourses course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses) //json banane k time hamne marshal kiya tha abb unmarshal kr rahe hain
		// so when you are not sure ki lcoCourses ki exact value pass krega ya copy to shi krne k liye aap referense pass krwate ho

		fmt.Printf("%#v\n", lcoCourses) // interfaces ko pprint krne k liye hame # sign use krna padta hai
	} else {
		fmt.Println("json is not valid")
	}

	// some cases where you want to add data to key value
	// when ever you are creating a map aka key, value pair you cannot sure that the value is always string it could be int and other types as well.

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData) // interfaces ko pprint krne k liye hame # sign use krna padta hai

	for k, v := range myOnlineData {
		fmt.Printf("key is  %v and value is %v and Type is: %T \n", k, v, v)
	}
}
