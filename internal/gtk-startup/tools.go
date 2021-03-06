package gtkStartup

import (
	"fmt"
	"os"
	"path"
)

// ErrorCheckWithPanic : panics on error
func ErrorCheckWithPanic(err error, message string) {
	if err != nil {
		panic(err.Error() + " : " + message)
	}
}

// ErrorCheckWithPanic : panics on error
func ErrorCheckWithoutPanic(err error, message string) {
	if err != nil {
		fmt.Println(err.Error() + " : " + message)
	}
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getResourcePath(fileName string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := path.Dir(exePath)

	gladePath := path.Join(exeDir, fileName)
	if fileExists(gladePath) {
		return gladePath, nil
	}
	gladePath = path.Join(exeDir, "assets", fileName)
	if fileExists(gladePath) {
		return gladePath, nil
	}
	gladePath = path.Join(exeDir, "../assets", fileName)
	if fileExists(gladePath) {
		return gladePath, nil
	}
	return gladePath, nil
}
