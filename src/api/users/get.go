package users

import (
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

func GetHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        u, err := s.Get(c.Request.Context(), id)
        if err != nil { c.JSON(404, gin.H{"error": err.Error()}); return }
        c.JSON(200, u)
    }
}

