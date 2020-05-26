package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Claims struct {
		jwt.StandardClaims
		UserName string
	}

	Config struct {
		JWTSecret string `envconfig:"JWT_SECRET" default:"194olaff~~~"`
	}

	Generator struct {
		Config        Config
		SigningMethod jwt.SigningMethod
	}

	StandardClaims = jwt.StandardClaims

	Signer interface {
		Sign(claims Claims) (string, error)
	}

	Verifier interface {
		Verify(tokenString string) (*Claims, error)
	}

	SignVerifier interface {
		Signer
		Verifier
	}
)

const (
	// DefaultIssuer is default issuer name
	DefaultIssuer = "anmotor"
	// DefaultLifeTime is default life time of a token
	DefaultLifeTime = time.Hour * 24
)

var (
	// ErrInvalidToken report that the JWT token is invalid
	ErrInvalidToken = errors.New("invalid token")
)

func New(conf Config) *Generator {
	return &Generator{
		Config:        conf,
		SigningMethod: jwt.SigningMethodES256,
	}
}

func (g *Generator) Sign(claims Claims) (string, error) {
	token := jwt.NewWithClaims(g.SigningMethod, claims)
	str, err := token.SignedString([]byte(g.Config.JWTSecret))
	return str, err
}

func (g *Generator) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.Config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
