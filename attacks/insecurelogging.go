package attacks

import (
	"os"
	"os/exec"

	"github.com/shosta/androSecTest/command"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/terminal"

	"github.com/shosta/androSecTest/variables"
)

// Get all the occurence of a string through a grep commanc and store it in a file.
func strInLog(str string, pkgname string) {

	var logFilePath = InsecLoggingDirPath(pkgname) + "/log.txt"
	var resFilePath = InsecLoggingDirPath(pkgname) + "/grep-" + str + ".txt"
	logging.PrintlnDebug("Insecure Logging file path: " + logFilePath)
	logging.PrintlnDebug("Res file path: " + resFilePath)
	cmd := "grep " + str + " " + logFilePath + " > " + resFilePath
	command.RunAlias(cmd)
}

// Get all the occurence of words related to "password" and store them in a file.
func passwordStrInLog(pkgname string) {
	strInLog("password", pkgname)
	strInLog("pass", pkgname)
	strInLog("passwd", pkgname)
}

// Get all the occurence of words related to "key" and store them in a file.
func keyStrInLog(pkgname string) {
	strInLog("key", pkgname)
}

// Get all the occurence of words related to "admin" and store them in a file.
func adminStrInLog(pkgname string) {
	strInLog("admin", pkgname)
	strInLog("adm", pkgname)
}

// A loop method that ask the user to enter a string, then search it in the log file through a grep command and ask the user if he wants to do another search.
func userInputStrInLog(pkgname string) {
	logging.Print(logging.Blue("Enter the string you want to look for in the log file.\n> "))
	usrinput := terminal.Waitfor()
	if usrinput != "" {
		logging.Println(logging.Green("Looking for \"") + logging.Bold(usrinput) + "\" in log file.")
		strInLog(usrinput, pkgname)
		logging.Print(logging.Blue("Do you want to look for another string? [y][n]\n> "))
		newSearch := terminal.Waitfor()
		if newSearch == "y" {
			userInputStrInLog(pkgname)
		}

		return
	}
}

// Launch a logcat command and push the result to a file.
func launchlogcat(pkgname string) {
	logging.Println("Log svg : " + variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.InsecureLoggingDir + "/log.txt")
	cmd := exec.Command("/bin/sh", "-c", "adb logcat > "+variables.SecurityAssessmentRootDir+"/"+pkgname+variables.AttacksDir+variables.InsecureLoggingDir+"/log.txt")

	// Start command asynchronously
	logging.PrintlnDebug("Launched logcat asynchronously.")
	cmd.Start()
	logging.PrintlnDebug("Wait for any user input to kill the logcat process.")
	logging.Println("Press any key to stop getting logs.")
	terminal.Waitfor()

	logging.PrintlnDebug("Stopped logcat.")
	cmd.Process.Signal(os.Kill)
}

// DoInsecureLog : Test if something insecure is logged through logcat while using the device.
// It tests the "password", "admin" and "key" related strings and then let the user test its own strings.
func DoInsecureLog(pkgname string) {
	logging.Println(logging.Green("Test Insecure Logging"))
	launchlogcat(pkgname)

	passwordStrInLog(pkgname)
	keyStrInLog(pkgname)
	adminStrInLog(pkgname)

	userInputStrInLog(pkgname)
}
