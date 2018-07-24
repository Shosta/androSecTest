package androidpkg

import (
	"os"

	folders "github.com/shosta/androSecTest/attacks"
	"github.com/shosta/androSecTest/command"
	"github.com/shosta/androSecTest/command/adb"
	"github.com/shosta/androSecTest/command/apktool"
	"github.com/shosta/androSecTest/logging"
)

// Setup :
func Setup(pkgname string) {
	unzip(pkgname)
	disassemble(pkgname)
	mkdbg(pkgname)
	allowbackup(pkgname)
	rebuild(pkgname)
	sign(pkgname)
	reinstall(pkgname)
}

// Unzip the package to the 'unzippedPackage' Folder
func unzip(pkgname string) {
	// var attacksDir = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	// var unzipDir = attacksDir + variables.UnzippedPackageDir
	sourceDirPath := folders.SourcePackageDirPath(pkgname)
	unzipDirPath := folders.UnzipDirPath(pkgname)
	logging.Println(logging.Green("Extract package : ") + logging.Bold(pkgname) + " to " + logging.Bold(unzipDirPath))

	//var cmd string = "unzip " + attacksDir + variables.SourcePackageDir + "/" + pkgname + ".apk '*' -d " + unzipDir
	cmdName := "unzip"
	cmdArgs := []string{
		sourceDirPath + "/" + pkgname + ".apk",
		"-d",
		unzipDirPath,
	}
	command.Run(cmdName, cmdArgs)

	logging.Println(logging.Bold("Done"))
}

// Disassemble the package using the apktool that is installed on the system.
func disassemble(pkgname string) {
	logging.Println(logging.Green("Disassemble package : ") + logging.Bold(pkgname))

	apktool.Disassemble(pkgname)

	logging.Println(logging.Bold("Done"))
}

func mkdbg(pkgname string) {
	logging.Println(logging.Green("Make package debuggable"))

	// var attacksDir = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	// var decodedDir = attacksDir + variables.DisassemblePackageDir
	disassembledDirPath := folders.DisassemblePackageDirPath(pkgname)
	logging.PrintlnVerbose(logging.Green("Extract package : ") + logging.Bold(pkgname) + " to " + logging.Bold(disassembledDirPath))

	//cmd = "sed -i -e 's/<application /<application android:debuggable=\"true\" /' /tmp/Attacks/DecodedPackage/AndroidManifest.xml"
	cmdName := "sed"
	cmdArgs := []string{
		"-i",
		"-e",
		"s/<application /<application android:debuggable=\"true\" /",
		disassembledDirPath + "/AndroidManifest.xml",
	}
	command.Run(cmdName, cmdArgs)

	logging.Println(logging.Bold("Done"))
}

// TODO : Should verify that the AllowBackup is not already available in the AppManifest.xml file.
func allowbackup(pkgname string) {
	logging.Println(logging.Green("Allow backup on package"))

	// var attacksDir = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	// var decodedDir = attacksDir + variables.DisassemblePackageDir
	disassembledDir := folders.DisassemblePackageDirPath(pkgname)

	//cmd = "sed -i -e 's/android:allowBackup=\"false\" / /' /tmp/Attacks/DecodedPackage/AndroidManifest.xml"
	cmdName := "sed"
	cmdArgs := []string{
		"-i",
		"-e",
		"s/android:allowBackup=\"false\" / /",
		disassembledDir + "/AndroidManifest.xml",
	}
	logging.PrintlnDebug("Remove the android:allowBackup=\"false\" if it is in the AndroidManifest.xml file.")
	command.Run(cmdName, cmdArgs)

	//cmd = "sed -i -e 's/<application /<application android:allowBackup=\"true\" /' /tmp/Attacks/DecodedPackage/AndroidManifest.xml"
	cmdArgs = []string{
		"-i",
		"-e",
		"s/<application /<application android:allowBackup=\"true\" /",
		disassembledDir + "/AndroidManifest.xml",
	}
	logging.PrintlnDebug("Add the android:allowBackup=\"true\" to the AndroidManifest.xml file.")
	command.Run(cmdName, cmdArgs)

	logging.Println(logging.Bold("Done"))
}

func rebuild(pkgname string) {
	logging.Println(logging.Green("Rebuild package : ") + logging.Bold(pkgname))

	apktool.Build(pkgname)

	logging.Println(logging.Bold("Done"))
}

//signapk /tmp/Attacks/DebuggablePackage/" + package_name + ".b.apk"
func sign(pkgname string) {
	pkgFilePath := folders.DebugPkgDirPath(pkgname) + "/" + pkgname
	logging.PrintlnDebug("Package to sign : " + pkgFilePath + ".b.apk")
	//var cmd string = "java -jar /home/shosta/ShostaSyncBox/Developpement/HackingTools/SignApkUtils/sign.jar " + pkgLoc + "b.apk"

	cmdArgs := []string{
		"-jar",
		"/home/shosta/ShostaSyncBox/Developpement/HackingTools/SignApkUtils/sign.jar",
		pkgFilePath + ".b.apk",
	}
	command.Run("java", cmdArgs)
	// command.RunAlias("/bin/bash", "-c", "signapk " + pkgLoc + ".b.apk")
}

//adb uninstall " + package_name
//adb install /tmp/Attacks/DebuggablePackage/" + package_name + ".b.s.apk"
func reinstall(pkgname string) {
	pkgFilePath := folders.DebugPkgDirPath(pkgname) + "/" + pkgname + ".b.s.apk"
	if _, err := os.Stat(pkgFilePath); os.IsNotExist(err) {
		logging.PrintlnError("Debuggable pakcage does not exist. Please review the repackaging errors and retry the process before we can install it on the device.")

		return
	}

	adb.Uninstall(pkgname)
	adb.Install(pkgFilePath)
}
