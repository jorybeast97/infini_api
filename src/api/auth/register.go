package auth

import (
    "time"
    "infini_api/src/api/contract"
    "infini_api/src/domain"
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

func RegisterHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req contract.RegisterRequest
        if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        if req.UserName == "" || req.Password == "" { c.JSON(400, contract.EmptyResponse{}); return }
        u := domain.User{UserName: req.UserName, NickName: req.NickName, Avatar: req.Avatar, Role: "editor", Status: "active", CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}
        created, err := s.Create(c.Request.Context(), u, req.Password)
        if err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        created.PasswordHash = ""
        c.JSON(201, contract.RegisterResponse{ User: contract.AuthUserBrief{ ID: created.ID, UserName: created.UserName, NickName: created.NickName, Role: created.Role, Avatar: created.Avatar } })
    }
}
