package models

type UserDetails struct {
	UserName   string
	Owner      string
	Key        bool
	Expiration string
	Size       int64
}

type ListUserDetails struct {
	ListUsers []UserDetails
}
