package auth

import (
    authservice "infini_api/src/service/auth"
    "github.com/gin-gonic/gin"
)

func LoginHandler(s *authservice.AuthServiceImpl) gin.HandlerFunc {
    return func(c *gin.Context) {
        var body struct{ Username, Password string }
        if err := c.ShouldBindJSON(&body); err != nil { c.JSON(400, gin.H{"error":"invalid json"}); return }
        token, name, role, err := s.Login(c.Request.Context(), body.Username, body.Password)
        if err != nil { c.JSON(401, gin.H{"error": err.Error()}); return }
        c.JSON(200, gin.H{"token": token, "user": gin.H{"id":"1","name":name,"role":role}})
    }
}
