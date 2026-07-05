package adyen

import "github.com/hanzo-fi/payments/pkg/domain/models"

var capabilities = []models.Capability{
	models.CAPABILITY_FETCH_ACCOUNTS,
	models.CAPABILITY_CREATE_WEBHOOKS,
	models.CAPABILITY_TRANSLATE_WEBHOOKS,
}
