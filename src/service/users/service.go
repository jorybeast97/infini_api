package usersservice

import (
    "context"
    "time"
    "infini_api/src/dao"
    "infini_api/src/domain"
    "golang.org/x/crypto/bcrypt"
)

type UsersService interface {
    List(ctx context.Context, q string, page, limit int) ([]domain.User, domain.Meta, error)
    Get(ctx context.Context, id string) (domain.User, error)
    Create(ctx context.Context, u domain.User, password string) (domain.User, error)
    Update(ctx context.Context, id string, u domain.User, password *string) (domain.User, error)
    Delete(ctx context.Context, id string) error
}

type UsersServiceImpl struct{ repo dao.UsersRepository }

func NewUsersService(repo dao.UsersRepository) *UsersServiceImpl { return &UsersServiceImpl{repo: repo} }

func (s *UsersServiceImpl) List(ctx context.Context, q string, page, limit int) ([]domain.User, domain.Meta, error) { return s.repo.List(ctx, q, page, limit) }
func (s *UsersServiceImpl) Get(ctx context.Context, id string) (domain.User, error) { return s.repo.Get(ctx, id) }

func (s *UsersServiceImpl) Create(ctx context.Context, u domain.User, password string) (domain.User, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil { return domain.User{}, err }
    u.PasswordHash = string(hash)
    if u.CreatedAt == 0 { u.CreatedAt = time.Now().Unix() }
    u.UpdatedAt = u.CreatedAt
    return s.repo.Create(ctx, u)
}

func (s *UsersServiceImpl) Update(ctx context.Context, id string, u domain.User, password *string) (domain.User, error) {
    if password != nil {
        hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
        if err != nil { return domain.User{}, err }
        u.PasswordHash = string(hash)
    }
    u.UpdatedAt = time.Now().Unix()
    return s.repo.Update(ctx, id, u)
}

func (s *UsersServiceImpl) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }

