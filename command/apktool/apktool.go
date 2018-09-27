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

// Package apktool : It executes binaries on an computer using apktool.
package apktool

import (
	"fmt"
	"os/exec"

	"github.com/shosta/androSecTest/attacks"

	"github.com/shosta/androSecTest/logging"
)

func runApktool(args ...string) string {
	cmd := exec.Command("apktool", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err) + ": " + string(output))
		return ""
	}
	return string(output)
}

// Disassemble :
// TODO Il faut prendre en compte les cas d'erreurs d'apktool.
func Disassemble(pkgname string) string {

	cmdArgs := []string{
		"d",
		attacks.SourcePackageDirPath(pkgname) + "/" + pkgname + ".apk",
		"-f",
		"-o",
		attacks.DisassemblePackageDirPath(pkgname),
	}

	var output = runApktool(cmdArgs...)
	logging.PrintlnVerbose(output)

	logging.Println(logging.Green("Package disassembled with success") + " to " + logging.Bold(attacks.DisassemblePackageDirPath(pkgname)))

	return output
}

// Build :
// TODO Il faut prendre en compte les cas d'erreurs d'apktool.
//cmd = "apktool b /tmp/Attacks/DecodedPackage -o /tmp/Attacks/DebuggablePackage/" + package_name + ".b.apk"
func Build(pkgname string) string {

	cmdArgs := []string{
		"b",
		attacks.DisassemblePackageDirPath(pkgname),
		"-o",
		attacks.DebugPkgDirPath(pkgname) + "/" + pkgname + ".b.apk",
	}

	var output = runApktool(cmdArgs...)
	logging.PrintlnVerbose(output)

	logging.Println(logging.Green("Package built with success") + " to " + logging.Bold(attacks.DebugPkgDirPath(pkgname)))

	return output
}
