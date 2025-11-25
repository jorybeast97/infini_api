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
	r.GET("/api/authors", apiauthors.ListHandler(s.Authors))
	r.GET("/api/posts", apiposts.ListHandler(s.Posts))
	r.GET("/api/hello", apihello.HelloHandler())
	r.GET("/api/users", apiusers.ListHandler(s.Users))
	r.GET("/api/users/:id", apiusers.GetHandler(s.Users))
	r.POST("/api/users", apiusers.CreateHandler(s.Users))
	r.PUT("/api/users/:id", apiusers.UpdateHandler(s.Users))
	r.DELETE("/api/users/:id", apiusers.DeleteHandler(s.Users))

	return r
}
