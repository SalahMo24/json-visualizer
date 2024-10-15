package jsonmodule

import "io/fs"

type Input = map[string]interface{}

type File struct {
	Name  string
	Stats fs.FileInfo
}

type FileHandler interface {
	Reader(filepath string) (map[string]interface{}, error)
	Writer(input Input) error
	GetStats() fs.FileInfo
	Append(newInputs Input) (Input, error)
}
