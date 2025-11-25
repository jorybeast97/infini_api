package authservice

import (
    "context"
    "infini_api/src/dao"
    "infini_api/src/domain"
    "golang.org/x/crypto/bcrypt"
)

type AuthService interface {
    Login(ctx context.Context, username, password string) (string, domain.User, error)
}

type AuthServiceImpl struct{ tokens TokenService; users dao.UsersRepository }

func NewAuthService(tokens TokenService, users dao.UsersRepository) *AuthServiceImpl { return &AuthServiceImpl{tokens: tokens, users: users} }

func (s *AuthServiceImpl) Login(ctx context.Context, username, password string) (string, domain.User, error) {
    if username == "" || password == "" { return "", domain.User{}, domain.AppError{Code: "INVALID_INPUT", Message: "missing credentials"} }
    u, err := s.users.GetByUserName(ctx, username)
    if err != nil { return "", domain.User{}, domain.AppError{Code: "UNAUTHORIZED", Message: "invalid credentials"} }
    if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
        return "", domain.User{}, domain.AppError{Code: "UNAUTHORIZED", Message: "invalid credentials"}
    }
    t, err := s.tokens.Sign(u.ID, u.NickName, u.Role)
    if err != nil { return "", domain.User{}, err }
    u.PasswordHash = ""
    return t, u, nil
}
