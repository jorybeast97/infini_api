package memory

import (
    "context"
    "infini_api/src/domain"
)

type AuthorsRepo struct{ s *MemoryStore }

func NewAuthorsRepo(s *MemoryStore) *AuthorsRepo { return &AuthorsRepo{s: s} }

func (r *AuthorsRepo) List(ctx context.Context, q string, page, limit int) ([]domain.Author, domain.Meta, error) {
    r.s.mu.RLock(); var list []domain.Author
    for _, a := range r.s.authors { if q == "" || containsAny(a.Name+" "+a.Role+" "+a.Bio, []string{q}) { list = append(list, a) } }
    r.s.mu.RUnlock()
    data, meta := paginate(list, page, limit)
    return data, meta, nil
}

func (r *AuthorsRepo) Get(ctx context.Context, id string) (domain.Author, error) {
    r.s.mu.RLock(); defer r.s.mu.RUnlock()
    for _, a := range r.s.authors { if a.ID == id { return a, nil } }
    return domain.Author{}, domain.AppError{Code: "NOT_FOUND", Message: "author not found"}
}

func (r *AuthorsRepo) Create(ctx context.Context, a domain.Author) (domain.Author, error) {
    r.s.mu.Lock(); a.ID = genID(); r.s.authors = append(r.s.authors, a); r.s.mu.Unlock(); return a, nil
}

func (r *AuthorsRepo) Update(ctx context.Context, id string, a domain.Author) (domain.Author, error) {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.authors { if r.s.authors[i].ID == id { a.ID = id; r.s.authors[i] = a; return a, nil } }
    return domain.Author{}, domain.AppError{Code: "NOT_FOUND", Message: "author not found"}
}

func (r *AuthorsRepo) Delete(ctx context.Context, id string) error {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.authors { if r.s.authors[i].ID == id { r.s.authors = append(r.s.authors[:i], r.s.authors[i+1:]...); return nil } }
    return domain.AppError{Code: "NOT_FOUND", Message: "author not found"}
}
