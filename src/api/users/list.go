package users

import (
    usersservice "infini_api/src/service/users"
    "github.com/gin-gonic/gin"
)

type usersListRequest struct { Q string `json:"q"`; Page int `json:"page"`; Limit int `json:"limit"` }
type usersListResponse struct { Data interface{} `json:"data"`; Meta interface{} `json:"meta"` }

func ListHandler(s *usersservice.UsersServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req usersListRequest
        if err := c.ShouldBindJSON(&req); err != nil { c.JSON(400, usersListResponse{}); return }
        if req.Page==0 { req.Page=1 }; if req.Limit==0 { req.Limit=20 }
        data, meta, err := s.List(c.Request.Context(), req.Q, req.Page, req.Limit)
        if err != nil { c.JSON(400, usersListResponse{}); return }
        c.JSON(200, usersListResponse{Data: data, Meta: meta})
    }
}
