package filemanager

import "time"

type IFileManager interface {
	Read(path string, filename string) [][]string
	GetUpdateDate(path string, filename string) time.Time
	Write(path string, filename string, data string)
}
