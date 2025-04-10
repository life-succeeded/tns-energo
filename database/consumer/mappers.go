package consumer

import domain "tns-energo/service/consumer"

func MapToDb(c domain.Consumer) Consumer {
	return Consumer{
		Surname:         c.Surname,
		Name:            c.Name,
		Patronymic:      c.Patronymic,
		LegalEntityName: c.LegalEntityName,
		PhoneNumber:     c.PhoneNumber,
	}
}

func MapToDomain(c Consumer) domain.Consumer {
	return domain.Consumer{
		Surname:         c.Surname,
		Name:            c.Name,
		Patronymic:      c.Patronymic,
		LegalEntityName: c.LegalEntityName,
		PhoneNumber:     c.PhoneNumber,
	}
}
