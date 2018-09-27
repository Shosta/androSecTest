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

// Package adb : It executes binaries on an android device using adb.
package adb

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/shosta/androSecTest/logging"
)

func runAdb(args ...string) string {
	cmd := exec.Command("adb", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err) + ": " + string(output))
		return ""
	}
	return string(output)
}

// Pull :
func Pull(pkgpath string, pkgdest string) string {
	logging.PrintlnDebug(logging.Green("Package path : ") + logging.Bold(pkgpath))
	logging.PrintlnDebug(logging.Green("Package destination : ") + logging.Bold(pkgdest))

	cmd := exec.Command("adb", "pull", pkgpath, pkgdest)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err) + ": " + string(output))
		return ""
	}
	return string(output)
}

// PkgPath : Get the path of the package name on the connected device on adb.
// It returns the package's path on the device as a string.
func PkgPath(pkgname string) string {
	logging.PrintlnDebug(logging.Green("Package name: ") + pkgname)

	var args = []string{"shell", "pm", "path", pkgname}
	var out = runAdb(args...)

	var pkgpath = strings.TrimSpace(strings.Split(out, ":")[1])
	logging.Println(logging.Green("Path: ") + pkgpath)

	return pkgpath
}

// ListPackages :
func ListPackages(pkgnamepart string) []string {
	logging.Println(logging.Green("Get the packages names") + " on the device that contains \"" + logging.Bold(pkgnamepart) + "\":")

	var args = []string{"shell", "pm", "list", "packages", "|", "grep " + pkgnamepart}
	var out = runAdb(args...)
	pkgs := strings.Split(out, "\n")

	return pkgs
}

//Uninstall : adb uninstall " + package_name
func Uninstall(pkgname string) string {
	logging.PrintlnDebug("Uninstall app on the device.")
	out := runAdb("uninstall", pkgname)

	return string(out)
}

//Install : adb install /tmp/Attacks/DebuggablePackage/" + package_name + ".b.s.apk"
func Install(localpkgpath string) string {
	logging.PrintlnDebug("Install app on the device.")
	logging.PrintlnDebug("Local package path: " + localpkgpath)
	out := runAdb("install", localpkgpath)

	return string(out)
}

// Devices : List the devices connected to the computer.
func Devices() string {
	logging.PrintlnDebug("List devices connected to the computer.")
	out := runAdb("devices", "-l")

	return string(out)
}
