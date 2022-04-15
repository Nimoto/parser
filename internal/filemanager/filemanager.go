package filemanager

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"
)

type filemanager struct {
}

func InitFileManager() IFileManager {
	return &filemanager{}
}

func (fm *filemanager) Read(path string, filename string) [][]string {
	f := fm.openFile(path, filename)
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return data
}

func (fm *filemanager) GetUpdateDate(path string, filename string) time.Time {
	fi, err := os.Stat(fm.getFullPath(path, filename))
	if err != nil {
		panic(err)
	}

	return fi.ModTime()
}

func (fm *filemanager) Write(path string, filename string, data string) {
	println(path)
	f := fm.openFile(path, filename)
	defer f.Close()

	b := []byte(data)

	err := os.WriteFile(fm.getFullPath(path, filename), b, 0644)
	if err != nil {
		panic(err)
	}
}

func (fm *filemanager) openFile(path string, filename string) *os.File {
	f, err := os.Open(fm.getFullPath(path, filename))
	if err != nil {
		panic(err)
	}

	return f
}

func (fm *filemanager) getFullPath(path string, filename string) string {
	d, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return d + "/" + filename
}
