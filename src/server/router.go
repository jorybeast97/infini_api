package server

import (
	apiauth "infini_api/src/api/auth"
	apiauthors "infini_api/src/api/authors"
	apihello "infini_api/src/api/hello"
	apiposts "infini_api/src/api/posts"
	apiusers "infini_api/src/api/users"
	"infini_api/src/config"
	"infini_api/src/middleware"

	"github.com/gin-gonic/gin"
)

func BuildRouter(cfg config.Config, s Services) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Static("/uploads", cfg.UploadDir)

    r.POST("/api/auth/login", apiauth.LoginHandler(s.Auth))
    r.POST("/api/auth/me", apiauth.MeHandler(s.Tokens, s.Users))
    r.POST("/api/auth/register", apiauth.RegisterHandler(s.Users))

    r.POST("/api/authors/list", apiauthors.ListHandler(s.Authors))

    r.POST("/api/posts/list", apiposts.ListHandler(s.Posts))
    r.POST("/api/posts/get", apiposts.GetHandler(s.Posts))
    r.POST("/api/posts/save", apiposts.SaveHandler(s.Posts))
    r.POST("/api/posts/delete", apiposts.DeleteHandler(s.Posts))

    r.POST("/api/hello", apihello.HelloHandler())

    r.POST("/api/users/list", apiusers.ListHandler(s.Users))
    r.POST("/api/users/get", apiusers.GetHandler(s.Users))
    r.POST("/api/users/create", apiusers.CreateHandler(s.Users))
    r.POST("/api/users/update", apiusers.UpdateHandler(s.Users))
    r.POST("/api/users/delete", apiusers.DeleteHandler(s.Users))

	return r
}
