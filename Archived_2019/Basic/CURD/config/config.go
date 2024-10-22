package config

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetFileName() string {

	filename := []string{"config.json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
