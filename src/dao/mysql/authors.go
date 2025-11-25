package mysql

import (
    "context"
    "infini_api/src/domain"
    "gorm.io/gorm"
)

type AuthorsRepo struct{ db *gorm.DB }

func NewAuthorsRepo(db *gorm.DB) *AuthorsRepo { return &AuthorsRepo{db: db} }

func (r *AuthorsRepo) List(ctx context.Context, q string, page, limit int) ([]domain.Author, domain.Meta, error) {
    var items []domain.Author
    tx := r.db.WithContext(ctx)
    if q != "" { tx = tx.Where("name LIKE ? OR role LIKE ? OR bio LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%") }
    var total int64
    if err := tx.Model(&domain.Author{}).Count(&total).Error; err != nil { return nil, domain.Meta{}, err }
    if err := tx.Offset((page-1)*limit).Limit(limit).Find(&items).Error; err != nil { return nil, domain.Meta{}, err }
    return items, domain.Meta{Page: page, Limit: limit, Total: int(total)}, nil
}

func (r *AuthorsRepo) Get(ctx context.Context, id string) (domain.Author, error) {
    var a domain.Author
    err := r.db.WithContext(ctx).First(&a, "id = ?", id).Error
    return a, err
}

func (r *AuthorsRepo) Create(ctx context.Context, a domain.Author) (domain.Author, error) {
    if a.ID == "" { a.ID = genID() }
    err := r.db.WithContext(ctx).Create(&a).Error
    return a, err
}

func (r *AuthorsRepo) Update(ctx context.Context, id string, a domain.Author) (domain.Author, error) {
    a.ID = id
    err := r.db.WithContext(ctx).Model(&domain.Author{}).Where("id = ?", id).Updates(a).Error
    return a, err
}

func (r *AuthorsRepo) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&domain.Author{}, "id = ?", id).Error
}

