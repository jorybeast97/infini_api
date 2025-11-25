package memory

import (
    "context"
    "sort"
    "strings"
    "infini_api/src/domain"
)

type PhotosRepo struct{ s *MemoryStore }

func NewPhotosRepo(s *MemoryStore) *PhotosRepo { return &PhotosRepo{s: s} }

func (r *PhotosRepo) List(ctx context.Context, q string, sortQ string, page, limit int) ([]domain.Photo, domain.Meta, error) {
    r.s.mu.RLock(); var list []domain.Photo
    for _, p := range r.s.photos { if q == "" || containsAny(p.Caption+" "+p.URL+" "+p.Location.Name, []string{q}) { list = append(list, p) } }
    r.s.mu.RUnlock()
    if sortQ != "" {
        parts := strings.Split(sortQ, ",")
        for i := len(parts)-1; i>=0; i-- {
            fd := strings.Split(parts[i], ":"); field := fd[0]; dir := "asc"; if len(fd)>1 { dir = fd[1] }
            sort.SliceStable(list, func(i,j int) bool {
                switch field {
                case "date":
                    if dir=="desc" { return list[i].Date > list[j].Date }
                    return list[i].Date < list[j].Date
                case "caption":
                    if dir=="desc" { return list[i].Caption > list[j].Caption }
                    return list[i].Caption < list[j].Caption
                default:
                    if dir=="desc" { return list[i].ID > list[j].ID }
                    return list[i].ID < list[j].ID
                }
            })
        }
    }
    data, meta := paginate(list, page, limit)
    return data, meta, nil
}

func (r *PhotosRepo) Get(ctx context.Context, id string) (domain.Photo, error) {
    r.s.mu.RLock(); defer r.s.mu.RUnlock()
    for _, p := range r.s.photos { if p.ID == id { return p, nil } }
    return domain.Photo{}, domain.AppError{Code: "NOT_FOUND", Message: "photo not found"}
}

func (r *PhotosRepo) Create(ctx context.Context, p domain.Photo) (domain.Photo, error) {
    r.s.mu.Lock(); p.ID = genID(); r.s.photos = append(r.s.photos, p); r.s.mu.Unlock(); return p, nil
}

func (r *PhotosRepo) Update(ctx context.Context, id string, p domain.Photo) (domain.Photo, error) {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.photos { if r.s.photos[i].ID == id { p.ID = id; r.s.photos[i] = p; return p, nil } }
    return domain.Photo{}, domain.AppError{Code: "NOT_FOUND", Message: "photo not found"}
}

func (r *PhotosRepo) Delete(ctx context.Context, id string) error {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.photos { if r.s.photos[i].ID == id { r.s.photos = append(r.s.photos[:i], r.s.photos[i+1:]...); return nil } }
    return domain.AppError{Code: "NOT_FOUND", Message: "photo not found"}
}
