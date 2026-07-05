package services

import (
	"context"

	"github.com/hanzo-fi/payments/pkg/domain/models"
)

func (s *Service) ConversionsGet(ctx context.Context, id models.ConversionID) (*models.Conversion, error) {
	conversion, err := s.storage.ConversionsGet(ctx, id)
	if err != nil {
		return nil, newStorageError(err, "cannot get conversion")
	}

	return conversion, nil
}
