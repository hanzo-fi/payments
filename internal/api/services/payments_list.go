package services

import (
	"context"

	"github.com/formancehq/go-libs/v5/pkg/storage/bun/paginate"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/storage"
)

func (s *Service) PaymentsList(ctx context.Context, query storage.ListPaymentsQuery) (*paginate.Cursor[models.Payment], error) {
	ps, err := s.storage.PaymentsList(ctx, query)
	if err != nil {
		return nil, newStorageError(err, "cannot list payments")
	}

	return ps, nil
}
