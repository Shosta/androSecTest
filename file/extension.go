/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Mozilla Public License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.mozilla.org/en-US/MPL/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package file : Check the extension of a file.
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
