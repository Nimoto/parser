package reader

import "time"

type IReader interface {
	Read(path string, filename string) []Data
}

type IMetaReader interface {
	Read(path string, filename string) (time.Time, error)
}
