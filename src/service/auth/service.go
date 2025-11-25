package authservice

import (
    "context"
    "strings"
    "infini_api/src/domain"
)

type AuthService interface {
    Login(ctx context.Context, username, password string) (string, string, string, error)
}

type AuthServiceImpl struct{ tokens TokenService }

func NewAuthService(tokens TokenService) *AuthServiceImpl { return &AuthServiceImpl{tokens: tokens} }

func (s *AuthServiceImpl) Login(ctx context.Context, username, password string) (string, string, string, error) {
    if username == "" || password == "" { return "", "", "", domain.AppError{Code: "INVALID_INPUT", Message: "missing credentials"} }
    if strings.ToLower(username) != "admin" || password != "password" { return "", "", "", domain.AppError{Code: "UNAUTHORIZED", Message: "invalid credentials"} }
    t, err := s.tokens.Sign("1", "Admin", "admin")
    if err != nil { return "", "", "", err }
    return t, "Admin", "admin", nil
}
