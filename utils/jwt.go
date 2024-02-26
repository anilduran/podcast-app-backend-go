package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret []byte

func GenerateToken(userId uuid.UUID, email string) (string, error) {

	secretString := os.Getenv("JWT_SECRET")

	if secretString == "" {
		return "", errors.New("jwt_secret is not found")
	}

	secret = []byte(secretString)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId.String(),
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretString)

}

func VerifyToken(token string) (uuid.UUID, error) {

	secretString := os.Getenv("JWT_SECRET")

	if secretString == "" {
		return uuid.Nil, errors.New("jwt_secret is not found")
	}

	secret = []byte(secretString)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})

	if err != nil {
		return uuid.Nil, errors.New("token is not valid")
	}

	if !parsedToken.Valid {
		return uuid.Nil, errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return uuid.Nil, errors.New("couldn't parse map claims")
	}

	userId, _ := uuid.Parse(claims["userId"].(string))

	return userId, nil

}
