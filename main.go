package main

import (
	"github.com/anmotor/internal/app/api"
	"github.com/anmotor/internal/pkg/http/server"
	"github.com/sirupsen/logrus"
)

func main() {
	router, err := api.NewRouter()
	if err != nil {
		logrus.Panic("Cannot init Router, err: ", err)
	}
	severConf := server.LoadConfigFromEnv()
	server.ListenAndServe(severConf, router)
}
