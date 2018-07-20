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
