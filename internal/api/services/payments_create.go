package services

import (
	"context"

	"github.com/hanzo-fi/payments/pkg/domain/models"
)

func (s *Service) PaymentsCreate(ctx context.Context, payment models.Payment) error {
	return handleEngineErrors(s.engine.CreateFormancePayment(ctx, payment))
}
