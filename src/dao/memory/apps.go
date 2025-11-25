package memory

import (
    "context"
    "infini_api/src/domain"
)

type AppsRepo struct{ s *MemoryStore }

func NewAppsRepo(s *MemoryStore) *AppsRepo { return &AppsRepo{s: s} }

func (r *AppsRepo) List(ctx context.Context) ([]domain.AppProject, error) {
    r.s.mu.RLock(); list := append([]domain.AppProject(nil), r.s.apps...); r.s.mu.RUnlock(); return list, nil
}

func (r *AppsRepo) Get(ctx context.Context, id string) (domain.AppProject, error) {
    r.s.mu.RLock(); defer r.s.mu.RUnlock()
    for _, a := range r.s.apps { if a.ID == id { return a, nil } }
    return domain.AppProject{}, domain.AppError{Code: "NOT_FOUND", Message: "app not found"}
}

func (r *AppsRepo) Create(ctx context.Context, a domain.AppProject) (domain.AppProject, error) {
    r.s.mu.Lock(); a.ID = genID(); r.s.apps = append(r.s.apps, a); r.s.mu.Unlock(); return a, nil
}

func (r *AppsRepo) Update(ctx context.Context, id string, a domain.AppProject) (domain.AppProject, error) {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.apps { if r.s.apps[i].ID == id { a.ID = id; r.s.apps[i] = a; return a, nil } }
    return domain.AppProject{}, domain.AppError{Code: "NOT_FOUND", Message: "app not found"}
}

func (r *AppsRepo) Delete(ctx context.Context, id string) error {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.apps { if r.s.apps[i].ID == id { r.s.apps = append(r.s.apps[:i], r.s.apps[i+1:]...); return nil } }
    return domain.AppError{Code: "NOT_FOUND", Message: "app not found"}
}
