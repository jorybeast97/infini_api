package mysql

import (
    "context"
    "time"
    "infini_api/src/domain"
    "gorm.io/gorm"
)

type PostsRepo struct{ db *gorm.DB }

func NewPostsRepo(db *gorm.DB) *PostsRepo { return &PostsRepo{db: db} }

func (r *PostsRepo) List(ctx context.Context, q, status string, hasLocation *bool, sort string, page, limit int) ([]domain.BlogPost, domain.Meta, error) {
    var items []domain.BlogPost
    tx := r.db.WithContext(ctx)
    if status != "" { tx = tx.Where("status = ?", status) }
    if hasLocation != nil {
        if *hasLocation { tx = tx.Where("location IS NOT NULL AND location != ''") } else { tx = tx.Where("location IS NULL OR location = ''") }
    }
    if q != "" { tx = tx.Where("title LIKE ? OR excerpt LIKE ? OR content LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%") }
    if sort != "" { tx = tx.Order(sort) }
    var total int64
    if err := tx.Model(&domain.BlogPost{}).Count(&total).Error; err != nil { return nil, domain.Meta{}, err }
    if err := tx.Offset((page-1)*limit).Limit(limit).Find(&items).Error; err != nil { return nil, domain.Meta{}, err }
    return items, domain.Meta{Page: page, Limit: limit, Total: int(total)}, nil
}

func (r *PostsRepo) Get(ctx context.Context, id string) (domain.BlogPost, error) {
    var p domain.BlogPost
    err := r.db.WithContext(ctx).First(&p, "id = ?", id).Error
    return p, err
}

func (r *PostsRepo) Create(ctx context.Context, p domain.BlogPost) (domain.BlogPost, error) {
    if p.ID == "" { p.ID = genID() }
    now := time.Now().Unix()
    if p.CreatedAt == 0 { p.CreatedAt = now }
    p.UpdatedAt = now
    err := r.db.WithContext(ctx).Create(&p).Error
    return p, err
}

func (r *PostsRepo) Update(ctx context.Context, id string, p domain.BlogPost) (domain.BlogPost, error) {
    p.ID = id
    p.UpdatedAt = time.Now().Unix()
    err := r.db.WithContext(ctx).Model(&domain.BlogPost{}).Where("id = ?", id).Updates(p).Error
    return p, err
}

func (r *PostsRepo) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&domain.BlogPost{}, "id = ?", id).Error
}
