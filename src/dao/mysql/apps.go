package mysql

import (
    "context"
    "time"
    "infini_api/src/domain"
    "gorm.io/gorm"
)

type AppsRepo struct{ db *gorm.DB }

func NewAppsRepo(db *gorm.DB) *AppsRepo { return &AppsRepo{db: db} }

func (r *AppsRepo) List(ctx context.Context) ([]domain.AppProject, error) {
    var items []domain.AppProject
    err := r.db.WithContext(ctx).Find(&items).Error
    return items, err
}

func (r *AppsRepo) Get(ctx context.Context, id string) (domain.AppProject, error) {
    var a domain.AppProject
    err := r.db.WithContext(ctx).First(&a, "id = ?", id).Error
    return a, err
}

func (r *AppsRepo) Create(ctx context.Context, a domain.AppProject) (domain.AppProject, error) {
    if a.ID == "" { a.ID = genID() }
    now := time.Now().Unix()
    if a.CreatedAt == 0 { a.CreatedAt = now }
    a.UpdatedAt = now
    err := r.db.WithContext(ctx).Create(&a).Error
    return a, err
}

func (r *AppsRepo) Update(ctx context.Context, id string, a domain.AppProject) (domain.AppProject, error) {
    a.ID = id
    a.UpdatedAt = time.Now().Unix()
    err := r.db.WithContext(ctx).Model(&domain.AppProject{}).Where("id = ?", id).Updates(a).Error
    return a, err
}

func (r *AppsRepo) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&domain.AppProject{}, "id = ?", id).Error
}
