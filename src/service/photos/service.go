package main

import "context"

type PhotosService interface {
    List(ctx context.Context, q string, sort string, page, limit int) ([]Photo, Meta, error)
    Get(ctx context.Context, id string) (Photo, error)
    Create(ctx context.Context, p Photo) (Photo, error)
    Update(ctx context.Context, id string, p Photo) (Photo, error)
    Delete(ctx context.Context, id string) error
}

type PhotosServiceImpl struct{ repo PhotosRepository }

func NewPhotosService(repo PhotosRepository) *PhotosServiceImpl { return &PhotosServiceImpl{repo: repo} }

func (s *PhotosServiceImpl) List(ctx context.Context, q string, sort string, page, limit int) ([]Photo, Meta, error) { return s.repo.List(ctx, q, sort, page, limit) }
func (s *PhotosServiceImpl) Get(ctx context.Context, id string) (Photo, error) { return s.repo.Get(ctx, id) }
func (s *PhotosServiceImpl) Create(ctx context.Context, p Photo) (Photo, error) { return s.repo.Create(ctx, p) }
func (s *PhotosServiceImpl) Update(ctx context.Context, id string, p Photo) (Photo, error) { return s.repo.Update(ctx, id, p) }
func (s *PhotosServiceImpl) Delete(ctx context.Context, id string) error { return s.repo.Delete(ctx, id) }

