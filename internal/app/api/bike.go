package api

import (
	"github.com/anmotor/internal/app/bike"
	"github.com/sirupsen/logrus"
)

func newBikeService() (*bike.Service, error) {
	s, err := dialDefaultMongoDB()
	if err != nil {
		logrus.Errorf("Fail to dial defautl mongodb API util")
		return nil, err
	}

	repo := bike.NewMongoDBRespository(s)

	return bike.New(repo), nil
}

func newBikeHandler(srv *bike.Service) *bike.Handler {
	return bike.NewHandler(srv)
}
