package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTProvider struct {
	secretKey string
	issuer    string
	expiry    time.Duration
}

func NewJWTProvider(secretKey, issuer string, expiry time.Duration) *JWTProvider {
	return &JWTProvider{
		secretKey: secretKey,
		issuer:    issuer,
		expiry:    expiry,
	}
}

func (j *JWTProvider) GenerateToken(userID string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"roles": roles,
		"iss":   j.issuer,
		"exp":   time.Now().Add(j.expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTProvider) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid token subject")
	}

	return userID, nil
}