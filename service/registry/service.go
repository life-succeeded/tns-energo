package registry

import (
	"bytes"
	"fmt"
	"strings"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"github.com/xuri/excelize/v2"
)

type Service struct {
	settings config.Settings
	registry RegistryStorage
}

func NewService(settings config.Settings, registry RegistryStorage) *Service {
	return &Service{
		settings: settings,
		registry: registry,
	}
}

func (s *Service) Parse(ctx libctx.Context, log liblog.Logger, fileBytes []byte) error {
	file, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("failed to close file: %v", err)
		}
	}()

	sheet := file.GetSheetName(0)
	rows, err := file.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("could not get rows: %w", err)
	}

	items := make([]Item, 0, len(rows)-1)
	for _, row := range rows[1:] {
		if len(row) != 6 {
			log.Errorf("invalid row length: %d", len(row))
			continue
		}

		haveAutomaton := false
		if strings.EqualFold(row[5], "+") {
			haveAutomaton = true
		}

		items = append(items, Item{
			AccountNumber: row[0],
			Surname:       row[1],
			Name:          row[2],
			Patronymic:    row[3],
			Object:        row[4],
			HaveAutomaton: haveAutomaton,
		})
	}

	err = s.registry.AddMany(ctx, items)
	if err != nil {
		return fmt.Errorf("could not add items: %w", err)
	}

	return nil
}

func (s *Service) GetItemByAccountNumber(ctx libctx.Context, log liblog.Logger, accountNumber string) (Item, error) {
	item, err := s.registry.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		return Item{}, fmt.Errorf("could not get item by account number: %w", err)
	}

	return item, nil
}

func (s *Service) GetItemsByAccountNumberRegular(ctx libctx.Context, log liblog.Logger, accountNumber string) ([]Item, error) {
	items, err := s.registry.GetByAccountNumberRegular(ctx, log, accountNumber)
	if err != nil {
		return nil, fmt.Errorf("could not get items by account number: %w", err)
	}

	return items, nil
}
