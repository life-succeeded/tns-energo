package task

import (
	"tns-energo/database/consumer"
	domain "tns-energo/service/task"
)

func MapToDb(t domain.Task) Task {
	return Task{
		Id:            t.Id,
		BrigadeId:     t.BrigadeId,
		Address:       t.Address,
		VisitDate:     t.VisitDate,
		Status:        int(t.Status),
		Consumer:      consumer.MapToDb(t.Consumer),
		AccountNumber: t.AccountNumber,
		Comment:       t.Comment,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

func MapToDomain(t Task) domain.Task {
	return domain.Task{
		Id:            t.Id,
		BrigadeId:     t.BrigadeId,
		Address:       t.Address,
		VisitDate:     t.VisitDate,
		Status:        domain.Status(t.Status),
		Consumer:      consumer.MapToDomain(t.Consumer),
		AccountNumber: t.AccountNumber,
		Comment:       t.Comment,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}

func MapSliceToDomain(tasks []Task) []domain.Task {
	result := make([]domain.Task, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, MapToDomain(t))
	}

	return result
}
