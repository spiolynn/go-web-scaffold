package upload

import (
	"fmt"
	"go-web-scaffold/pkgs/file"
	"go-web-scaffold/pkgs/logging"
	"go-web-scaffold/pkgs/setting"
	"go-web-scaffold/pkgs/util"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	// 文件后缀
	imageAllowExts []string = setting.G_cfg_yaml.App.Imageallowexts

	// 文件保存根路径
	runtimebasepath string = setting.G_cfg_yaml.App.Runtimebasepath

	// 图片文件保存相对路径
	imagesavepath string = setting.G_cfg_yaml.App.Imagesavepath

	imagemaxsize int = setting.G_cfg_yaml.App.Imagemaxsize

	imageallowextstype []string = setting.G_cfg_yaml.App.Imageallowextstype

	filepreurl string = setting.G_cfg_yaml.App.Filepreurl

	imageurl string = setting.G_cfg_yaml.App.Imageurl
)

// article file
//func GetFileFullUrl(name string) string {
//	return filepreurl + GetImageUrl() + name
//}

// 对文件名进行混淆
func GetMixFileName(name string) string {
	return GetFileName(name)
}

// 对文件名进行混淆
func GetFileName(name string) string {
	ext := file.GetExt(name)
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMD5(filename)
	return filename + ext
}

//func GetFilePath() string {
//	return setting.G_cfg_yaml.App.Filesavepath
//}
//
//
//func GetFileFullPath() string {
//	return setting.AConfigData.AApp.RuntimeRootPath + GetFilePath()
//}

// 检测文件后缀合法
//func CheckFileExt(filename string) bool {
//	ext := file.GetExt(filename)
//	for _, allowExt := range setting.AConfigData.AApp.FileAllowExts{
//		if strings.ToUpper(allowExt) == strings.ToUpper(ext){
//			return true
//		}
//	}
//	return false
//}

//func CheckFileSize(f multipart.File) bool {
//	size ,err := file.GetSize(f)
//	if err != nil {
//		logging.Warn("Check File with error: %v", err)
//		return false
//	}
//	return size <= setting.AConfigData.AApp.FileMaxSize
//}

// 检测文件夹是否存在、不存在则创建、运行根目录
func CheckFile(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}

// -------------------------------------------------------------------

// upload image pull path
func GetImageFullPath() string {
	return filepath.Join(runtimebasepath, GetImagePath())
}

// upload image 相对路径
func GetImagePath() string {
	return imagesavepath
}

// 图片后缀检查
func CheckImageExt(filename string) bool {
	ext := file.GetExt(filename)
	for _, allowExt := range imageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	logging.Logs.Warn("%v image ext error", filename)
	return false
}

// 图片大小检测
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		logging.Logs.Warn("image size error")
		return false
	}
	return size/1024/1024 <= imagemaxsize
}

// 检测文件类型
func CheckImageContentType(f multipart.File) bool {

	buffer := make([]byte, 512)

	// 文件读取游标恢复
	_, err1 := f.Seek(0, 0)
	if err1 != nil {
		return false
	}
	_, err := f.Read(buffer)
	if err != nil {
		return false
	} else {
		contentType := http.DetectContentType(buffer)

		for _, allowExt := range imageallowextstype {
			if strings.ToUpper(allowExt) == strings.ToUpper(contentType) {
				return true
			}
		}

		logging.Logs.Warn("image ContentType error ", contentType)
		return false
	}
}

// 获取文件头十六进制信息
func GetFileContentType(out *os.File) (string, error) {

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func GetImageFullUrl(name string) string {
	return filepreurl + GetImageUrl() + name
}

func GetImageUrl() string {
	return imageurl
}
