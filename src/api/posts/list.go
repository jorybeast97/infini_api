package posts

import (
    "strconv"
    postsservice "infini_api/src/service/posts"
    "github.com/gin-gonic/gin"
)

func ListHandler(s *postsservice.PostsServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        q := c.Query("q")
        status := c.Query("status")
        sortQ := c.Query("sort")
        var hasLocation *bool
        if v := c.Query("hasLocation"); v != "" { b := v == "true"; hasLocation = &b }
        page, _ := strconv.Atoi(c.Query("page")); if page==0 { page=1 }
        limit, _ := strconv.Atoi(c.Query("limit")); if limit==0 { limit=20 }
        data, meta, err := s.List(c.Request.Context(), q, status, hasLocation, sortQ, page, limit)
        if err != nil { c.JSON(400, gin.H{"error": err.Error()}); return }
        c.JSON(200, gin.H{"data": data, "meta": meta})
    }
}
