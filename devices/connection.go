/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Mozilla Public License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.mozilla.org/en-US/MPL/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

	if len(devicesArray) <= 0 {
		logging.PrintlnDebug(logging.Red("No device") + " connected to the computer.\nPlease connect a device before launching that app.")
		return false
	}

	return true
}
