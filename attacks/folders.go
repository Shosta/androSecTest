package attacks

import (
	"os"

	"github.com/shosta/androSecTest/config"
	"github.com/shosta/androSecTest/logging"

	"github.com/shosta/androSecTest/variables"
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
