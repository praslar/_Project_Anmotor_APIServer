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
		{
			Path:        "/api/v1/bike",
			Method:      http.MethodGet,
			Handler:     h.FindAll,
			Middlewares: []router.Middleware{auth.RequireAuthMiddleware},
		},
		{
			Path:        "/api/v1/bike/{bike_id:[a-z0-9-\\-]+}",
			Method:      http.MethodGet,
			Handler:     h.FindByID,
			Middlewares: []router.Middleware{auth.RequireAuthMiddleware},
		},
		{
			Path:        "/api/v1/bike/{bike_id:[a-z0-9-\\-]+}",
			Method:      http.MethodPatch,
			Handler:     h.Update,
			Middlewares: []router.Middleware{auth.RequireAuthMiddleware},
		},
		{
			Path:        "/api/v1/bike/{bike_id:[a-z0-9-\\-]+}",
			Method:      http.MethodDelete,
			Handler:     h.Delete,
			Middlewares: []router.Middleware{auth.RequireAuthMiddleware},
		},
	}
}
