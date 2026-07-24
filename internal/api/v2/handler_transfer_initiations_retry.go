package v2

import (
	"net/http"

	"github.com/hanzo-fi/go-libs/v5/pkg/transport/api"
	"github.com/hanzo-fi/payments/internal/api/backend"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/otel"
	"go.opentelemetry.io/otel/attribute"
)

func transferInitiationsRetry(backend backend.Backend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer().Start(r.Context(), "v2_transferInitiationsRetry")
		defer span.End()

		span.SetAttributes(attribute.String("transferInitiationID", transferInitiationID(r)))
		id, err := models.PaymentInitiationIDFromString(transferInitiationID(r))
		if err != nil {
			otel.RecordError(span, err)
			api.BadRequest(w, ErrInvalidID, err)
			return
		}

		_, err = backend.PaymentInitiationsRetry(ctx, id, true)
		if err != nil {
			otel.RecordError(span, err)
			handleServiceErrors(w, r, err)
			return
		}

		api.NoContent(w)
	}
}
