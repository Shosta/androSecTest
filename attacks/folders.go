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

// Package attacks : creates all the folders that are going to store the application and the attacks results.
package attacks

import (
	"os"

	"github.com/Shosta/androSecTest/config"
	"github.com/Shosta/androSecTest/logging"

	"github.com/Shosta/androSecTest/variables"
)

// The folder to start working on a package is : ~/android/security/'packagename'/attacks/...

func createRootFolder() {
	os.MkdirAll(config.SecurityAssessmentRootDir, os.ModePerm)
}

// CreateAttacksFolder : Create the attack folder
func CreateAttacksFolder(pkgname string) {
	createRootFolder()

	var pkgdir = config.SecurityAssessmentRootDir + "/" + pkgname
	os.MkdirAll(pkgdir, os.ModePerm)

	os.MkdirAll(pkgdir+variables.AttacksDir+variables.SourcePackageDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.SourcePackageDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.UnzippedPackageDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.UnzippedPackageDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.DisassemblePackageDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.DisassemblePackageDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.DecompiledPackageDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.DecompiledPackageDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.DebuggablePackageDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.DebuggablePackageDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.InsecureLoggingDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.InsecureLoggingDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.InsecureStorageDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.InsecureStorageDir)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.InsecureBackupDir, os.ModePerm)
	logging.PrintlnDebug("Done : " + variables.InsecureBackupDir)
}

// InsecStorageDirPath : Return the Insecure Storage folder path.
func InsecStorageDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.InsecureStorageDir
}

// InsecLoggingDirPath : Return the Insecure Logging folder path.
func InsecLoggingDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.InsecureLoggingDir
}

// UnzipDirPath : Return the folder path where we store the "unzip" command result files.
func UnzipDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.UnzippedPackageDir
}

// SourcePackageDirPath : Return the folder that contains the package we pulled initially from the device.
func SourcePackageDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.SourcePackageDir
}

// DisassemblePackageDirPath : Return the folder path where we store the "apktool -d" command result files (.smali).
func DisassemblePackageDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.DisassemblePackageDir
}

// DecompiledPackageDirPath : Return the folder path where we store the "jadx -d" command result files (.java).
func DecompiledPackageDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.DecompiledPackageDir
}

// DebugPkgDirPath : Return the folder path where we store the "apktool b" command result files (.b.apk).
func DebugPkgDirPath(pkgname string) string {
	return config.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.DebuggablePackageDir
}

func createLeakageDir(pkgname string) {
	os.MkdirAll(DecompiledPackageDirPath(pkgname)+"/leakages", os.ModePerm)
}
