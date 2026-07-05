package plugins

import "github.com/hanzo-fi/payments/pkg/domain/models"

type Plugin struct {
	models.PSPPlugin
	models.OpenBankingPlugin
}
