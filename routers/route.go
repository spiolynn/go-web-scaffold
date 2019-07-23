package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-scaffold/pkgs/setting"
	"go-web-scaffold/routers/api/v1"
	"time"
)

// 路由控制层

// 路由注册
func InitRouter() *gin.Engine {

	r := gin.New()

	// 日志中间件
	//r.Use(gin.Logger())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())


	gin.SetMode(setting.G_cfg_yaml.Server.Runmode)
	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/query", v1.Query)
	}


	return r
}

