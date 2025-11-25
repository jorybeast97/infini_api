package memory

import (
    "context"
    "sort"
    "strings"
    "infini_api/src/domain"
)

type PostsRepo struct{ s *MemoryStore }

func NewPostsRepo(s *MemoryStore) *PostsRepo { return &PostsRepo{s: s} }

func (r *PostsRepo) List(ctx context.Context, q, status string, hasLocation *bool, sortQ string, page, limit int) ([]domain.BlogPost, domain.Meta, error) {
    r.s.mu.RLock(); var list []domain.BlogPost
    for _, p := range r.s.posts {
        if status != "" && p.Status != status { continue }
        if hasLocation != nil { want := *hasLocation; if (p.Location != nil) != want { continue } }
        if q == "" || containsAny(p.Title+" "+p.Excerpt+" "+p.Content, []string{q}) { list = append(list, p) }
    }
    r.s.mu.RUnlock()
    if sortQ != "" {
        parts := strings.Split(sortQ, ",")
        for i := len(parts)-1; i>=0; i-- {
            fd := strings.Split(parts[i], ":"); field := fd[0]; dir := "asc"; if len(fd)>1 { dir = fd[1] }
            sort.SliceStable(list, func(i,j int) bool { var vi,vj string; switch field { case "date": vi,vj=list[i].Date,list[j].Date; case "title": vi,vj=list[i].Title,list[j].Title; default: vi,vj=list[i].ID,list[j].ID }; if dir=="desc" { return vi>vj }; return vi<vj })
        }
    }
    data, meta := paginate(list, page, limit)
    return data, meta, nil
}

func (r *PostsRepo) Get(ctx context.Context, id string) (domain.BlogPost, error) {
    r.s.mu.RLock(); defer r.s.mu.RUnlock()
    for _, p := range r.s.posts { if p.ID == id { return p, nil } }
    return domain.BlogPost{}, domain.AppError{Code: "NOT_FOUND", Message: "post not found"}
}

func (r *PostsRepo) Create(ctx context.Context, p domain.BlogPost) (domain.BlogPost, error) {
    if p.Status == "" { p.Status = "draft" }
    r.s.mu.Lock(); p.ID = genID(); r.s.posts = append(r.s.posts, p); r.s.mu.Unlock(); return p, nil
}

func (r *PostsRepo) Update(ctx context.Context, id string, p domain.BlogPost) (domain.BlogPost, error) {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.posts { if r.s.posts[i].ID == id { p.ID = id; r.s.posts[i] = p; return p, nil } }
    return domain.BlogPost{}, domain.AppError{Code: "NOT_FOUND", Message: "post not found"}
}

func (r *PostsRepo) Delete(ctx context.Context, id string) error {
    r.s.mu.Lock(); defer r.s.mu.Unlock()
    for i := range r.s.posts { if r.s.posts[i].ID == id { r.s.posts = append(r.s.posts[:i], r.s.posts[i+1:]...); return nil } }
    return domain.AppError{Code: "NOT_FOUND", Message: "post not found"}
}
