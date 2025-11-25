package memory

import (
    "math"
    "strings"
    "sync"
    "infini_api/src/domain"
    "crypto/rand"
)

type MemoryStore struct {
    mu      sync.RWMutex
    apps    []domain.AppProject
    posts   []domain.BlogPost
    photos  []domain.Photo
    authors []domain.Author
    users   []domain.User
}

func NewMemoryStore() *MemoryStore { return &MemoryStore{} }

func Seed(s *MemoryStore) {
    s.mu.Lock()
    if len(s.apps) == 0 {
        s.apps = []domain.AppProject{
            {ID: "1", Name: "ZenTask", Description: "A minimalist productivity app focused on flow state.", Icon: "CheckCircle", URL: "#", Tags: []string{"Productivity", "React Native", "iOS"}},
            {ID: "2", Name: "CryptoPulse", Description: "Real-time cryptocurrency market visualization tool.", Icon: "TrendingUp", URL: "#", Tags: []string{"Finance", "D3.js", "Web3"}},
            {ID: "3", Name: "EchoNotes", Description: "AI-powered voice memo organizer and summarizer.", Icon: "Mic", URL: "#", Tags: []string{"AI", "Audio", "Utility"}},
        }
    }
    if len(s.posts) == 0 {
        s.posts = []domain.BlogPost{
            {ID: "1", Title: "The Future of Declarative UI", Excerpt: "Why keeping your UI logic declarative makes scaling easier than you think.", Content: "Full content here...", Date: 1710460800, ReadTime: "5 min read", Location: &domain.Location{Lat: 37.7749, Lng: -122.4194, Name: "San Francisco, CA"}},
            {ID: "2", Title: "Nomad Life in Tokyo", Excerpt: "Spending a month working remotely from Shibuya.", Content: "Full content here...", Date: 1699574400, ReadTime: "8 min read", Location: &domain.Location{Lat: 35.6895, Lng: 139.6917, Name: "Tokyo, Japan"}},
            {ID: "3", Title: "Building performant lists in React", Excerpt: "Virtualization techniques deep dive.", Content: "Full content here...", Date: 1695340800, ReadTime: "10 min read", Location: nil},
        }
    }
    if len(s.photos) == 0 {
        s.photos = []domain.Photo{
            {ID: "1", URL: "https://picsum.photos/600/400?random=1", Caption: "Sunset over the Golden Gate", Date: 1710028800, Location: domain.Location{Lat: 37.8199, Lng: -122.4783, Name: "Golden Gate Bridge"}},
            {ID: "2", URL: "https://picsum.photos/600/800?random=2", Caption: "Neon streets of Shinjuku", Date: 1700006400, Location: domain.Location{Lat: 35.6909, Lng: 139.7005, Name: "Shinjuku"}},
            {ID: "3", URL: "https://picsum.photos/600/600?random=3", Caption: "Coffee shop vibes in Berlin", Date: 1691193600, Location: domain.Location{Lat: 52.5200, Lng: 13.4050, Name: "Berlin, Germany"}},
            {ID: "4", URL: "https://picsum.photos/600/400?random=4", Caption: "Hiking in the Swiss Alps", Date: 1689811200, Location: domain.Location{Lat: 46.8182, Lng: 8.2275, Name: "Swiss Alps"}},
        }
    }
    if len(s.authors) == 0 {
        gh := "#"; li := "#"; tw := "#"
        s.authors = []domain.Author{
            {ID: "1", Name: "Infini", Role: "Frontend Architect", Avatar: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=150&h=150&fit=crop", Bio: "Obsessed with pixel perfection and accessible UI systems. Building the digital future.", Social: domain.Social{Github: &gh, Twitter: &tw, Linkedin: &li}},
            {ID: "2", Name: "Aria", Role: "UX Researcher", Avatar: "https://images.unsplash.com/photo-1494790108377-be9c29b29330?w=150&h=150&fit=crop", Bio: "Decoding human behavior to build intuitive digital experiences.", Social: domain.Social{Linkedin: &li, Twitter: &tw}},
            {ID: "3", Name: "Cipher", Role: "Backend Engineer", Avatar: "https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?w=150&h=150&fit=crop", Bio: "Scalable infrastructure and secure data pipelines.", Social: domain.Social{Github: &gh}},
        }
    }
    if len(s.users) == 0 {
        s.users = []domain.User{}
    }
    s.mu.Unlock()
}

func paginate[T any](items []T, page, limit int) ([]T, domain.Meta) {
    if page < 1 { page = 1 }
    if limit < 1 { limit = 20 }
    if limit > 100 { limit = 100 }
    total := len(items)
    start := (page - 1) * limit
    if start > total { return []T{}, domain.Meta{Page: page, Limit: limit, Total: total} }
    end := int(math.Min(float64(start+limit), float64(total)))
    return items[start:end], domain.Meta{Page: page, Limit: limit, Total: total}
}

func containsAny(s string, qs []string) bool {
    sl := strings.ToLower(s)
    for _, q := range qs { if q != "" && strings.Contains(sl, strings.ToLower(q)) { return true } }
    return false
}

func genID() string {
    b := make([]byte, 8)
    _, err := rand.Read(b)
    if err != nil { return "id" }
    const hex = "0123456789abcdef"
    out := make([]byte, 16)
    for i, v := range b { out[i*2] = hex[v>>4]; out[i*2+1] = hex[v&0x0f] }
    return string(out)
}
