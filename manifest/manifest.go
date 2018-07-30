package manifest

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/shosta/androSecTest/attacks"
	"github.com/shosta/androSecTest/logging"
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
		fmt.Println(err)
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
