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

// Package androidpkg : Repackage a package after enabling some penetration testing features in it.
package androidpkg

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Shosta/androSecTest/settings"

	"github.com/Shosta/androSecTest/images"
	"github.com/Shosta/androSecTest/manifest"

	"github.com/Shosta/androSecTest/command/sed"

	folders "github.com/Shosta/androSecTest/attacks"
	"github.com/Shosta/androSecTest/command"
	"github.com/Shosta/androSecTest/command/adb"
	"github.com/Shosta/androSecTest/command/apktool"
	"github.com/Shosta/androSecTest/logging"
)

// Setup :
func Setup(pkgname string) {
	unzip(pkgname)
	disassemble(pkgname)
	mkdbg(pkgname)
	allowbackup(pkgname)
	addDbgBadgeOnAppIcon(pkgname)
	rebuild(pkgname)
	sign(pkgname)
	reinstall(pkgname)
}


// Unzip the package to the 'unzippedPackage' Folder
// cmd = "unzip " + attacksDir + variables.SourcePackageDir + "/" + pkgname + ".apk '*' -d " + unzipDir
func unzip(pkgname string) {
	sourceDirPath := folders.SourcePackageDirPath(pkgname)
	unzipDirPath := folders.UnzipDirPath(pkgname)
	logging.Println(logging.Green("Unzip package : ") + logging.Bold(pkgname) + " to " + logging.Bold(unzipDirPath))

	cmdName := "unzip"
	cmdArgs := []string{
		sourceDirPath + "/" + pkgname + ".apk",
		"-d",
		unzipDirPath,
	}
	command.Run(cmdName, cmdArgs)

	logging.Println(logging.Bold("Done"))
}

// Disassemble the package using the apktool that is installed on the system.
func disassemble(pkgname string) {
	logging.Println(logging.Green("Disassemble package : ") + logging.Bold(pkgname))
	logging.Println("Work in progress...")
	apktool.Disassemble(pkgname)

	logging.Println(logging.Bold("Done"))
}

func mkdbg(pkgname string) {
	logging.Println(logging.Green("Make package debuggable"))

	disassembledDirPath := folders.DisassemblePackageDirPath(pkgname)
	sed.Replace(disassembledDirPath+"/AndroidManifest.xml", "<application ", "<application android:debuggable=\"true\" ")

	logging.Println(logging.Bold("Done"))
}

// TODO : Should verify that the AllowBackup is not already available in the AppManifest.xml file.
func allowbackup(pkgname string) {
	logging.Println(logging.Green("Allow backup on package"))

	disassembledDir := folders.DisassemblePackageDirPath(pkgname)

	sed.Replace(disassembledDir+"/AndroidManifest.xml", "android:allowBackup=\"true\"", "")
	sed.Replace(disassembledDir+"/AndroidManifest.xml", "android:allowBackup=\"false\"", "")
	logging.PrintlnDebug("Remove the android:allowBackup=\"false\" if it is in the AndroidManifest.xml file.")

	sed.Replace(disassembledDir+"/AndroidManifest.xml", "<application ", "<application android:allowBackup=\"true\" ")
	logging.PrintlnDebug("Add the android:allowBackup=\"true\" to the AndroidManifest.xml file.")

	logging.Println(logging.Bold("Done"))
}

// On boucle sur tous les dossiers/fichiers pour trouver les fichiers qui contiennent iconName.
// Cela permet d'envoyer dans une channel les chemins absolus vers les icones.
func searchForIconPaths(dirPath string, iconName string, wg *sync.WaitGroup, dirlistchan chan string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			searchForIconPaths(dirPath+"/"+file.Name(), iconName, wg, dirlistchan)
		}
		// fmt.Println(file.Name())
		if strings.Contains(file.Name(), iconName) && strings.HasSuffix(strings.ToUpper(file.Name()), ".PNG") {
			dirlistchan <- dirPath + "/" + file.Name()
		}
	}
}

// Adding a watermark on the app icons based on the package name.
// It is using  Go routines and channels to do it asynchronously (even if it brings absolutely nothing in terms of performance but it was a great way for me to learn Go routines).
func addDbgBadgeOnAppIcon(pkgname string) {
	// TODO : Le tester sur de nombreuses applications.
	disassembleDirPath := folders.DisassemblePackageDirPath(pkgname)

	iconName := manifest.IconName(pkgname)

	resDir := disassembleDirPath + "/res"

	wg := new(sync.WaitGroup)
	dirlistchan := make(chan string, 1000)
	wg.Add(1)
	go func() {
		defer wg.Done()
		searchForIconPaths(resDir, iconName, wg, dirlistchan)
	}()

	go func() {
		wg.Wait()
		close(dirlistchan)
	}()

	for icon := range dirlistchan {
		// TODO : Must find an icon set that has a default icon that is large enough to handle any icon.
		dpi := "_xxhdpi" // Default to the larger icon.
		if strings.HasSuffix(filepath.Dir(icon), "xxxhdpi") {
			dpi = "_xxhdpi"
		} else if strings.HasSuffix(filepath.Dir(icon), "xxhdpi") {
			dpi = "_xxhdpi"
		} else if strings.HasSuffix(filepath.Dir(icon), "xhdpi") {
			dpi = "_xhdpi"
		} else if strings.HasSuffix(filepath.Dir(icon), "hdpi") {
			dpi = "_hdpi"
		}

		images.Watermark("./.res/watermark/dbg/unlock"+dpi+".png", icon)
	}

	wg.Wait()
}

func rebuild(pkgname string) {
	logging.Println(logging.Green("Rebuild package : ") + logging.Bold(pkgname) + "\nWork in progress...")

	apktool.Build(pkgname)

	logging.Println(logging.Bold("Done"))
}

//signapk /tmp/Attacks/DebuggablePackage/" + package_name + ".b.apk"
func sign(pkgname string) {
	pkgFilePath := folders.DebugPkgDirPath(pkgname) + "/" + pkgname
	logging.PrintlnDebug("Package to sign : " + pkgFilePath + ".b.apk")
	//var cmd string = "java -jar /home/shosta/ShostaSyncBox/Developpement/HackingTools/SignApkUtils/sign.jar " + pkgLoc + "b.apk"

	cmdArgs := []string{
		"-jar",
		settings.SignApk(),
		"/home/Developpement/HackingTools/SignApkUtils/testkey.x509.pem",
		"/home/Developpement/HackingTools/SignApkUtils/testkey.pk8",
		pkgFilePath + ".b.apk",
		pkgFilePath + ".b.s.apk",
	}
	command.Run("java", cmdArgs)
	// command.RunAlias("/bin/bash", "-c", "signapk " + pkgLoc + ".b.apk")
}

//adb uninstall package_name
//adb install "/tmp/Attacks/DebuggablePackage/" + package_name + ".b.s.apk"
func reinstall(pkgname string) {
	pkgFilePath := folders.DebugPkgDirPath(pkgname) + "/" + pkgname + ".b.s.apk"
	if _, err := os.Stat(pkgFilePath); os.IsNotExist(err) {
		logging.PrintlnError("Debuggable package does not exist. Please review the repackaging errors and retry the process before we can install it on the device.")

		return
	}

	adb.Uninstall(pkgname)
	adb.Install(pkgFilePath)
}
