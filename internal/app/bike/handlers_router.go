package bike

import (
	"net/http"

	"github.com/anmotor/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/bikes",
			Method:  http.MethodPost,
			Handler: h.Create,
			// Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
	}
}