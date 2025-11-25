package dao

import (
    "context"
    "infini_api/src/domain"
)

type AuthorsRepository interface {
	List(ctx context.Context, q string, page, limit int) ([]domain.Author, domain.Meta, error)
	Get(ctx context.Context, id string) (domain.Author, error)
	Create(ctx context.Context, a domain.Author) (domain.Author, error)
	Update(ctx context.Context, id string, a domain.Author) (domain.Author, error)
	Delete(ctx context.Context, id string) error
}

type PostsRepository interface {
	List(ctx context.Context, q, status string, hasLocation *bool, sort string, page, limit int) ([]domain.BlogPost, domain.Meta, error)
	Get(ctx context.Context, id string) (domain.BlogPost, error)
	Create(ctx context.Context, p domain.BlogPost) (domain.BlogPost, error)
	Update(ctx context.Context, id string, p domain.BlogPost) (domain.BlogPost, error)
	Delete(ctx context.Context, id string) error
}

type PhotosRepository interface {
	List(ctx context.Context, q string, sort string, page, limit int) ([]domain.Photo, domain.Meta, error)
	Get(ctx context.Context, id string) (domain.Photo, error)
	Create(ctx context.Context, p domain.Photo) (domain.Photo, error)
	Update(ctx context.Context, id string, p domain.Photo) (domain.Photo, error)
	Delete(ctx context.Context, id string) error
}

type AppsRepository interface {
    List(ctx context.Context) ([]domain.AppProject, error)
    Get(ctx context.Context, id string) (domain.AppProject, error)
    Create(ctx context.Context, a domain.AppProject) (domain.AppProject, error)
    Update(ctx context.Context, id string, a domain.AppProject) (domain.AppProject, error)
    Delete(ctx context.Context, id string) error
}

type UsersRepository interface {
    List(ctx context.Context, q string, page, limit int) ([]domain.User, domain.Meta, error)
    Get(ctx context.Context, id string) (domain.User, error)
    Create(ctx context.Context, u domain.User) (domain.User, error)
    Update(ctx context.Context, id string, u domain.User) (domain.User, error)
    Delete(ctx context.Context, id string) error
}
