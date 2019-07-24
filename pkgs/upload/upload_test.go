package upload

import (
	"fmt"
	"go-web-frame/pkgs/setting"
	"os"
	"testing"
)

func beforeTest() {
	setting.Setup("../../confs/app.ini")
}

func TestGetFileName(t *testing.T) {
	beforeTest()
	filename := GetFileName("123.pdf")
	fmt.Printf("%s", filename)
}

func TestGetFilePath(t *testing.T) {
	beforeTest()
	fmt.Printf("%s\n", GetFilePath())
}

func TestCheckFile(t *testing.T) {
	beforeTest()
	CheckFile("123.txt")
}

func TestGetImageFullPath(t *testing.T) {
	beforeTest()
	fmt.Printf("%s\n", GetFileFullPath())
}

func TestCheckImageExt(t *testing.T) {
	beforeTest()
	fmt.Print(CheckImageExt("123.png"))
}

func TestGetFileContentType(t *testing.T) {
	f, err := os.Open("C:\\GOLib\\src\\go-web-frame\\runtime\\temp\\upload\\images\\ee0dc655da02e88747edcd64313856a3.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Get the content
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content Type: " + contentType)
}

func TestCheckImageContentType(t *testing.T) {
	beforeTest()

	f, err := os.Open("C:\\GOLib\\src\\go-web-frame\\runtime\\temp\\upload\\images\\ee0dc655da02e88747edcd64313856a3.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	isok := CheckImageContentType(f)

	fmt.Print(isok)

}
