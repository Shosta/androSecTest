package androidpkg

import (
	"github.com/shosta/androSecTest/command"
	"github.com/shosta/androSecTest/command/adb"
	"github.com/shosta/androSecTest/command/apktool"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/variables"
)

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
	var attacksDir string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	var unzipDir string = attacksDir + variables.UnzippedPackageDir
	logging.Println(logging.Green("Extract package : ") + logging.Bold(pkgname) + " to " + logging.Bold(unzipDir))

	//var cmd string = "unzip " + attacksDir + variables.SourcePackageDir + "/" + pkgname + ".apk '*' -d " + unzipDir
	cmdName := "unzip"
	cmdArgs := []string{
		attacksDir + variables.SourcePackageDir + "/" + pkgname + ".apk",
		"-d",
		unzipDir,
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

	var attacksDir string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	var decodedDir string = attacksDir + variables.DecodedPackageDir
	logging.PrintlnVerbose(logging.Green("Extract package : ") + logging.Bold(pkgname) + " to " + logging.Bold(decodedDir))

	//cmd = "sed -i -e 's/<application /<application android:debuggable=\"true\" /' /tmp/Attacks/DecodedPackage/AndroidManifest.xml"
	cmdName := "sed"
	cmdArgs := []string{
		"-i",
		"-e",
		"s/<application /<application android:debuggable=\"true\" /",
		decodedDir + "/AndroidManifest.xml",
	}
	command.Run(cmdName, cmdArgs)

	logging.Println(logging.Bold("Done"))
}

func allowbackup(pkgname string) {
	logging.Println(logging.Green("Allow backup on package"))

	var attacksDir string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir
	var decodedDir string = attacksDir + variables.DecodedPackageDir

	//cmd = "sed -i -e 's/android:allowBackup=\"false\" / /' /tmp/Attacks/DecodedPackage/AndroidManifest.xml"
	cmdName := "sed"
	cmdArgs := []string{
		"-i",
		"-e",
		"s/android:allowBackup=\"false\" / /",
		decodedDir + "/AndroidManifest.xml",
	}
	logging.PrintlnDebug("Remove the android:allowBackup=\"false\" if it is in the AndroidManifest.xml file.")
	command.Run(cmdName, cmdArgs)

	//cmd = "sed -i -e 's/<application /<application android:allowBackup=\"true\" /' /tmp/Attacks/DecodedPackage/AndroidManifest.xml"
	cmdArgs = []string{
		"-i",
		"-e",
		"s/<application /<application android:allowBackup=\"true\" /",
		decodedDir + "/AndroidManifest.xml",
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
	var pkgLoc string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.DebuggablePackageDir + "/" + pkgname
	logging.PrintlnDebug("Package to sign : " + pkgLoc + ".b.apk")
	//var cmd string = "java -jar /home/shosta/ShostaSyncBox/Developpement/HackingTools/SignApkUtils/sign.jar " + pkgLoc + "b.apk"

	cmdArgs := []string{
		"-jar",
		"/home/shosta/ShostaSyncBox/Developpement/HackingTools/SignApkUtils/sign.jar",
		pkgLoc + ".b.apk",
	}
	command.Run("java", cmdArgs)

	// cmdArgs := []string{
	// 	"-i",
	// 	"-c",
	// 	"signapk",
	// 	pkgLoc + ".b.apk",
	// }
	// command.RunAlias("/bin/bash -i -c signapk " + pkgLoc + ".b.apk")
	// command.RunCmd("bash -ic -rcfile ~/.bash_aliases signapk " + pkgLoc + ".b.apk")
}

//adb uninstall " + package_name
//adb install /tmp/Attacks/DebuggablePackage/" + package_name + ".b.s.apk"
func reinstall(pkgname string) {
	adb.Uninstall(pkgname)

	var localpkgPath string = variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.DebuggablePackageDir + "/" + pkgname + ".b.s.apk"
	adb.Install(localpkgPath)
}
