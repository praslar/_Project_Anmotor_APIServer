package user

import (
	"context"
	"fmt"

	"github.com/anmotor/internal/app/types"
	db "github.com/anmotor/internal/pkg/database"

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
		return nil, fmt.Errorf("internal server error, %v", err)
	}

	if db.IsErrNotFound(err) {
		return nil, fmt.Errorf("user not found, %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("internal error")
	}
	return user.Strip(), nil
}
