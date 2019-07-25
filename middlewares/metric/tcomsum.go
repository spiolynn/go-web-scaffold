package metric

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-scaffold/pkgs/logging"
	"strconv"
	"time"
)

/*
请求拦截中间件
*/

// 记录时间
func Timing(c *gin.Context) {
	t1 := time.Now()
	RequestURL := c.Request.RequestURI
	logging.Logm.Info("request start [" + RequestURL + "]")
	logging.Logm.Info(fmt.Sprintf("request head:  %+v", c.Request.Header))

	// gin 只允许 get 一次body
	//body, _ := ioutil.ReadAll(c.Request.Body)
	//
	//if len(body) < 100 {
	//	logging.Logm.Info(fmt.Sprintf("request body:   %s", string(body)))
	//}

	c.Next()
	elapsed := time.Since(t1)
	t := elapsed * time.Nanosecond
	elapsed_ms := int64(t / time.Millisecond)
	logging.Logm.Info(fmt.Sprintf("request end   [" + RequestURL + "]"))
	logging.Logm.Info(fmt.Sprintf("[" + RequestURL + "] elapsed_ms " + strconv.FormatInt(elapsed_ms, 10) + "ms"))
}

//// 打印报文体
//func RequestOutPut(c *gin.Context) {
//	//setting.Logger.Info(fmt.Sprintf("request head:\n %+v",c.Request.Header))
//	//body,_ := ioutil.ReadAll(c.Request.Body)
//	//setting.Logger.Info(fmt.Sprintf("request body:\n %s",string(body)))
//	c.Next()
//}
