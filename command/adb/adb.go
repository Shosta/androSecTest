// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program can be used as go_android_GOARCH_exec by the Go tool.
// It executes binaries on an android device using adb.
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

// Pull
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

// Get the path of the package name on the connected device on adb.
// It returns the package's path on the device as a string.
func PkgPath(pkgname string) string {
	logging.PrintlnDebug(logging.Green("Package name: ") + pkgname)

	var args = []string{"shell", "pm", "path", pkgname}
	var out = runAdb(args...)

	var pkgpath = strings.TrimSpace(strings.Split(out, ":")[1])
	logging.Println(logging.Green("Path: ") + pkgpath)

	return pkgpath
}

func ListPackages(pkgnamepart string) []string {
	logging.Println(logging.Green("Get the packages names") + " on the device that contains \"" + logging.Bold(pkgnamepart) + "\":")

	var args = []string{"shell", "pm", "list", "packages", "|", "grep " + pkgnamepart}
	var out = runAdb(args...)
	pkgs := strings.Split(out, "\n")

	return pkgs
}

//adb uninstall " + package_name
func Uninstall(pkgname string) string {
	logging.PrintlnDebug("Uninstall app on the device.")
	out := runAdb("uninstall", pkgname)

	return string(out)
}

//adb install /tmp/Attacks/DebuggablePackage/" + package_name + ".b.s.apk"
func Install(localpkgpath string) string {
	logging.PrintlnDebug("Install app on the device.")
	logging.PrintlnDebug("Local package path: " + localpkgpath)
	out := runAdb("install", localpkgpath)

	return string(out)
}
