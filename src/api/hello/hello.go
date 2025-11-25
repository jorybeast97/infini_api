package hello

import "github.com/gin-gonic/gin"

func HelloHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "hello world"})
    }
}

