package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateTokenPair(userId string) (TokenPair, error)
}

type authServiceImpl struct {
	jwtSecret []byte
}

func NewService(jwtSecret []byte) Service {
	return &authServiceImpl{
		jwtSecret: jwtSecret,
	}
}

func (as *authServiceImpl) GenerateTokenPair(userId string) (TokenPair, error) {
	accessTokenClaims := jwt.RegisteredClaims{
		Issuer:    "auth-api.com",
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

	refreshTokenClaims := jwt.RegisteredClaims{
		Issuer:    "auth-api.com",
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	signedAccessToken, err := accessToken.SignedString(as.jwtSecret)
	if err != nil {
		return TokenPair{}, fmt.Errorf("Service.GenerateToken - Error signing access token: %w", err)
	}

	signedRefreshToken, err := refreshToken.SignedString(as.jwtSecret)
	if err != nil {
		return TokenPair{}, fmt.Errorf("Service.GenerateToken - Error signing refresh token: %w", err)
	}

	tokenPair := TokenPair{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return tokenPair, nil
}
