package users

import (
    "time"
    "infini_api/src/domain"
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

func CreateHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var body struct{
            UserName string `json:"userName"`
            NickName string `json:"nickName"`
            Password string `json:"password"`
            Avatar   string `json:"avatar"`
            Bio      string `json:"bio"`
            Role     string `json:"role"`
            Status   string `json:"status"`
        }
        if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":"invalid json"}); return }
        u := domain.User{UserName: body.UserName, NickName: body.NickName, Avatar: body.Avatar, Bio: body.Bio, Role: body.Role, Status: body.Status, CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
        created, err := s.Create(c.Request.Context(), u, body.Password)
        if err != nil { c.JSON(400, gin.H{"error": err.Error()}); return }
        c.JSON(201, created)
    }
}

