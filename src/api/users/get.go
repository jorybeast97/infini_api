package users

import (
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

type usersGetRequest struct { ID string `json:"id"` }

func GetHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req usersGetRequest
        if err := c.ShouldBindJSON(&req); err != nil || req.ID=="" { c.JSON(400, gin.H{}); return }
        u, err := s.Get(c.Request.Context(), req.ID)
        if err != nil { c.JSON(404, gin.H{}); return }
        c.JSON(200, u)
    }
}
