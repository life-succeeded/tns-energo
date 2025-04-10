package task

import (
	"fmt"
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Service struct {
	tasks Storage
}

func NewService(tasks Storage) *Service {
	return &Service{
		tasks: tasks,
	}
}

func (s *Service) AddOne(ctx libctx.Context, log liblog.Logger, request AddOneRequest) (Task, error) {
	var (
		now  = time.Now()
		task = Task{
			InspectorId:   request.InspectorId,
			Address:       request.Address,
			VisitDate:     request.VisitDate,
			Status:        Planned,
			Consumer:      request.Consumer,
			AccountNumber: request.AccountNumber,
			Comment:       request.Comment,
			CreatedAt:     now,
			UpdatedAt:     now,
		}
	)

	id, err := s.tasks.AddOne(ctx, task)
	if err != nil {
		return Task{}, fmt.Errorf("could not add task: %w", err)
	}

	task.Id = id

	return task, nil
}

func (s *Service) GetByInspectorId(ctx libctx.Context, log liblog.Logger, inspectorId int) ([]Task, error) {
	tasks, err := s.tasks.GetByInspectorId(ctx, log, inspectorId)
	if err != nil {
		return nil, fmt.Errorf("could not get tasks by inspector id: %w", err)
	}

	return tasks, nil
}

func (s *Service) UpdateStatus(ctx libctx.Context, log liblog.Logger, id string, request UpdateStatusRequest) error {
	err := s.tasks.UpdateStatus(ctx, id, request.NewStatus)
	if err != nil {
		return fmt.Errorf("could not update task status: %w", err)
	}

	return nil
}
