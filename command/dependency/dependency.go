package command

import (
	"os/exec"

	"github.com/shosta/androSecTest/logging"
)

func AreAllReady() bool {
	var areAllReady = true

	areAllReady = isAdbInstalled()
	if areAllReady != true {
		return false
	}

	areAllReady = isApktoolInstalled()
	if areAllReady != true {
		return false
	}

	areAllReady = isSignApkInstalled()
	if areAllReady != true {
		return false
	}

	return true
}

func isAdbInstalled() bool {
	path, err := exec.LookPath("adb")
	if err != nil {
		logging.PrintlnError("didn't find 'adb' executable\n")
		return false
	}
	logging.PrintlnVerbose("'adb' executable is in " + path)

	return true
}

func isApktoolInstalled() bool {
	path, err := exec.LookPath("apktool")
	if err != nil {
		logging.PrintlnError("didn't find 'apktool' executable\n")
		return false
	}
	logging.PrintlnVerbose("'apktool' executable is in " + path)

	return true
}

func isSignApkInstalled() bool {
	path, err := exec.LookPath("signapk")
	if err != nil {
		logging.PrintlnError("didn't find 'signapk' executable\n")
		return false
	}
	logging.PrintlnVerbose("'signapk' executable is in " + path)

	return true
}
