package services

import (
	"context"

	"github.com/formancehq/go-libs/v5/pkg/storage/bun/paginate"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/storage"
)

func (s *Service) BankAccountsList(ctx context.Context, query storage.ListBankAccountsQuery) (*paginate.Cursor[models.BankAccount], error) {
	bas, err := s.storage.BankAccountsList(ctx, query)
	if err != nil {
		return nil, newStorageError(err, "cannot list bank accounts")
	}

	return bas, nil
}
