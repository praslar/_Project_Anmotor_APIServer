package bike

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/anmotor/internal/app/types"
	"github.com/anmotor/internal/pkg/http/respond"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type (
	service interface {
		Create(ctx context.Context, req *types.CreateBike) (*types.Bike, error)
		Update(ctx context.Context, id string, req types.UpdateBike) (*types.UpdateBike, error)
		Delete(ctx context.Context, id string) error

		FindAll(ctx context.Context) ([]*types.Bike, error)
		FindByID(context.Context, string) (*types.Bike, error)
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

	bike, err := h.srv.Create(r.Context(), &req)
	if err != nil {
		logrus.Errorf("Fail to Create bike due to, %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: bike,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["bike_id"]

	if id == "" {
		logrus.Error("Fail to Update Bike due to empty Bike ID ")
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}

	var req types.UpdateBike

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logrus.Errorf("Fail to parse JSON to Update Bike Request struct, %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	bike, err := h.srv.Update(r.Context(), id, req)
	if err != nil {
		logrus.Errorf("Fail to Update bike due to, %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: bike,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bike_id"]
	if id == "" {
		logrus.Error("Fail to delete bike due to empty bike ID ")
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	if err := h.srv.Delete(r.Context(), id); err != nil {
		logrus.Errorf("Fail to delete bike due to %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {

	bikes, err := h.srv.FindAll(r.Context())

	if err != nil {
		logrus.Errorf("Fail to get all bike due to, %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: bikes,
	})
}
func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["bike_id"]
	if id == "" {
		logrus.Error("Fail to get bike due to empty bike ID ")
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}

	bike, err := h.srv.FindByID(r.Context(), id)

	if err != nil {
		logrus.Errorf("Fail to get bike due to, %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}

	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: bike,
	})
}
