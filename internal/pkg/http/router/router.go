package router

import (
	"net/http"

	"github.com/gorilla/mux"

	envconfig "github.com/anmotor/internal/pkg/env"
)

type (
	// Config hold configurations of router
	Config struct {

		// need to set manually
		Middlewares     []Middleware
		Routes          []Route
		NotFoundHandler http.Handler
	}

	Middleware = func(http.Handler) http.Handler

	// Route hold configuration of routing
	Route struct {
		Desc        string
		Path        string
		Method      string
		Queries     []string
		Handler     http.HandlerFunc
		Middlewares []Middleware
	}
)

func New(conf *Config) (http.Handler, error) {

	r := mux.NewRouter()

	for _, middleware := range conf.Middlewares {
		r.Use(middleware)
	}

	for _, rt := range conf.Routes {
		var h http.Handler
		h = http.HandlerFunc(rt.Handler)

		for i := len(rt.Middlewares) - 1; i >= 0; i-- {
			h = rt.Middlewares[i](h)
		}
		r.Path(rt.Path).Methods(rt.Method).Handler(h).Queries(rt.Queries...)
	}

	if conf.NotFoundHandler != nil {
		r.NotFoundHandler = conf.NotFoundHandler
	}

	return r, nil
}

func LoadConfigFromEnv() *Config {
	var conf Config
	envconfig.Load(&conf)
	return &conf
}
