package routes

import (
	"controllers"

	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/createuser", controllers.AddUser).Methods("POST")
	myRouter.HandleFunc("/fixpermissionuser", controllers.FixPermission).Methods("POST")
	myRouter.HandleFunc("/deleteuser", controllers.DeleteUser).Methods("POST")
	myRouter.HandleFunc("/createuserwithkey", controllers.AddUserWithKey).Methods("POST")
	myRouter.HandleFunc("/listusers", controllers.ListUsers).Methods("GET")

	return myRouter
}
