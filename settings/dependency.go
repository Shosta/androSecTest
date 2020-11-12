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

	areAllReady, _ = IsSignApkInstalled()
	if areAllReady != true {
		return false
	}

	areAllReady, _ = IsJadxInstalled()
	if areAllReady != true {
		return false
	}

	return true
}

// isAdbInstalled : Check if adb is in the PATH
func isAdbInstalled() bool {
	path, err := exec.LookPath("adb")
	if err != nil {
		logging.PrintlnError("didn't find 'adb' executable\n")
		return false
	}
	logging.PrintlnVerbose("'adb' executable is in " + path)

	return true
}

// IsApktoolInstalled : Check if the path in the User Settings for Apktool is valid
func IsApktoolInstalled() (bool, string) {
	us, err := loadUsrSettings()
	path, err := exec.LookPath(us.Tools.Apktool)
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false, ""
	}

	path, err = exec.LookPath(us.Tools.Apktool)
	if err != nil {
		logging.PrintlnError("didn't find 'apktool' executable\n")
		return false, ""
	}
	logging.PrintlnVerbose("'apktool' executable is in " + path)

	return true, path
}

// IsSignApkInstalled : Check if the path in the User Settings for SignApk is valid
func IsSignApkInstalled() (bool, string) {
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false, ""
	}

	path, err := exec.LookPath(us.Tools.SignApk)
	if err != nil {
		logging.PrintlnError("didn't find 'signapk' executable\n")
		return false, ""
	}
	logging.PrintlnVerbose("'signapk' executable is in " + path)

	return true, path
}

// IsJadxInstalled : Check if the path in the User Settings for Jdax is valid
func IsJadxInstalled() (bool, string) {
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false, ""
	}

	err = nil
	path, err := exec.LookPath(us.Tools.Jadx)
	if err != nil {
		logging.PrintlnError("didn't find 'jadx' executable\n")
		return false, ""
	}
	logging.PrintlnVerbose("'jadx' executable is in " + path)

	return true, path
}

// IsHumptyDumptyInstalled : Check if the path in the User Settings for Humpty-Dumpty is valid
func IsHumptyDumptyInstalled() (bool, string) {
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return false, ""
	}

	err = nil
	path, err := exec.LookPath(us.HackingTools.HumptyDumpty)
	if err != nil {
		logging.PrintlnError("didn't find 'humpty-dumpty' executable\n")
		return false, ""
	}
	logging.PrintlnVerbose("'humpty-dumpty' executable is in " + path)

	return true, path
}
