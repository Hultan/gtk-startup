package tools

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
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

// GetResourcePath : Gets the path to a resource file
func GetResourcePath(directory, file string) string {
	return path.Join(GetExecutablePath(), directory, file)
}

// GetExecutablePath : Returns the path of the executable
func GetExecutablePath() string {
	executable, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(executable)
}
