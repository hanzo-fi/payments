package stripe

import (
	"fmt"

	"github.com/hanzo-fi/payments/pkg/domain/models"
	errorsutils "github.com/hanzo-fi/payments/pkg/domain/errors"
)

func (p *Plugin) validatePayoutTransferRequest(pi models.PSPPaymentInitiation) error {
	if pi.DestinationAccount == nil {
		return errorsutils.NewWrappedError(
			fmt.Errorf("destination account is required in transfer/payout request"),
			models.ErrInvalidRequest,
		)
	}

	return nil
}
