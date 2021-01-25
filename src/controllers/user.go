package controllers

import (
	//"fmt"

	"encoding/json"
	"net/http"
	"user"

	//"strings"
	"commands"
	//"os"
	//"log"
)

func AddUser(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.CreateUser(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

	rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(200)

	rw.Write(js)
}

func UsersLogged(rw http.ResponseWriter, req *http.Request) {

	js, _ := json.Marshal(commands.GetUsersLogged())
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func ChangePassword(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.ChangePassword(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

	rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(200)

	rw.Write(js)
}

func ChangeExpiration(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.ChangeExpiration(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

	rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(200)

	rw.Write(js)

}

func Unlink_User(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.Unlink_User(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

	rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.WriteHeader(200)

	rw.Write(js)

}

func DownloadKey(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.ReturnPathKey(person))
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
	//return path
}

func AddUserWithKey(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.CreateUserKey(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func DeleteUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var person user.User

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.DeleteUser(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func FixPermission(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	person := user.User{}

	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	js, err := json.Marshal(commands.FixPermission(person))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func DiskPercentage(rw http.ResponseWriter, req *http.Request) {
	js, err := json.Marshal(commands.DiskPercent())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func ListUsers(rw http.ResponseWriter, req *http.Request) {

	commands.ListDirectory("/data/users/bey")
	js, err := json.Marshal(commands.ListUsers())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}
