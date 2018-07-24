// Package devices : Check that a device is connected and adb commands can then be performed on it.
package devices

import (
	"strings"

	"github.com/shosta/androSecTest/command/adb"
	"github.com/shosta/androSecTest/logging"
)

func connectedDevices(adbOutput string) []string {
	devices := strings.Split(strings.TrimSpace(adbOutput), "List of devices attached")[1]
	devicesArray := strings.Split(devices, "\n")

	return devicesArray[1:len(devicesArray)]
}

// IsConnected : Check is a device is connected through USB to the computer and is ready to receive "adb" commands.
// It uses the command : "adb devices -l" to verify it.
func IsConnected() bool {
	output := adb.Devices()
	logging.PrintlnDebug("Devices : \n" + output)
	devicesArray := connectedDevices(output)

	if len(devicesArray) > 0 {
		logging.PrintlnDebug("No device connected to the computer.")
		return true
	}

	return false
}
