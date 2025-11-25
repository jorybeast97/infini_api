package posts

import (
    "infini_api/src/api/contract"
    postsservice "infini_api/src/service/posts"
    "github.com/gin-gonic/gin"
)

func ListHandler(s *postsservice.PostsServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req contract.PostsListRequest
        if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        if req.Page == 0 { req.Page = 1 }
        if req.Limit == 0 { req.Limit = 20 }
        data, meta, err := s.List(c.Request.Context(), req.Q, req.Status, req.HasLocation, req.Sort, req.Page, req.Limit)
        if err != nil { c.JSON(400, contract.EmptyResponse{}); return }
        c.JSON(200, contract.PostsListResponse{Data: data, Meta: meta})
    }
}
