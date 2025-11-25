package posts

import (
    "infini_api/src/api/contract"
    "infini_api/src/domain"
    postsservice "infini_api/src/service/posts"
    "github.com/gin-gonic/gin"
)

func SaveHandler(s *postsservice.PostsServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req contract.PostsSaveRequest
        if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        var p domain.BlogPost
        var err error
        if req.Post.ID == "" || req.Post.ID == "new" {
            p, err = s.Create(c.Request.Context(), req.Post)
        } else {
            p, err = s.Update(c.Request.Context(), req.Post.ID, req.Post)
        }
        if err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        c.JSON(200, contract.PostsSaveResponse{Data: p})
    }
}
