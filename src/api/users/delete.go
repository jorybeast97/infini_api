package users

import (
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

func DeleteHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        if err := s.Delete(c.Request.Context(), id); err != nil { c.JSON(400, gin.H{"error": err.Error()}); return }
        c.Status(204)
    }
}

