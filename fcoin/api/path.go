package api

import (
	"path/filepath"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func GetPath(endPoint string, methodName string) (ret string) {
	fcoinDir, _ := filepath.Abs("../")
	filePath := fcoinDir + "/config/path.yaml"
	fileSource := file.NewSource(file.WithPath(filePath))
	config.Load(fileSource)
	ret = config.Get("fcoin", "endpoints", endPoint, methodName).String("")
	return
}
