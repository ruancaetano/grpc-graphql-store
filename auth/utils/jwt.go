package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ruancaetano/grpc-graphql-store/users/pbusers"
	"golang.org/x/exp/slices"
)

type JwtCustomClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateJwtUserToken(user *pbusers.User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	expiresAt := time.Now().Add(time.Hour).Unix()

	claims := JwtCustomClaims{
		Id:    user.GetId(),
		Email: user.GetEmail(),
		Name:  user.GetName(),
		Role:  user.GetRole(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateJwt(token string, expectedRoles []string) (*JwtCustomClaims, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	parsedToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*JwtCustomClaims)

	if !ok {
		return nil, fmt.Errorf("failed to get claims")
	}

	if !parsedToken.Valid || !slices.Contains(expectedRoles, claims.Role) {
		return nil, fmt.Errorf("permission denied")
	}

	return claims, nil
}
