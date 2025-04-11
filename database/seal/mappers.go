package seal

import domain "tns-energo/service/seal"

func MapToDb(s domain.Seal) Seal {
	return Seal{
		Number: s.Number,
		Place:  s.Place,
	}
}

func MapSliceToDb(seals []domain.Seal) []Seal {
	result := make([]Seal, 0, len(seals))
	for _, s := range seals {
		result = append(result, MapToDb(s))
	}

	return result
}

func MapToDomain(s Seal) domain.Seal {
	return domain.Seal{
		Number: s.Number,
		Place:  s.Place,
	}
}

func MapSliceToDomain(seals []Seal) []domain.Seal {
	result := make([]domain.Seal, 0, len(seals))
	for _, s := range seals {
		result = append(result, MapToDomain(s))
	}

	return result
}
