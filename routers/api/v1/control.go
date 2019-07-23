package v1

import "github.com/gin-gonic/gin"

// 业务逻辑层 API 实现

func Query(c *gin.Context) {
	c.JSON(200, gin.H{"message": "test"})
}



