//go:build !ee

package registry

import (
	adyen "github.com/hanzo-fi/payments/ce/plugins/adyen"
	atlar "github.com/hanzo-fi/payments/ce/plugins/atlar"
	bankingcircle "github.com/hanzo-fi/payments/ce/plugins/bankingcircle"
	column "github.com/hanzo-fi/payments/ce/plugins/column"
	currencycloud "github.com/hanzo-fi/payments/ce/plugins/currencycloud"
	dummypay "github.com/hanzo-fi/payments/ce/plugins/dummypay"
	generic "github.com/hanzo-fi/payments/ce/plugins/generic"
	increase "github.com/hanzo-fi/payments/ce/plugins/increase"
	mangopay "github.com/hanzo-fi/payments/ce/plugins/mangopay"
	modulr "github.com/hanzo-fi/payments/ce/plugins/modulr"
	moneycorp "github.com/hanzo-fi/payments/ce/plugins/moneycorp"
	plaid "github.com/hanzo-fi/payments/ce/plugins/plaid"
	powens "github.com/hanzo-fi/payments/ce/plugins/powens"
	qonto "github.com/hanzo-fi/payments/ce/plugins/qonto"
	stripe "github.com/hanzo-fi/payments/ce/plugins/stripe"
	tink "github.com/hanzo-fi/payments/ce/plugins/tink"
	wise "github.com/hanzo-fi/payments/ce/plugins/wise"
	pkgplugins "github.com/hanzo-fi/payments/pkg/domain/plugins"
)

func init() {
	load(map[string]pkgplugins.Registration{
		adyen.ProviderName:         adyen.Registration,
		atlar.ProviderName:         atlar.Registration,
		bankingcircle.ProviderName: bankingcircle.Registration,
		column.ProviderName:        column.Registration,
		currencycloud.ProviderName: currencycloud.Registration,
		DummyPSPName:               dummypay.Registration,
		generic.ProviderName:       generic.Registration,
		increase.ProviderName:      increase.Registration,
		mangopay.ProviderName:      mangopay.Registration,
		modulr.ProviderName:        modulr.Registration,
		moneycorp.ProviderName:     moneycorp.Registration,
		plaid.ProviderName:         plaid.Registration,
		powens.ProviderName:        powens.Registration,
		qonto.ProviderName:         qonto.Registration,
		stripe.ProviderName:        stripe.Registration,
		tink.ProviderName:          tink.Registration,
		wise.ProviderName:          wise.Registration,
	})
}
