package contract

import "infini_api/src/domain"

// Auth
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
type LoginResponse struct {
    Token string        `json:"token"`
    User  AuthUserBrief `json:"user"`
}
type AuthUserBrief struct {
    ID       string `json:"id"`
    UserName string `json:"userName"`
    NickName string `json:"nickName"`
    Role     string `json:"role"`
    Avatar   string `json:"avatar"`
}
type RegisterRequest struct {
    UserName string `json:"userName"`
    NickName string `json:"nickName,omitempty"`
    Password string `json:"password"`
    Avatar   string `json:"avatar,omitempty"`
}
type RegisterResponse struct { User AuthUserBrief `json:"user"` }
type MeResponse struct { User AuthUserBrief `json:"user"` }

// Posts
type PostsListRequest struct {
    Q           string `json:"q,omitempty"`
    Status      string `json:"status,omitempty"`
    HasLocation *bool  `json:"hasLocation,omitempty"`
    Sort        string `json:"sort,omitempty"`
    Page        int    `json:"page"`
    Limit       int    `json:"limit"`
}
type PostsListResponse struct {
    Data []domain.BlogPost `json:"data"`
    Meta domain.Meta       `json:"meta"`
}
type PostsGetRequest struct { ID string `json:"id"` }
type PostsGetResponse struct { Data domain.BlogPost `json:"data"` }
type PostsSaveRequest struct { Post domain.BlogPost `json:"post"` }
type PostsSaveResponse struct { Data domain.BlogPost `json:"data"` }
type PostsDeleteRequest struct { ID string `json:"id"` }
type EmptyResponse struct{ }

