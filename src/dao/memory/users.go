package memory

import (
    "context"
    "time"
    "infini_api/src/domain"
)

type UsersRepo struct{ s *MemoryStore }

func NewUsersRepo(s *MemoryStore) *UsersRepo { return &UsersRepo{s: s} }

func (r *UsersRepo) List(ctx context.Context, q string, page, limit int) ([]domain.User, domain.Meta, error) {
    r.s.mu.RLock(); var list []domain.User
    for _, u := range r.s.users { if q == "" || containsAny(u.UserName+" "+u.NickName+" "+u.Bio, []string{q}) { list = append(list, u) } }
    r.s.mu.RUnlock()
    data, meta := paginate(list, page, limit)
    return data, meta, nil
}

func (r *UsersRepo) Get(ctx context.Context, id string) (domain.User, error) {
    r.s.mu.RLock(); defer r.s.mu.RUnlock()
    for _, u := range r.s.users { if u.ID == id { return u, nil } }
    return domain.User{}, domain.AppError{Code: "NOT_FOUND", Message: "user not found"}
}

func (r *UsersRepo) GetByUserName(ctx context.Context, userName string) (domain.User, error) {
    r.s.mu.RLock(); defer r.s.mu.RUnlock()
    for _, u := range r.s.users { if u.UserName == userName { return u, nil } }
    return domain.User{}, domain.AppError{Code: "NOT_FOUND", Message: "user not found"}
}

func (r *UsersRepo) Create(ctx context.Context, u domain.User) (domain.User, error) {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    if u.ID == "" { u.ID = genID() }
    if u.CreatedAt == 0 { u.CreatedAt = time.Now().Unix() }
    u.UpdatedAt = u.CreatedAt
    r.s.users = append(r.s.users, u)
    return u, nil
}

func (r *UsersRepo) Update(ctx context.Context, id string, u domain.User) (domain.User, error) {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.users { if r.s.users[i].ID == id { u.ID = id; u.UpdatedAt = time.Now().Unix(); r.s.users[i] = u; return u, nil } }
    return domain.User{}, domain.AppError{Code: "NOT_FOUND", Message: "user not found"}
}

func (r *UsersRepo) Delete(ctx context.Context, id string) error {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.users { if r.s.users[i].ID == id { r.s.users = append(r.s.users[:i], r.s.users[i+1:]...); return nil } }
    return domain.AppError{Code: "NOT_FOUND", Message: "user not found"}
}
