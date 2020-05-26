package auth

import (
	"net/http"

	"github.com/anmotor/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:   "/authentication",
			Method: http.MethodPost,
			// Handler: ,
		},
	}
}
