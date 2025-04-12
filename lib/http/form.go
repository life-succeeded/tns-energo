package http

import "mime/multipart"

type FormDataField struct {
	Name  string
	Value string
}

type FormDataFile struct {
	FieldName string
	FileName  string
	Payload   multipart.File
}
