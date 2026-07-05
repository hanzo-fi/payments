package v3

import (
	"encoding/json"
	"net/http"

	"github.com/formancehq/go-libs/v5/pkg/transport/api"
	"github.com/hanzo-fi/payments/internal/api/backend"
	"github.com/hanzo-fi/payments/internal/api/common"
	"github.com/hanzo-fi/payments/internal/connectors/plugins/registry"
	"github.com/hanzo-fi/payments/internal/otel"
)

func connectorsConfigs(backend backend.Backend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := otel.Tracer().Start(r.Context(), "v3_connectorsConfigs")
		defer span.End()

		confs := backend.ConnectorsConfigs()

		err := json.NewEncoder(w).Encode(api.BaseResponse[registry.Configs]{
			Data: &confs,
		})
		if err != nil {
			otel.RecordError(span, err)
			common.InternalServerError(w, r, err)
			return
		}
	}
}
