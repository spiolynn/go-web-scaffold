package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-web-scaffold/pkgs/app"
	"go-web-scaffold/pkgs/e"
	"go-web-scaffold/pkgs/logging"
	"go-web-scaffold/pkgs/util"
	"go-web-scaffold/routers/api"
	"go-web-scaffold/service"
	"net/http"
)

// 业务逻辑层 API 实现
func Query(c *gin.Context) {
	c.JSON(200, gin.H{"message": "test"})
}

/*
@feature : 获取admins表列表信息
@url     : Get /Admins/
@param   : json
{
	username: ***
	password: ***
}
*/

func QueryAdminList(c *gin.Context) {

	logging.Logs.Infof("QueryAdminList in : %v", c.Request.RequestURI)
	// 1 返回数据准备
	appG := app.Gin{c}
	var data = make(map[string]interface{})

	// 2 json 数据绑定
	var _DQueryAdminList Admins
	if err := c.ShouldBindJSON(&_DQueryAdminList); err != nil {
		logging.Logs.Errorf(" Bind fail  %v ", err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
		return
	}
	logging.Logs.Infof("in: %v", _DQueryAdminList)

	// 3 数据校验
	valid := validation.Validation{}
	ok, _ := valid.Valid(&_DQueryAdminList)

	if !ok {
		// 验证失败
		validErrorMsg := api.GetErrorInfo(valid.Errors)
		data["msg"] = validErrorMsg
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
	} else {
		// 验证通过
		_mQueryAdminList := util.Struct_to_Map(_DQueryAdminList)
		jsonStr, err := service.QueryAdminList(_mQueryAdminList)
		if err != nil {
			appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, data)
		} else {
			data["status"] = "02"
			data["adminlist"] = jsonStr
			logging.Logs.Infof("out: %v", data)
			appG.Response(http.StatusOK, e.SUCCESS, data)
		}
	}

}
