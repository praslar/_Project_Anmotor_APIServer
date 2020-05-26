package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/anmotor/internal/app/types"
	"github.com/anmotor/internal/pkg/http/respond"
	"github.com/sirupsen/logrus"
)

type (
	service interface {
		Auth(ctx context.Context, username, password string) (string, *types.User, error)
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
	token, user, err := h.srv.Auth(r.Context(), req.Username, req.Password)
	if err != nil {
		logrus.Errorf("unauthorized, err: %v", http.StatusUnauthorized)
		respond.Error(w, err, http.StatusUnauthorized)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: map[string]interface{}{
			"token":     token,
			"user_info": user,
		},
	})
}
