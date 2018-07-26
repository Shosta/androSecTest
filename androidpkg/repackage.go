package androidpkg

import (
	"os"

	"github.com/shosta/androSecTest/command/sed"

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
// cmd = "unzip " + attacksDir + variables.SourcePackageDir + "/" + pkgname + ".apk '*' -d " + unzipDir
func unzip(pkgname string) {
	sourceDirPath := folders.SourcePackageDirPath(pkgname)
	unzipDirPath := folders.UnzipDirPath(pkgname)
	logging.Println(logging.Green("Extract package : ") + logging.Bold(pkgname) + " to " + logging.Bold(unzipDirPath))

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

	disassembledDirPath := folders.DisassemblePackageDirPath(pkgname)
	logging.PrintlnVerbose(logging.Green("Extract package : ") + logging.Bold(pkgname) + " to " + logging.Bold(disassembledDirPath))

	sed.Replace(disassembledDirPath+"/AndroidManifest.xml", "s/<application ", "<application android:debuggable=\"true\" ")

	logging.Println(logging.Bold("Done"))
}

// TODO : Should verify that the AllowBackup is not already available in the AppManifest.xml file.
func allowbackup(pkgname string) {
	logging.Println(logging.Green("Allow backup on package"))

	disassembledDir := folders.DisassemblePackageDirPath(pkgname)

	sed.Replace(disassembledDir+"/AndroidManifest.xml", "android:allowBackup=\"false\" ", " ")
	logging.PrintlnDebug("Remove the android:allowBackup=\"false\" if it is in the AndroidManifest.xml file.")

	sed.Replace(disassembledDir+"/AndroidManifest.xml", "<application ", "<application android:allowBackup=\"true\" ")
	logging.PrintlnDebug("Add the android:allowBackup=\"true\" to the AndroidManifest.xml file.")

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
