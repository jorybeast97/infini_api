package posts

import (
    "infini_api/src/api/contract"
    postsservice "infini_api/src/service/posts"
    "github.com/gin-gonic/gin"
)

func DeleteHandler(s *postsservice.PostsServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req contract.PostsDeleteRequest
        if err := c.ShouldBindJSON(&req); err != nil || req.ID == "" { c.JSON(400, contract.EmptyResponse{}); return }
        if err := s.Delete(c.Request.Context(), req.ID); err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        c.JSON(200, contract.EmptyResponse{})
    }
}

