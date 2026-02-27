package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type AuthService struct {
	secretKey string
}

func NewAuthService(secretKey string) *AuthService {
	return &AuthService{secretKey: secretKey}
}

func (a *AuthService) CreateTokens(userID string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(15 * time.Minute).Unix()
	td.RtExpires = time.Now().Add(7 * 24 * time.Hour).Unix()

	var err error
	td.AccessToken, err = a.createToken(userID, td.AtExpires)
	if err != nil {
		return nil, err
	}

	td.RefreshToken, err = a.createToken(userID, td.RtExpires)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (a *AuthService) createToken(userID string, expires int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = expires

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secretKey))
}

func (a *AuthService) ValidateRefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	userID := claims["user_id"].(string)
	return userID, nil
}