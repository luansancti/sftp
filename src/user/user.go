package user

import (
	"helper"
	"path"
)

type User struct {
	User           string
	Expiration     int
	Password       string
	PathUser       string
	PathUserPublic string
	PathUserUp     string
	PathPublic     string
}

func NewUser(username string, password string, expiration int) User {
	user := User{}
	user.User = username
	user.Password = password
	user.Expiration = expiration
	user.PathUser = path.Join(helper.GetConfigPaths().UsersPath, username)
	user.PathUserUp = path.Join(helper.GetConfigPaths().UsersPath, username, "upload")
	user.PathUserPublic = path.Join(helper.GetConfigPaths().UsersPath, username, "public")
	user.PathPublic = helper.GetConfigPaths().PublicPath
	return user
}
