package attacks

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/shosta/androSecTest/command"
	"github.com/shosta/androSecTest/logging"

	"github.com/shosta/androSecTest/variables"
)

// Get all the occurence of a string through a grep commanc and store it in a file.
func strInLog(str string) {
	var insecureLoggingPath string = variables.SecAssessmentPath + variables.InsecureLoggingDir
	var logFilePath string = insecureLoggingPath + "/log.txt"
	var resFilePath string = insecureLoggingPath + "/grep-" + str + ".txt"
	logging.PrintlnDebug("Insecure Logging file path: " + logFilePath)
	logging.PrintlnDebug("Res file path: " + resFilePath)
	cmd := "grep " + str + " " + logFilePath + " > " + resFilePath
	command.RunAlias(cmd)
}

// Get all the occurence of words related to "password" and store them in a file.
func passwordStrInLog() {
	strInLog("password")
	strInLog("pass")
	strInLog("passwd")
}

// Get all the occurence of words related to "key" and store them in a file.
func keyStrInLog() {
	strInLog("key")
}

// Get all the occurence of words related to "admin" and store them in a file.
func adminStrInLog() {
	strInLog("admin")
	strInLog("adm")
}

// A loop method that ask the user to enter a string, then search it in the log file through a grep command and ask the user if he wants to do another search.
func userInputStrInLog() {
	logging.Print(logging.Blue("Enter the string you want to look for in the log file.\n> "))
	usrinput := userinput()
	if usrinput != "" {
		logging.Println(logging.Green("Looking for \"") + logging.Bold(usrinput) + "\" in log file.")
		strInLog(usrinput)
		logging.Print(logging.Blue("Do you want to look for another string? [y][n]\n> "))
		newSearch := userinput()
		if newSearch == "y" {
			userInputStrInLog()
		}

		return
	}
}

// Wait for a user input on the CLI.
// It returns the user input as a string.
func userinput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		logging.PrintlnDebug("User wrote: " + scanner.Text())
		return scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}

	return ""
}

// Launch a logcat command and push the result to a file.
func launchlogcat(pkgname string) {
	logging.Println("Log svg : " + variables.SecurityAssessmentRootDir + "/" + pkgname + variables.AttacksDir + variables.InsecureLoggingDir + "/log.txt")
	cmd := exec.Command("/bin/sh", "-c", "adb logcat > "+variables.SecurityAssessmentRootDir+"/"+pkgname+variables.AttacksDir+variables.InsecureLoggingDir+"/log.txt")

	// Start command asynchronously
	logging.PrintlnDebug("Launched logcat asynchronously.")
	cmd.Start()
	logging.PrintlnDebug("Wait for any user input to kill the logcat process.")
	userinput()

	logging.PrintlnDebug("Stopped logcat.")
	cmd.Process.Signal(os.Kill)
}

// Test if something insecure is logged through logcat while using the device.
// It tests the "password", "admin" and "key" related strings and then let the user test its own strings.
func InsecureLogging(pkgname string) {
	logging.Println(logging.Green("Test Insecure Logging"))
	launchlogcat(pkgname)

	passwordStrInLog()
	keyStrInLog()
	adminStrInLog()

	userInputStrInLog()
}
