package api

import (
	"net/http"

	"github.com/anmotor/internal/app/auth"
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

	userSrv, err := newUserService()
	if err != nil {
		return nil, err
	}

	//===========================================================
	indexHandler := NewIndexHandler()
	jwtSignVerifier := newJWTSignVerifier()
	userInfoMiddleware := auth.UserInfoMiddleware(jwtSignVerifier)
	authHandler := newAuthHandler(jwtSignVerifier, userSrv)

	routes := []router.Route{
		{
			Path:    "/",
			Method:  get,
			Handler: indexHandler.ServeHTTP,
		},
	}

	routes = append(routes, bikeHandler.Routes()...)
	routes = append(routes, authHandler.Routes()...)

	conf := router.LoadConfigFromEnv()
	conf.Routes = routes

	conf.Middlewares = []router.Middleware{
		userInfoMiddleware,
	}

	r, err := router.New(conf)

	if err != nil {
		return nil, err
	}
	return middleware.CORS(r), nil
}
