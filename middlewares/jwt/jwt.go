package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-web-scaffold/pkgs/e"
	"go-web-scaffold/pkgs/security"
	"net/http"
)

// JWT中间件
// 校驗token是否过期
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 函数
		var code int
		var data interface{}
		code = e.SUCCESS

		//token := c.Query("token")
		token := c.GetHeader("token")

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := security.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			// 异常返回
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
