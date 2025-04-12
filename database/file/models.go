package file

type File struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}
