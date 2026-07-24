package services

import (
	"context"

	"github.com/hanzo-fi/go-libs/v5/pkg/storage/bun/paginate"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/storage"
)

func (s *Service) WorkflowsInstancesList(ctx context.Context, query storage.ListInstancesQuery) (*paginate.Cursor[models.Instance], error) {
	cursor, err := s.storage.InstancesList(ctx, query)
	return cursor, newStorageError(err, "cannot list instances")
}
