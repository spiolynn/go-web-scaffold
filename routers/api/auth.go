package api

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-web-scaffold/models"
	"go-web-scaffold/pkgs/app"
	"go-web-scaffold/pkgs/e"
	"go-web-scaffold/pkgs/logging"
	"go-web-scaffold/pkgs/security"
	"net/http"
)

//用户鉴权

type auth struct {
	Username string `json:"Required; MaxSize(50)" valid:"Required; MaxSize(50)"`
	Password string `json:"Required; MaxSize(50)" valid:"Required; MaxSize(50)"`
}

//校验用户
func CheckAdmin(c *gin.Context) {

	appG := app.Gin{c}
	data := make(map[string]string)
	code := e.INVALID_PARAMS

	// 数据获取 by param
	username := c.Query("username")
	password := c.Query("password")

	// 接口数据合法校验
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if ok {
		// 业务逻辑
		isExist, _ := models.CheckAdmin(username, password)
		if isExist {
			token, err := security.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
				logging.Logs.Warn("Fail to Generate Token: ", err)
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH_USER_ISNOTEXIST
		}
	} else {
		// 验证失败
		validErrorMsg := GetErrorInfo(valid.Errors)
		data["msg"] = validErrorMsg
	}

	appG.Response(http.StatusOK, code, data)
}

// 解析错误信息
func GetErrorInfo(errors []*validation.Error) string {
	var valid_strs []string
	for _, err := range errors {
		logging.Logs.Info(err.Key, err.Message)
		valid_strs = append(valid_strs, fmt.Sprintf("|%v %v|", err.Key, err.Message))
	}
	strs := fmt.Sprintf("%v", valid_strs)
	return strs
}
