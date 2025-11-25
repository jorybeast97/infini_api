package authservice

import (
    jwt "github.com/golang-jwt/jwt/v5"
    "time"
)

type TokenService interface {
    Sign(userID, name, role string) (string, error)
    Parse(tokenStr string) (string, string, string, error)
}

type JWTService struct{ secret []byte }

func NewJWTService(secret string) *JWTService { return &JWTService{secret: []byte(secret)} }

func (j *JWTService) Sign(userID, name, role string) (string, error) {
    claims := jwt.MapClaims{"sub": userID, "name": name, "role": role, "exp": time.Now().Add(24*time.Hour).Unix(), "iat": time.Now().Unix()}
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return t.SignedString(j.secret)
}

func (j *JWTService) Parse(tokenStr string) (string, string, string, error) {
    tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) { return j.secret, nil })
    if err != nil || !tkn.Valid { return "", "", "", err }
    c := tkn.Claims.(jwt.MapClaims)
    return c["sub"].(string), c["name"].(string), c["role"].(string), nil
}
