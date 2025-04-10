package file

import domain "tns-energo/service/file"

func MapToDb(f domain.File) File {
	return File{
		Name: f.Name,
		URL:  f.URL,
	}
}

func MapSliceToDb(files []domain.File) []File {
	result := make([]File, 0, len(files))
	for _, f := range files {
		result = append(result, MapToDb(f))
	}

	return result
}

func MapToDomain(f File) domain.File {
	return domain.File{
		Name: f.Name,
		URL:  f.URL,
	}
}

func MapSliceToDomain(files []File) []domain.File {
	result := make([]domain.File, 0, len(files))
	for _, f := range files {
		result = append(result, MapToDomain(f))
	}

	return result
}
