package contract

import domain "tns-energo/service/contract"

func MapToDb(c domain.Contract) Contract {
	return Contract{
		Number: c.Number,
		Date:   c.Date,
	}
}

func MapToDomain(c Contract) domain.Contract {
	return domain.Contract{
		Number: c.Number,
		Date:   c.Date,
	}
}
