/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Mozilla Public License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.mozilla.org/en-US/MPL/

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package manifest : Provides manipulation on an Android manifest file to retrieve or change some information in it.
package manifest

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Shosta/androSecTest/attacks"
	"github.com/Shosta/androSecTest/logging"
)

// Manifest :
type Manifest struct {
	Application Application `xml:"application"`
}

// Application :
type Application struct {
	Icon        string `xml:"icon,attr"`
	AllowBackup string `xml:"allowBackup,attr"`
}

func icon(pkgname string) string {
	// Open our xmlFile
	attacks.DisassemblePackageDirPath(pkgname)
	xmlFile, err := os.Open("/home/shosta/android/security/com.orange.orangeetmoi/attacks/disassemblePackage/AndroidManifest.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		logging.PrintlnError(err)
	}
	logging.PrintlnDebug("Successfully Opened AndroidManifest.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	manifest := Manifest{}
	err = xml.Unmarshal(byteValue, &manifest)
	if err != nil {
		fmt.Println(err)
	}
	logging.PrintlnDebug("Icon : " + manifest.Application.Icon)

	return manifest.Application.Icon
}

// IconName : Parse the Android application manifest file in order to retrieve the icon's name in order to later be able to work on it.
func IconName(pkgname string) string {
	// Split the Android Manifest's "icon" attribute that contains the "folder/icon name" and return only the icon name.
	return strings.Split(icon(pkgname), "/")[1]
}
