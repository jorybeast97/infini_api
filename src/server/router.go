package server

import (
    "infini_api/src/config"
    "infini_api/src/middleware"
    apiauth "infini_api/src/api/auth"
    apiauthors "infini_api/src/api/authors"
    apiposts "infini_api/src/api/posts"
    "github.com/gin-gonic/gin"
)

func BuildRouter(cfg config.Config, s Services) *gin.Engine {
    r := gin.Default()
    r.Use(middleware.CORS())
    r.Static("/uploads", cfg.UploadDir)

    r.POST("/api/auth/login", apiauth.LoginHandler(s.Auth))
    r.GET("/api/authors", apiauthors.ListHandler(s.Authors))
    r.GET("/api/posts", apiposts.ListHandler(s.Posts))

    return r
}
