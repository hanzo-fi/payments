package v3

import (
	"net/http"

	"github.com/hanzo-fi/go-libs/v5/pkg/transport/api"
	"github.com/hanzo-fi/payments/internal/api/backend"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/otel"
	"go.opentelemetry.io/otel/attribute"
)

func connectorsReset(backend backend.Backend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer().Start(r.Context(), "v3_connectorsReset")
		defer span.End()

		span.SetAttributes(attribute.String("connectorID", connectorID(r)))
		connectorID, err := models.ConnectorIDFromString(connectorID(r))
		if err != nil {
			otel.RecordError(span, err)
			api.BadRequest(w, ErrInvalidID, err)
			return
		}

		task, err := backend.ConnectorsReset(ctx, connectorID)
		if err != nil {
			otel.RecordError(span, err)
			handleServiceErrors(w, r, err)
			return
		}

		api.Accepted(w, task.ID.String())
	}
}
