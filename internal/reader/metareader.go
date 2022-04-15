package reader

import (
	"errors"
	"parser/internal/filemanager"
	"time"
)

type metareader struct {
	fm filemanager.IFileManager
}

func InitMetaReader(fm filemanager.IFileManager) IMetaReader {
	return &metareader{
		fm: fm,
	}
}

func (r *metareader) formatData(data [][]string) (time.Time, error) {
	for _, line := range data {
		t, err := time.Parse(time.RFC3339Nano, line[0])
		if err != nil {
			panic(err)
		}

		return t, nil
	}

	return time.Time{}, errors.New("not found")
}

func (r *metareader) Read(path string, filename string) (time.Time, error) {
	return r.formatData(r.fm.Read(path, filename))
}
