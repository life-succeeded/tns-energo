package brigade

import (
	"fmt"
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Service struct {
	brigades Storage
}

func NewService(brigades Storage) *Service {
	return &Service{
		brigades: brigades,
	}
}

func (s *Service) Create(ctx libctx.Context, _ liblog.Logger, request CreateRequest) (Brigade, error) {
	now := time.Now()
	brigade := Brigade{
		FirstInspector:  request.FirstInspector,
		SecondInspector: request.SecondInspector,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	id, err := s.brigades.AddOne(ctx, brigade)
	if err != nil {
		return Brigade{}, fmt.Errorf("could not add brigade: %w", err)
	}

	brigade.Id = id

	return brigade, nil
}

func (s *Service) GetById(ctx libctx.Context, _ liblog.Logger, id string) (Brigade, error) {
	brigade, err := s.brigades.GetById(ctx, id)
	if err != nil {
		return Brigade{}, fmt.Errorf("could not find brigade: %w", err)
	}

	return brigade, nil
}

func (s *Service) GetAll(ctx libctx.Context, log liblog.Logger) ([]Brigade, error) {
	brigades, err := s.brigades.GetAll(ctx, log)
	if err != nil {
		return nil, fmt.Errorf("could not find brigades: %w", err)
	}

	return brigades, nil
}

func (s *Service) Update(ctx libctx.Context, _ liblog.Logger, id string, request UpdateRequest) error {
	b, err := s.brigades.GetById(ctx, id)
	if err != nil {
		return fmt.Errorf("could not find brigade: %w", err)
	}

	b.FirstInspector = request.FirstInspector
	b.SecondInspector = request.SecondInspector
	b.UpdatedAt = time.Now()

	err = s.brigades.Update(ctx, id, b)
	if err != nil {
		return fmt.Errorf("could not update brigade %s: %w", b.Id, err)
	}

	return nil
}
