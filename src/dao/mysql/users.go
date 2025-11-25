package mysql

import (
    "context"
    "time"
    "infini_api/src/domain"
    "gorm.io/gorm"
)

type UsersRepo struct{ db *gorm.DB }

func NewUsersRepo(db *gorm.DB) *UsersRepo { return &UsersRepo{db: db} }

func (r *UsersRepo) List(ctx context.Context, q string, page, limit int) ([]domain.User, domain.Meta, error) {
    var items []domain.User
    tx := r.db.WithContext(ctx)
    if q != "" { tx = tx.Where("user_name LIKE ? OR nick_name LIKE ? OR bio LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%") }
    var total int64
    if err := tx.Model(&domain.User{}).Count(&total).Error; err != nil { return nil, domain.Meta{}, err }
    if err := tx.Offset((page-1)*limit).Limit(limit).Find(&items).Error; err != nil { return nil, domain.Meta{}, err }
    return items, domain.Meta{Page: page, Limit: limit, Total: int(total)}, nil
}

func (r *UsersRepo) Get(ctx context.Context, id string) (domain.User, error) {
    var u domain.User
    err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error
    return u, err
}

func (r *UsersRepo) GetByUserName(ctx context.Context, userName string) (domain.User, error) {
    var u domain.User
    err := r.db.WithContext(ctx).Where("user_name = ?", userName).First(&u).Error
    return u, err
}

func (r *UsersRepo) Create(ctx context.Context, u domain.User) (domain.User, error) {
    now := time.Now().Unix()
    if u.ID == "" { u.ID = genID() }
    if u.CreatedAt == 0 { u.CreatedAt = now }
    u.UpdatedAt = now
    err := r.db.WithContext(ctx).Create(&u).Error
    return u, err
}

func (r *UsersRepo) Update(ctx context.Context, id string, u domain.User) (domain.User, error) {
    u.ID = id
    u.UpdatedAt = time.Now().Unix()
    err := r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Updates(u).Error
    return u, err
}

func (r *UsersRepo) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id).Error
}
