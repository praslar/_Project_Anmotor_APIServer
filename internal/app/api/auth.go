package api

import (
	"github.com/anmotor/internal/app/auth"
	envconfig "github.com/anmotor/internal/pkg/env"
	"github.com/anmotor/internal/pkg/jwt"
)

func newAuthHandler(signer jwt.Signer, authenticator auth.UserAuthen) *auth.Handler {
	srv := auth.NewService(signer, authenticator)
	return auth.NewHandler(srv)
}

func newJWTSignVerifier() jwt.SignVerifier {
	var conf jwt.Config
	envconfig.Load(&conf)
	return jwt.New(conf)
}
