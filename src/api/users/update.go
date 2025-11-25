package users

import (
    "infini_api/src/domain"
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

type usersUpdateRequest struct {
    ID       string  `json:"id"`
    NickName *string `json:"nickName"`
    Password *string `json:"password"`
    Avatar   *string `json:"avatar"`
    Bio      *string `json:"bio"`
    Role     *string `json:"role"`
    Status   *string `json:"status"`
}

func UpdateHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var body usersUpdateRequest
        if err := c.ShouldBindJSON(&body); err != nil || body.ID=="" { c.JSON(400, gin.H{}); return }
        u := domain.User{}
        if body.NickName != nil { u.NickName = *body.NickName }
        if body.Avatar != nil { u.Avatar = *body.Avatar }
        if body.Bio != nil { u.Bio = *body.Bio }
        if body.Role != nil { u.Role = *body.Role }
        if body.Status != nil { u.Status = *body.Status }
        updated, err := s.Update(c.Request.Context(), body.ID, u, body.Password)
        if err != nil { c.JSON(400, gin.H{}); return }
        c.JSON(200, updated)
    }
}
