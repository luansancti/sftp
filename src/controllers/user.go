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

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
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

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func DiskPercentage(rw http.ResponseWriter, req *http.Request) {
	// js, err := json.Marshal(commands.DiskPercent())
	// if err != nil {
	// 	http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// rw.Header().Set("Content-Type", "application/json")
	// rw.Write(js)
	commands.DiskPercent()

}

func ListUsers(rw http.ResponseWriter, req *http.Request) {

	js, err := json.Marshal(commands.ListUsers())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}
