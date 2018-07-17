package androidpkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shosta/androSecTest/attacks"
	"github.com/shosta/androSecTest/command/adb"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/variables"
)

// Allow the user to select the package he wants to save locally on its computer through a simple part of the package name.
// The package is saved in a folder on the $Home folder.
// It returns the package name
func Savelocal(pkgnamepart string) string {

	var pkgname = packageName(pkgnamepart)

	attacks.CreateAttacksFolder(pkgname)
	var pkgpath = adb.Path(pkgname)
	pull(pkgname, pkgpath)

	return pkgname
}

// Allow the user to select the package he wants to work on through a simple part of the package name.
func packageName(pkgnamepart string) string {
	var pkgs []string = adb.ListPackages(pkgnamepart)

	var pkgname string = choose(pkgs)

	return pkgname
}

// Display a list of packages and let the user choose one of them.
// It returns the package name, the user chose.
func choose(pkgs []string) string {

	for i, pkg := range pkgs {
		if pkg != "" {
			fmt.Println(logging.Blue("[") + logging.Red(strconv.Itoa(i+1)) + logging.Blue("] ") + logging.Bold(strings.Split(pkg, ":")[1]))
		}
	}

	// Wait for input from user in order to choose which apk to retrive through adb
	fmt.Println("Which package do you want to investigate?")

	var input string
	fmt.Scanln(&input)

	var i, err = strconv.Atoi(input)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return strings.Split(pkgs[i-1], ":")[1]
}

/**
Get the package path on the connected devices via adb regarding the package name, without the "apk" extension.

Params:
package_name The package name, without the "apk" extension.

Comment: Use the shell command : adb shell pm path 'pkgname'
*/
func path(pkgname string) string {
	var path = adb.Path(pkgname)

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
