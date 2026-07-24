package v3

import (
	"net/http"

	"github.com/hanzo-fi/go-libs/v5/pkg/transport/api"
	"github.com/hanzo-fi/payments/internal/api/backend"
	"github.com/hanzo-fi/payments/pkg/domain/models"
	"github.com/hanzo-fi/payments/internal/otel"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

type PaymentServiceUserDeleteConnectionResponse struct {
	TaskID string `json:"taskID"`
}

func paymentServiceUsersDeleteConnection(backend backend.Backend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer().Start(r.Context(), "v3_paymentServiceUsersDeleteConnection")
		defer span.End()

		span.SetAttributes(attribute.String("paymentServiceUserID", paymentServiceUserID(r)))
		id, err := uuid.Parse(paymentServiceUserID(r))
		if err != nil {
			otel.RecordError(span, err)
			api.BadRequest(w, ErrInvalidID, err)
			return
		}

		span.SetAttributes(attribute.String("connectorID", connectorID(r)))
		connectorID, err := models.ConnectorIDFromString(connectorID(r))
		if err != nil {
			otel.RecordError(span, err)
			api.BadRequest(w, ErrInvalidID, err)
			return
		}

		connectionID := connectionID(r)
		span.SetAttributes(attribute.String("connectionID", connectionID))

		task, err := backend.PaymentServiceUsersConnectionsDelete(ctx, connectorID, id, connectionID)
		if err != nil {
			otel.RecordError(span, err)
			handleServiceErrors(w, r, err)
			return
		}

		api.Accepted(w, PaymentServiceUserDeleteConnectionResponse{
			TaskID: task.ID.String(),
		})
	}
}
