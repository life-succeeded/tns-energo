package minio

import "bytes"

type File struct {
	Name string
	Data *bytes.Buffer
}
