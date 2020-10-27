package settings

import (
	"fmt"
	"os/exec"

	"github.com/Shosta/androSecTest/logging"
)

// AreAllReady :
func AreAllReady() bool {
	var areAllReady = true

	areAllReady = isAdbInstalled()
	if areAllReady != true {
		return false
	}

	areAllReady, _ = IsApktoolInstalled()
	if areAllReady != true {
		return false
	}

	areAllReady = isSignApkInstalled()
	if areAllReady != true {
		return false
	}

	areAllReady = isJadxInstalled()
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

// IsApktoolInstalled : Return if apktool is in the user's PATH so that we could call it directly when executing a command.
func IsApktoolInstalled() (bool, string) {
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false, ""
	}

	path, err := exec.LookPath(us.Tools.Apktool)
	if err != nil {
		logging.PrintlnError("didn't find 'apktool' executable\n")
		return false, ""
	} else if err != nil {
		logging.PrintlnError("didn't find 'apktool' executable\n")
		return false, ""
	}
	logging.PrintlnVerbose("'apktool' executable is in " + path)

	return true, path
}

// TODO : Move the signapk executable path to an external folder.
// Add a setup process at the beginning of the program. And an argument to redo the setup if necessary.
func isSignApkInstalled() bool {
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false
	}

	path, err := exec.LookPath(us.Tools.SignApk)
	if err != nil {
		logging.PrintlnError("didn't find 'signapk' executable\n")
		return false
	}
	logging.PrintlnVerbose("'signapk' executable is in " + path)

	return true
}

// TODO : Move the jadx executable path to an external folder.
// Add a setup process at the beginning of the program. And an argument to redo the setup if necessary.
func isJadxInstalled() bool {
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false
	}

	err = nil
	path, err := exec.LookPath(us.Tools.Jadx)
	if err != nil {
		logging.PrintlnError("didn't find 'jadx' executable\n")
		return false
	}
	logging.PrintlnVerbose("'jadx' executable is in " + path)

	return true
}
