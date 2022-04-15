package reader

import (
	"errors"
	"parser/internal/filemanager"
	"strconv"
)

type datareader struct {
	fm filemanager.IFileManager
}

func InitDataReader(fm filemanager.IFileManager) IReader {
	return &datareader{
		fm: fm,
	}
}

func (r *datareader) formatData(data [][]string) []Data {
	var dataList []Data
	for i, line := range data {
		if i > 0 {
			var rec Data
			id, err := strconv.ParseUint(line[0], 10, 32)
			if err != nil {
				panic(err)
			}
			rec.id = uint32(id)

			rec.name = line[1]

			req, err := strconv.ParseBool(line[2])
			if err != nil {
				panic(err)
			}
			rec.required = req

			dataList = append(dataList, rec)
		} else {
			if line[0] != "id" && line[1] != "name" && line[2] != "required" {
				panic(errors.New("invalid file"))
			}
		}
	}

	return dataList
}

func (r *datareader) Read(path string, filename string) []Data {
	return r.formatData(r.fm.Read(path, filename))
}
