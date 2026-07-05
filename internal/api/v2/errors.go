package v2

import (
	"errors"
	"net/http"

	"github.com/hanzo-fi/payments/internal/connectors"
	"github.com/hanzo-fi/payments/internal/connectors/engine"

	"github.com/formancehq/go-libs/v5/pkg/transport/api"
	"github.com/hanzo-fi/payments/internal/api/common"
	"github.com/hanzo-fi/payments/internal/api/services"
	"github.com/hanzo-fi/payments/internal/storage"
	errorsutils "github.com/hanzo-fi/payments/pkg/domain/errors"
)

const (
	ErrUniqueReference                 = "CONFLICT"
	ErrNotFound                        = "NOT_FOUND"
	ErrInvalidID                       = "INVALID_ID"
	ErrMissingOrInvalidBody            = "MISSING_OR_INVALID_BODY"
	ErrValidation                      = "VALIDATION"
	ErrConnectorCapabilityNotSupported = "CONNECTOR_CAPABILITY_NOT_SUPPORTED"
)

func handleServiceErrors(w http.ResponseWriter, r *http.Request, err error) {
	var capabilityNotSupported *engine.ErrConnectorCapabilityNotSupported

	switch {
	case errors.Is(err, storage.ErrDuplicateKeyValue):
		api.BadRequest(w, ErrUniqueReference, err)
	case errors.Is(err, storage.ErrNotFound):
		api.NotFound(w, err)
	case errors.Is(err, storage.ErrValidation):
		api.BadRequest(w, ErrValidation, err)
	case errors.Is(err, services.ErrValidation), errors.Is(err, connectors.ErrValidation):
		cause := errorsutils.Cause(err)
		api.BadRequest(w, ErrValidation, cause)
	case errors.Is(err, services.ErrNotFound):
		api.NotFound(w, err)
	case errors.As(err, &capabilityNotSupported):
		api.BadRequest(w, ErrConnectorCapabilityNotSupported, err)
	default:
		common.InternalServerError(w, r, err)
	}
}
