package auth

import (
    "infini_api/src/api/contract"
    authservice "infini_api/src/service/auth"
    "github.com/gin-gonic/gin"
)

func LoginHandler(s *authservice.AuthServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req contract.LoginRequest
        if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        token, u, err := s.Login(c.Request.Context(), req.Username, req.Password)
        if err != nil { c.JSON(401, contract.EmptyResponse{}); return }
        resp := contract.LoginResponse{ Token: token, User: contract.AuthUserBrief{ ID: u.ID, UserName: u.UserName, NickName: u.NickName, Role: u.Role, Avatar: u.Avatar } }
        c.JSON(200, resp)
    }
}
