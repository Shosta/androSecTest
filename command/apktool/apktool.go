package apktool

import (
	"fmt"
	"os/exec"

	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/variables"
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

// TODO Il faut prendre en compte les cas d'erreurs d'apktool.
func Disassemble(pkgname string) string {
	var attacksDir string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	var sourcePkgDir string = attacksDir + variables.SourcePackageDir
	var decodedDir string = attacksDir + variables.DecodedPackageDir

	cmdArgs := []string{
		"d",
		sourcePkgDir + "/" + pkgname + ".apk",
		"-f",
		"-o",
		decodedDir,
	}

	var output = runApktool(cmdArgs...)
	logging.PrintlnVerbose(output)

	logging.Println(logging.Green("Package disassembled with success") + " to " + logging.Bold(decodedDir))

	return output
}

// TODO Il faut prendre en compte les cas d'erreurs d'apktool.
//cmd = "apktool b /tmp/Attacks/DecodedPackage -o /tmp/Attacks/DebuggablePackage/" + package_name + ".b.apk"
func Build(pkgname string) string {
	var attacksDir string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	var decodedDir string = attacksDir + variables.DecodedPackageDir
	var debugPkgDir string = attacksDir + variables.DebuggablePackageDir

	cmdArgs := []string{
		"b",
		decodedDir,
		"-o",
		debugPkgDir + "/" + pkgname + ".b.apk",
	}

	var output = runApktool(cmdArgs...)
	logging.PrintlnVerbose(output)

	logging.Println(logging.Green("Package built with success") + " to " + logging.Bold(debugPkgDir))

	return output
}
