package services

import (
	"context"

	"github.com/formancehq/go-libs/v5/pkg/storage/bun/paginate"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/storage"
)

func (s *Service) SchedulesList(ctx context.Context, query storage.ListSchedulesQuery) (*paginate.Cursor[models.Schedule], error) {
	cursor, err := s.storage.SchedulesList(ctx, query)
	return cursor, newStorageError(err, "cannot list schedules")
}
