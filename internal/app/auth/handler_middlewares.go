package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/anmotor/internal/app/types"
	"github.com/anmotor/internal/pkg/http/respond"
	"github.com/anmotor/internal/pkg/jwt"

	"github.com/sirupsen/logrus"
)

type (
	contextKey string
)

const (
	authContextKey contextKey = "r_authorized_user"
)

func UserInfoMiddleware(verifier jwt.Verifier) func(http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("Authorization")
			if key == "" {
				inner.ServeHTTP(w, r)
				return
			}
			claims, err := verifier.Verify(key)
			if err != nil {
				logrus.WithContext(r.Context()).Errorf("invalid JWT token, err: %v", err)
				inner.ServeHTTP(w, r)
				return
			}
			newCtx := NewContext(r.Context(), claimsToUser(claims))
			r = r.WithContext(newCtx)
			logrus.WithContext(r.Context()).WithField("user_id", claims.UserName)
			logrus.WithContext(r.Context()).Debugf("decode jwt successful")
			inner.ServeHTTP(w, r)
		})
	}
}

func RequireAuthMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := FromContext(r.Context()); user == nil {
			logrus.Errorf("unauthorized, err: %v", http.StatusUnauthorized)
			respond.Error(w, errors.New("Unauthorized"), http.StatusUnauthorized)
			return
		}
		inner.ServeHTTP(w, r)
	})
}

func NewContext(ctx context.Context, user *types.User) context.Context {
	return context.WithValue(ctx, authContextKey, user)
}

func FromContext(ctx context.Context) *types.User {
	if v, ok := ctx.Value(authContextKey).(*types.User); ok {
		return v
	}
	return nil
}
