package auth

import (
    "infini_api/src/api/contract"
    authservice "infini_api/src/service/auth"
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

func MeHandler(tokens authservice.TokenService, users *usersservice.UsersServiceImpl) gin.HandlerFunc {
	return func(c *gin.Context) {
		authz := c.GetHeader("Authorization")
		if len(authz) < 8 || authz[:7] != "Bearer " {
			c.JSON(401, gin.H{"error": "missing bearer token"})
			return
		}
		tokenStr := authz[7:]
		id, _, _, err := tokens.Parse(tokenStr)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			return
		}
        u, err := users.Get(c.Request.Context(), id)
        if err != nil { c.JSON(404, contract.EmptyResponse{}); return }
        c.JSON(200, contract.MeResponse{ User: contract.AuthUserBrief{ ID: u.ID, UserName: u.UserName, NickName: u.NickName, Role: u.Role, Avatar: u.Avatar } })
	}
}
