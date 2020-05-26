package auth

import (
	"github.com/anmotor/internal/pkg/jwt"
)

type (
	UserAuthen interface {
	}

	Service struct {
		jwtSigner      jwt.Signer
		authentication UserAuthen
	}
)

func NewService(signer jwt.Signer, authentication UserAuthen) *Service {
	return &Service{
		jwtSigner:      signer,
		authentication: authentication,
	}
}
