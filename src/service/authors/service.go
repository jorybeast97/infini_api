package authorsservice

import (
    "context"
    "infini_api/src/dao"
    "infini_api/src/domain"
)

type AuthorsService interface {
    List(ctx context.Context, q string, page, limit int) ([]domain.Author, domain.Meta, error)
    Get(ctx context.Context, id string) (domain.Author, error)
    Create(ctx context.Context, a domain.Author) (domain.Author, error)
    Update(ctx context.Context, id string, a domain.Author) (domain.Author, error)
    Delete(ctx context.Context, id string) error
}

type AuthorsServiceImpl struct{ repo dao.AuthorsRepository }

func NewAuthorsService(repo dao.AuthorsRepository) *AuthorsServiceImpl { return &AuthorsServiceImpl{repo: repo} }

func (s *AuthorsServiceImpl) List(ctx context.Context, q string, page, limit int) ([]domain.Author, domain.Meta, error) { return s.repo.List(ctx, q, page, limit) }
func (s *AuthorsServiceImpl) Get(ctx context.Context, id string) (domain.Author, error) { return s.repo.Get(ctx, id) }
func (s *AuthorsServiceImpl) Create(ctx context.Context, a domain.Author) (domain.Author, error) { return s.repo.Create(ctx, a) }
func (s *AuthorsServiceImpl) Update(ctx context.Context, id string, a domain.Author) (domain.Author, error) { return s.repo.Update(ctx, id, a) }
func (s *AuthorsServiceImpl) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
