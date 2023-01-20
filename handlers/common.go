package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go-mongo-starter/models"
	"os"
	"time"
)

const (
	JwtExpirationTime = time.Hour * 1
)

func generateToken(u *models.User) (*string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{}
	claims["id"] = u.ID
	claims["authorized"] = true
	claims["iss"] = "digital-farm-api"
	claims["exp"] = time.Now().Add(JwtExpirationTime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &signed, nil
}

func ExtractClaims(c *fiber.Ctx) (*jwt.MapClaims, error) {
	// get token from header
	header := c.Get("Authorization")
	if header == "" {
		return nil, fmt.Errorf("token not found")
	}

	bearer := "Bearer "
	if len(header) < len(bearer) || header[:len(bearer)] != bearer {
		return nil, fmt.Errorf("invalid token")
	}

	tokenString := header[len(bearer):]

	// validate token
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		panic("env var 'JWT_SECRET' not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return &claims, nil
}

func ExtractUserID(c *fiber.Ctx) (*string, error) {
	claims, err := ExtractClaims(c)
	if err != nil {
		return nil, err
	}

	id, ok := (*claims)["id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	return &id, nil
}

func GetParam(c *fiber.Ctx, key string) (*string, error) {
	param := c.Params(key, "")
	if param == "" {
		return nil, fmt.Errorf("invalid param: %s", key)
	}
	return &param, nil
}
