package services

import (
	"context"

	"github.com/formancehq/go-libs/v5/pkg/storage/bun/paginate"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/storage"
)

func (s *Service) PaymentInitiationsList(ctx context.Context, query storage.ListPaymentInitiationsQuery) (*paginate.Cursor[models.PaymentInitiation], error) {
	pis, err := s.storage.PaymentInitiationsList(ctx, query)
	if err != nil {
		return nil, newStorageError(err, "cannot list payment initiations")
	}

	return pis, nil
}
