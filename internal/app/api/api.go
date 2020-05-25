package api

import (
	"net/http"

	"github.com/anmotor/internal/pkg/http/middleware"
	"github.com/anmotor/internal/pkg/http/router"
)

const (
	get     = http.MethodGet
	post    = http.MethodPost
	put     = http.MethodPut
	delete  = http.MethodDelete
	options = http.MethodOptions
)

func NewRouter() (http.Handler, error) {
	bikeSrv, err := newBikeService()
	if err != nil {
		return nil, err
	}
	bikeHandler := newBikeHandler(bikeSrv)

	indexHandler := NewIndexHandler()

	routes := []router.Route{
		{
			Path:    "/",
			Method:  get,
			Handler: indexHandler.ServeHTTP,
		},
	}

	routes = append(routes, bikeHandler.Routes()...)

	conf := router.LoadConfigFromEnv()
	conf.Routes = routes

	r, err := router.New(conf)

	if err != nil {
		return nil, err
	}
	return middleware.CORS(r), nil
}
