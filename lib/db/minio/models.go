package minio

import "bytes"

type File struct {
	Name string
	Data *bytes.Buffer
}

type Policy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Effect    string    `json:"Effect"`
	Principal Principal `json:"Principal"`
	Action    []string  `json:"Action"`
	Resource  []string  `json:"Resource"`
}

type Principal struct {
	AWS []string `json:"AWS"`
}
