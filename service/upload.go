package service

import (
	"ginmall/conf"
	"io"

	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLoacalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}

	avatarPath := basePath + userName + ".jpg" // todo 提取后缀
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func UploadProductToLoacalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}

	productPath := basePath + productName + ".jpg" // todo 提取后缀
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// 判断路径存在与否
func DirExistOrNot(fileAddr string) bool {
	if s, err := os.Stat(fileAddr); err != nil {
		return false
	} else {
		return s.IsDir()
	}
}

// 路径不存在 创建文件夹
func CreateDir(dirname string) bool {
	err := os.MkdirAll(dirname, 0755)
	return err == nil
}
