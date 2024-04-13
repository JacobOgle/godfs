package internal

import (
	"log"
	"os"
)

// FileOps shared by files
// Read files
// Append only on writes
// Chunk files for distribution accross nodes --	generally 64MB chunks
type FileOps interface {
	Append()
	Read()
	Chunk()
}

type Chunk struct{}

// File is a basic representation of a file
type File struct {
	content []byte
	chunks  []Chunk
	size    int
}

func (f *File) ReadFileContentToStruct(path string) {
	fdata, err := os.ReadFile(path)
	if err != nil {
		// handle error...
		// TODO: write error handlers
		log.Fatal(err) // this is a temporary hack for compiler warning suppression
	}
	f.content = append(f.content, fdata...)
}

func (f File) Append() {}

func (f File) Read() {}

func (f *File) Chunk() {
	// 64MB may be too big for this small project.
	if len(f.content) > 0 {
		// TODO - Implement a chunking strategy
	}
}
