package bike

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
		Create(ctx context.Context, req *types.CreateBike) (*types.Bike, error)
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
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req types.CreateBike

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.Errorf("Fail to parse JSON to Create Bike Request struct, %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	project, err := h.srv.Create(r.Context(), &req)
	if err != nil {
		logrus.Errorf("Fail to Create Project due to, %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: project,
	})
}
