package registry

import (
	"bytes"
	"fmt"
	"strings"
	"tns-energo/config"
	"tns-energo/database/registry"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"github.com/xuri/excelize/v2"
)

type Service interface {
	Parse(ctx libctx.Context, log liblog.Logger, fileBytes []byte) error
	GetItemByAccountNumber(ctx libctx.Context, log liblog.Logger, accountNumber string) (Item, error)
}

type Impl struct {
	settings config.Settings
	registry registry.Repository
}

func NewService(settings config.Settings, registry registry.Repository) *Impl {
	return &Impl{
		settings: settings,
		registry: registry,
	}
}

func (s *Impl) Parse(ctx libctx.Context, log liblog.Logger, fileBytes []byte) error {
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

	items := make([]registry.Item, 0, len(rows)-1)
	for _, row := range rows[1:] {
		if len(row) != 6 {
			log.Errorf("invalid row length: %d", len(row))
			continue
		}

		haveAutomaton := false
		if strings.EqualFold(row[5], "+") {
			haveAutomaton = true
		}

		items = append(items, registry.Item{
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

func (s *Impl) GetItemByAccountNumber(ctx libctx.Context, log liblog.Logger, accountNumber string) (Item, error) {
	item, err := s.registry.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		return Item{}, fmt.Errorf("could not get item by account number: %w", err)
	}

	return Item{
		Id:            item.Id,
		AccountNumber: item.AccountNumber,
		Surname:       item.Surname,
		Name:          item.Name,
		Patronymic:    item.Patronymic,
		Object:        item.Object,
		HaveAutomaton: item.HaveAutomaton,
	}, nil
}
