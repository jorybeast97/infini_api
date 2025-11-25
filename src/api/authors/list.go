package authors

import (
    authorsservice "infini_api/src/service/authors"
    "github.com/gin-gonic/gin"
)

type authorsListRequest struct { Q string `json:"q"`; Page int `json:"page"`; Limit int `json:"limit"` }
type authorsListResponse struct { Data interface{} `json:"data"`; Meta interface{} `json:"meta"` }

func ListHandler(s *authorsservice.AuthorsServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req authorsListRequest
        if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, authorsListResponse{}); return }
        if req.Page==0 { req.Page=1 }
        if req.Limit==0 { req.Limit=20 }
        data, meta, err := s.List(c.Request.Context(), req.Q, req.Page, req.Limit)
        if err != nil { c.JSON(400, authorsListResponse{}); return }
        c.JSON(200, authorsListResponse{Data: data, Meta: meta})
    }
}
