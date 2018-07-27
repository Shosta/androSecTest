package file

import (
	"path"
	"strings"
)

// IsPNG : Return true if the file extension is "png".
func IsPNG(filePath string) bool {
	if strings.ToUpper(path.Ext(filePath)) == ".PNG" {
		return true
	}

	return false
}

// IsJPG : Return true if the file extension is "jpg" or "jpeg".
func IsJPG(filePath string) bool {
	if strings.ToUpper(path.Ext(filePath)) == ".JPG" || strings.ToUpper(path.Ext(filePath)) == ".JPEG" {
		return true
	}

	return false
}
