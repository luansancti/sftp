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
	Data    []UserDetails
}

type DirectoryPerc struct {
	Success bool
	Message string
	Data    []PercentageUsed
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

func ResponseListUsers(message string, success bool, data []UserDetails) ListUser {
	ResponseListUsers := ListUser{}
	ResponseListUsers.Message = message
	ResponseListUsers.Success = success
	ResponseListUsers.Data = data
	return ResponseListUsers
}

func ResponseDirectoryPerc(message string, success bool, data []PercentageUsed) DirectoryPerc {
	ResponseListUsers := DirectoryPerc{}
	ResponseListUsers.Message = message
	ResponseListUsers.Success = success
	ResponseListUsers.Data = data
	return ResponseListUsers
}
