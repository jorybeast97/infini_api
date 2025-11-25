package mysql

import (
    "context"
    "time"
    "infini_api/src/domain"
    "gorm.io/gorm"
)

type PhotosRepo struct{ db *gorm.DB }

func NewPhotosRepo(db *gorm.DB) *PhotosRepo { return &PhotosRepo{db: db} }

func (r *PhotosRepo) List(ctx context.Context, q string, sort string, page, limit int) ([]domain.Photo, domain.Meta, error) {
    var items []domain.Photo
    tx := r.db.WithContext(ctx)
    if q != "" { tx = tx.Where("caption LIKE ? OR url LIKE ?", "%"+q+"%", "%"+q+"%") }
    if sort != "" { tx = tx.Order(sort) }
    var total int64
    if err := tx.Model(&domain.Photo{}).Count(&total).Error; err != nil { return nil, domain.Meta{}, err }
    if err := tx.Offset((page-1)*limit).Limit(limit).Find(&items).Error; err != nil { return nil, domain.Meta{}, err }
    return items, domain.Meta{Page: page, Limit: limit, Total: int(total)}, nil
}

func (r *PhotosRepo) Get(ctx context.Context, id string) (domain.Photo, error) {
    var p domain.Photo
    err := r.db.WithContext(ctx).First(&p, "id = ?", id).Error
    return p, err
}

func (r *PhotosRepo) Create(ctx context.Context, p domain.Photo) (domain.Photo, error) {
    if p.ID == "" { p.ID = genID() }
    now := time.Now().Unix()
    if p.CreatedAt == 0 { p.CreatedAt = now }
    p.UpdatedAt = now
    err := r.db.WithContext(ctx).Create(&p).Error
    return p, err
}

func (r *PhotosRepo) Update(ctx context.Context, id string, p domain.Photo) (domain.Photo, error) {
    p.ID = id
    p.UpdatedAt = time.Now().Unix()
    err := r.db.WithContext(ctx).Model(&domain.Photo{}).Where("id = ?", id).Updates(p).Error
    return p, err
}

func (r *PhotosRepo) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&domain.Photo{}, "id = ?", id).Error
}
