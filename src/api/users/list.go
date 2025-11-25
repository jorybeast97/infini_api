package users

import (
    "strconv"
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

func ListHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        q := c.Query("q")
        page, _ := strconv.Atoi(c.Query("page")); if page==0 { page=1 }
        limit, _ := strconv.Atoi(c.Query("limit")); if limit==0 { limit=20 }
        data, meta, err := s.List(c.Request.Context(), q, page, limit)
        if err != nil { c.JSON(400, gin.H{"error": err.Error()}); return }
        c.JSON(200, gin.H{"data": data, "meta": meta})
    }
}

