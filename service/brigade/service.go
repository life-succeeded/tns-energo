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
	brigade := Brigade{
		FirstInspector:  request.FirstInspector,
		SecondInspector: request.SecondInspector,
		CreatedAt:       time.Now(),
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
