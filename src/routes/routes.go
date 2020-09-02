package routes

import (
	"controllers"
	"github.com/gorilla/mux"
)


func LoadRoutes() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/createuser", controllers.AddUser).Methods("POST")
	//myRouter.HandleFunc("/melao", otherDirectory).Methods("GET")
	//myRouter.HandleFunc("/me", test).Methods("POST")

	return myRouter
}