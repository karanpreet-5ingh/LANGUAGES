package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") //when the get request is comming up this guy servehome is resposible for what to show there at route "/"

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("hey there mod uses")

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome my friend welcome!!</h1>"))
}
