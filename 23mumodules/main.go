package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// bring thirdparty library (mux) to run as a server 
func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang series by Saad</h1>"))
}
