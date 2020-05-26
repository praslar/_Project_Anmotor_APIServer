package auth

import (
	"time"

	"github.com/anmotor/internal/app/types"
	"github.com/anmotor/internal/pkg/jwt"
)

func userToClaims(user *types.User, lifetime time.Duration) jwt.Claims {
	return jwt.Claims{
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(lifetime).Unix(),
			Issuer:    jwt.DefaultIssuer,
			Subject:   user.UserName,
		},
	}
}

func claimsToUser(claims *jwt.Claims) *types.User {
	return &types.User{
		UserName: claims.UserName,
	}
}