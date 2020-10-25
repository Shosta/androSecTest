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

// Package androidpkg : Save the local files on the computer to then check the Insecure Storage vulnerability.
package androidpkg

import (
	"github.com/Shosta/androSecTest/attacks"
	"github.com/Shosta/androSecTest/command/adb"
	"github.com/Shosta/androSecTest/logging"
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

// Pull a package from the connected devices via adb.
func pull(pkgname string, pkgpath string) {
	destLocation := attacks.SourcePackageDirPath(pkgname) + "/" + pkgname + ".apk"
	logging.Println(logging.Green("Pull package from ") + logging.Bold(pkgpath))

	var out = adb.Pull(pkgpath, destLocation)

	logging.PrintlnVerbose(out)
	logging.Println(logging.Green("Package stored at ") + logging.Bold(destLocation))
}
