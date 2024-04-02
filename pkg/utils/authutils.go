package utils

import (
	"alumni-management-server/pkg/config"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// GetPasswordHash returns the hashed password.
func GetPasswordHash(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

// CheckPassword checks if the password is correct.
func CheckPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

// GetJwtForUser returns the JWT for the user using username.
func GetJwtForUser(username string) (string, error) {
	now := time.Now().UTC()
	ttl := time.Minute * time.Duration(config.LocalConfig.JwtExpireMinutes)
	claims := jwt.StandardClaims{
		ExpiresAt: now.Add(ttl).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		Subject:   username,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.LocalConfig.JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
