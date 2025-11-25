package postsservice

import (
    "context"
    "infini_api/src/dao"
    "infini_api/src/domain"
)

type PostsService interface {
    List(ctx context.Context, q, status string, hasLocation *bool, sort string, page, limit int) ([]domain.BlogPost, domain.Meta, error)
    Get(ctx context.Context, id string) (domain.BlogPost, error)
    Create(ctx context.Context, p domain.BlogPost) (domain.BlogPost, error)
    Update(ctx context.Context, id string, p domain.BlogPost) (domain.BlogPost, error)
    Delete(ctx context.Context, id string) error
}

type PostsServiceImpl struct{ repo dao.PostsRepository }

func NewPostsService(repo dao.PostsRepository) *PostsServiceImpl { return &PostsServiceImpl{repo: repo} }

func (s *PostsServiceImpl) List(ctx context.Context, q, status string, hasLocation *bool, sort string, page, limit int) ([]domain.BlogPost, domain.Meta, error) { return s.repo.List(ctx, q, status, hasLocation, sort, page, limit) }
func (s *PostsServiceImpl) Get(ctx context.Context, id string) (domain.BlogPost, error) { return s.repo.Get(ctx, id) }
func (s *PostsServiceImpl) Create(ctx context.Context, p domain.BlogPost) (domain.BlogPost, error) { return s.repo.Create(ctx, p) }
func (s *PostsServiceImpl) Update(ctx context.Context, id string, p domain.BlogPost) (domain.BlogPost, error) { return s.repo.Update(ctx, id, p) }
func (s *PostsServiceImpl) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }
