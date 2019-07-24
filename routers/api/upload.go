package api

import (
	"github.com/gin-gonic/gin"
	"go-web-scaffold/pkgs/app"
	"go-web-scaffold/pkgs/e"
	"go-web-scaffold/pkgs/logging"
	"go-web-scaffold/pkgs/upload"
	"net/http"
	"path/filepath"
)

/*
上传文件
*/

func UploadImage(c *gin.Context) {

	// init response
	appG := app.Gin{c}
	code := e.SUCCESS
	data := make(map[string]string)

	// check1: 必须为图片
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Logs.Warn("Request from file err: ", err)
		code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
		appG.Response(http.StatusOK, code, data)
		return
	}

	// check2: 图片合法
	if image == nil {
		code = e.ERROR_UPLOAD_IMAGE_ISEMPTH
	} else {
		// 生成混淆文件名
		imageName := upload.GetMixFileName(image.Filename)
		// 生成upload image 路径
		fullPath := upload.GetImageFullPath()
		// 相对 image save 路径
		savePath := upload.GetImagePath()

		src := filepath.Join(fullPath, imageName)

		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) || !upload.CheckImageContentType(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckFile(fullPath)
			if err != nil {
				logging.Logs.Warn("%v", err)
				code = e.ERROR_UPLOAD_IMAGE_SAVEPATH_ISEMPTH
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Logs.Warn("Fail to save image: ", err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = filepath.Join(savePath, imageName)
			}
		}
	}

	appG.Response(http.StatusOK, code, data)
}
