package e

/*
错误码
设计：
统一格式：A-BB-CCC
A:错误级别，如X代表系统级错误，Y代表服务级错误；
B:项目或模块名称，一般公司不会超过99个项目；
C:具体错误编号，自增即可，一个项目999种错误应该够用；
举例:
1: 系统级别
2: 业务级别

01: 模块1
02: 模块2

001: XX 错误

001001
*/

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	// 上传模块
	ERROR_UPLOAD_SAVE_IMAGE_FAIL        = 101001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL       = 101002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT     = 101003
	ERROR_UPLOAD_IMAGE_ISEMPTH          = 101004
	ERROR_UPLOAD_IMAGE_SAVEPATH_ISEMPTH = 101005
)
