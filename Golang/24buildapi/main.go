package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` // yahana  pr since hame Author niche declare kr diya h or wo type class ka hai to use hmm use kr skte hain
	// we use *Author not &Author or simple Author because "Author" might make changes in a copy of struct Author  and &Author is not used because Author is not yet declared

}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website "`
}

// fake db
var courses []Course

// simple method call kiya hai jisme check kr rahe hai ki CourseID and CourseName khali na aa jae.
// middleware, helper  - file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

//controllers  - file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Get all the courses \n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //need to encose json
}
