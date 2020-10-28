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
	myRouter.HandleFunc("/percentagedisk", controllers.DiskPercentage).Methods("GET")
	myRouter.HandleFunc("/downlaodkey", controllers.DownloadKey).Methods("POST")
	myRouter.HandleFunc("/userslogged", controllers.UsersLogged).Methods("GET")
	myRouter.HandleFunc("/unlink_user", controllers.Unlink_User).Methods("POST")
	myRouter.HandleFunc("/changepassword", controllers.ChangePassword).Methods("POST")
	myRouter.HandleFunc("/changeexpiration", controllers.ChangeExpiration).Methods("POST")

	return myRouter
}
