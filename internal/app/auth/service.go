package auth

import (
	"context"

	"github.com/anmotor/internal/app/status"
	"github.com/anmotor/internal/app/types"
	"github.com/anmotor/internal/pkg/jwt"

	"github.com/sirupsen/logrus"
)

type (
	UserAuthen interface {
		AuthenUser(ctx context.Context, username, password string) (*types.User, error)
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

func (s *Service) Auth(ctx context.Context, username, password string) (string, *types.User, error) {
	user, err := s.authentication.AuthenUser(ctx, username, password)
	if err != nil {
		logrus.WithContext(ctx).Errorf("fail to login with %s, err: %#v", username, err)
		return "", nil, status.Gen().BadRequest
	}

	token, err := s.jwtSigner.Sign(userToClaims(user, jwt.DefaultLifeTime))

	if err != nil {
		logrus.WithContext(ctx).Errorf("fail to gerenate JWT token, err: %#v", err)
		return "", nil, err
	}

	return token, user, nil
}
