package controllers

import (
	//"fmt"
	"encoding/json"
	"net/http"
	"user"

	//"helper"
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
	if commands.CreateUser(person) {

	}
}
