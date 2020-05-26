package auth

import (
	"encoding/json"
	"net/http"

	"github.com/anmotor/internal/pkg/http/respond"
)

type (
	service interface {
	}

	Handler struct {
		srv service
	}
)

func NewHandler(srv service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Username string
		Password string
	}{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
}
