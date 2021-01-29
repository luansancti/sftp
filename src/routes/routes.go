package routes

import (
	"net/http"
	"controllers"
	"os"
	"github.com/gorilla/handlers"
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

func SetRequest() http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(LoadRoutes())

}
