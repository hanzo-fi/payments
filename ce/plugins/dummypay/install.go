//go:build !it

package dummypay

import (
	"context"

	"github.com/hanzo-fi/payments/pkg/domain/plugins"
	"github.com/hanzo-fi/payments/pkg/domain/models"
)

func (p *Plugin) install(_ context.Context, _ models.InstallRequest) (models.InstallResponse, error) {
	return models.InstallResponse{}, plugins.ErrNotImplemented
}
