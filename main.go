package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"routes"
)



func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func otherDirectory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Other!")
	fmt.Println("Endpoint Hit: other")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello World")
}
type test_struct struct {
    Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t test_struct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
	log.Println(t.Test)
	fmt.Println(t.Test)
}

func handleRequests() {
	fmt.Println("Go Tutorial")
	log.Fatal(http.ListenAndServe(":8081", routes.LoadRoutes()))
}

func main() {
	handleRequests()
	//helper.CheckUserExists("bruno")
}
