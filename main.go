package main

import (
	"fmt"
	"log"
	"net/http"
	"routes"
)

func handleRequests() {
	fmt.Println("Go Tutorial")
	log.Fatal(http.ListenAndServe(":8081", routes.SetRequest()))
}

func main() {
	handleRequests()
}
