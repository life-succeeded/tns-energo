package registry

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/device"

	"github.com/xuri/excelize/v2"
)

type Service struct {
	registry Storage
}

func NewService(registry Storage) *Service {
	return &Service{
		registry: registry,
	}
}

func (s *Service) Parse(ctx libctx.Context, log liblog.Logger, payload []byte) error {
	file, err := excelize.OpenReader(bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("could not open excel file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("failed to close excel file: %v", err)
		}
	}()

	sheet := file.GetSheetName(0)
	rows, err := file.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("could not get rows: %w", err)
	}

	items := make([]Item, 0, len(rows)-9)
	now := time.Now()
	for _, row := range rows[9:] {
		addressSlice := make([]string, 0, 4)

		if len(row[3]) != 0 {
			addressSlice = append(addressSlice, row[3])
		}

		if len(row[4]) != 0 {
			addressSlice = append(addressSlice, row[4])
		}

		if len(row[5]) != 0 {
			addressSlice = append(addressSlice, fmt.Sprintf("д %v", row[5]))
		}

		if len(row[6]) != 0 {
			addressSlice = append(addressSlice, row[6])
		}

		deploymentDate, err := time.Parse("01-02-06", row[15])
		if err != nil {
			log.Errorf("could not parse deployment date: %v", err)
			continue
		}

		items = append(items, Item{
			Address: strings.Join(addressSlice, ", "),
			OldDevice: device.Device{
				Type:   row[10],
				Number: row[11],
			},
			NewDevice: device.Device{
				Type:           row[13],
				Number:         row[14],
				DeploymentDate: deploymentDate,
			},
			CreatedAt: now,
			UpdatedAt: now,
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
