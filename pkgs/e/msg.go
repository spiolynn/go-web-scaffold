package e

// 错误消息
var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	//ERROR_EXIST_TAG:                 "已存在该标签名称",
	//ERROR_NOT_EXIST_TAG:             "该标签不存在",
	//ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",

	//ERROR_VALIDATION_FAIL:           "数据验证失败",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:        "图片保存失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT:     "图片文件类型不符合规范",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:       "输入文件必须图片文件",
	ERROR_UPLOAD_IMAGE_ISEMPTH:          "图片文件为空",
	ERROR_UPLOAD_IMAGE_SAVEPATH_ISEMPTH: "后台保存文件路径不存在",

	ERROR_AUTH:                 "用户名不存在 或密码错误",
	ERROR_AUTH_USER_ISNOTEXIST: "用户名不存在",
}

// Get Error Msg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
