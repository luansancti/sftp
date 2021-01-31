package models

import "time"

type DirectoryInfo struct {
	Name        string
	Size        int64
	IsDirectory bool
	ModTime     time.Time
}
