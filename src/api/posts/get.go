package posts

import (
    "infini_api/src/api/contract"
    postsservice "infini_api/src/service/posts"
    "github.com/gin-gonic/gin"
)

func GetHandler(s *postsservice.PostsServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req contract.PostsGetRequest
        if err := c.ShouldBindJSON(&req); err != nil || req.ID == "" { c.JSON(400, contract.EmptyResponse{}); return }
        p, err := s.Get(c.Request.Context(), req.ID)
        if err != nil { c.JSON(404, contract.EmptyResponse{}); return }
        c.JSON(200, contract.PostsGetResponse{Data: p})
    }
}

