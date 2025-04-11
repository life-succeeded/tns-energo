package inspector

import domain "tns-energo/service/inspector"

func MapToDb(i domain.Inspector) Inspector {
	return Inspector{
		Surname:    i.Surname,
		Name:       i.Name,
		Patronymic: i.Patronymic,
	}
}

func MapToDomain(i Inspector) domain.Inspector {
	return domain.Inspector{
		Surname:    i.Surname,
		Name:       i.Name,
		Patronymic: i.Patronymic,
	}
}
