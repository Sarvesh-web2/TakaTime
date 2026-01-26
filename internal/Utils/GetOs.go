package utils

import "runtime"

func GetOS() string {
	osName := runtime.GOOS
	return osName
}
