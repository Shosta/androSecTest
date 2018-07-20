package androidpkg

import (
	"github.com/shosta/androSecTest/attacks"
	"github.com/shosta/androSecTest/command/adb"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/variables"
)

// Savelocal : Allow the user to select the package he wants to save locally on its computer through a simple part of the package name.
// The package is saved in a folder on the $Home folder.
// It returns the package name
func Savelocal(pkgname string) {

	attacks.CreateAttacksFolder(pkgname)
	var pkgpath = adb.PkgPath(pkgname)
	pull(pkgname, pkgpath)
}

/**
Get the package path on the connected devices via adb regarding the package name, without the "apk" extension.

Params:
package_name The package name, without the "apk" extension.

Comment: Use the shell command : adb shell pm path 'pkgname'
*/
func path(pkgname string) string {
	var path = adb.PkgPath(pkgname)

	return path
}

// Pull a package on the connected devices via adb.
func pull(pkgname string, pkgpath string) {
	var destLocation = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.SourcePackageDir + "/" + pkgname + ".apk"
	logging.Println(logging.Green("Pull package from ") + logging.Bold(pkgpath))

	var out = adb.Pull(pkgpath, destLocation)

	logging.PrintlnVerbose(out)
	logging.Println(logging.Green("Package stored at ") + logging.Bold(destLocation))
}
