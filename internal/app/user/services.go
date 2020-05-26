package user

import (
	"context"

	"github.com/anmotor/internal/app/status"
	"github.com/anmotor/internal/app/types"
	db "github.com/anmotor/internal/pkg/database"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type (
	repoProvider interface {
		FindUser(ctx context.Context, username string) (*types.User, error)
	}

	Service struct {
		Repo repoProvider
	}
)

func NewService(repo repoProvider) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) AuthenUser(ctx context.Context, username, password string) (*types.User, error) {
	user, err := s.Repo.FindUser(ctx, username)
	if err != nil && !db.IsErrNotFound(err) {
		logrus.Errorf("failed to check existing user by username, err: %v", err)
		return nil, status.Gen().Internal
	}
	if db.IsErrNotFound(err) {
		logrus.Debugf("user not found, username: %s", username)
		return nil, status.User().UserNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logrus.Error("invalid password")
		return nil, status.Auth().InvalidUserPassword
	}
	return user.Strip(), nil
}
