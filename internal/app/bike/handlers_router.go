package bike

import (
	"net/http"

	"github.com/anmotor/internal/app/auth"
	"github.com/anmotor/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:        "/api/v1/bike",
			Method:      http.MethodPost,
			Handler:     h.Create,
			Middlewares: []router.Middleware{auth.RequireAuthMiddleware},
		},
	}
}
