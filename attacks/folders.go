package attacks

import (
	"os"

	"github.com/shosta/androSecTest/variables"
)

// The folder to start working on a package is : ~/android/security/'packagename'/attacks/...

func createRootFolder() {
	os.MkdirAll(variables.SecurityAssessmentRootDir, os.ModePerm)
}

// CreateAttacksFolder : Create the attack folder
func CreateAttacksFolder(pkgname string) {
	createRootFolder()

	var pkgdir = variables.SecurityAssessmentRootDir + "/" + pkgname
	os.MkdirAll(pkgdir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.SourcePackageDir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.UnzippedPackageDir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.DecodedPackageDir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.DebuggablePackageDir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.InsecureLoggingDir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.InsecureStorageDir, os.ModePerm)
	os.MkdirAll(pkgdir+variables.AttacksDir+variables.InsecureBackupDir, os.ModePerm)
}

// InsecStorageDirPath : Return the Insecure Storage folder path.
func InsecStorageDirPath(pkgname string) string {
	return variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.InsecureStorageDir
}

// DecodedPackageDirPath : Return the Decoded Package folder path.
func DecodedPackageDirPath(pkgname string) string {
	return variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.DecodedPackageDir
}

func createLeakageDir(pkgname string) {
	os.MkdirAll(DecodedPackageDirPath(pkgname)+"/leakages", os.ModePerm)
}
