package models

import "time"

type UserDetails struct {
	UserName   string
	Owner      string
	Key        bool
	Expiration time.Time
	Size       int64
}

type ListUserDetails struct {
	ListUsers []UserDetails
}
