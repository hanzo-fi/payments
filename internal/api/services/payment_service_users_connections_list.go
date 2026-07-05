package services

import (
	"context"

	"github.com/formancehq/go-libs/v5/pkg/storage/bun/paginate"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/storage"
	"github.com/google/uuid"
)

func (s *Service) PaymentServiceUsersConnectionsList(ctx context.Context, psuID uuid.UUID, connectorID *models.ConnectorID, query storage.ListOpenBankingConnectionsQuery) (*paginate.Cursor[models.OpenBankingConnection], error) {
	ps, err := s.storage.OpenBankingConnectionsList(ctx, psuID, connectorID, query)
	if err != nil {
		return nil, newStorageError(err, "cannot list open banking connections")
	}

	return ps, nil
}
