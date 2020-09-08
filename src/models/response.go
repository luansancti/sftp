package models

type CreateUser struct {
	Success bool
	Message string
}

type FixPermission struct {
	CreateUser
}

type ListUser struct {
	Success bool
	Message string
	Data    ListUserDetails
}

type DeleteUser struct {
	CreateUser
}

type DefaultResponse struct {
	Success bool
	Message string
}

func ResponseDefault(message string, success bool) DefaultResponse {
	response := DefaultResponse{}
	response.Message = message
	response.Success = success
	return response
}

func ResponseCreate(message string, success bool) CreateUser {
	responseCreate := CreateUser{}
	responseCreate.Message = message
	responseCreate.Success = success
	return responseCreate
}

func ResponsePermission(message string, success bool) FixPermission {
	reponsePermission := FixPermission{}
	reponsePermission.Message = message
	reponsePermission.Success = success
	return reponsePermission
}
