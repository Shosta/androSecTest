/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Mozilla Public License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.mozilla.org/en-US/MPL/

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package settings : Provides features to store the user settings in a file.
// It stores or updates the path to the executables that are required to pursue the penetration testing.
package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Shosta/androSecTest/logging"
	"github.com/Shosta/androSecTest/terminal"
)

var version string
var jadxpath string
var apktoolpath string
var signapkpath string
var humptydumptypath string

// UserSettings :
type UserSettings struct {
	Application  Application  `json:"application"`
	Tools        Tools        `json:"tools"`
	HackingTools HackingTools `json:"hackingtools"`
}

// Application :
type Application struct {
	Version string `json:"version"`
}

// Tools :
type Tools struct {
	Jadx    string `json:"jadx"`
	Apktool string `json:"apktool"`
	SignApk string `json:"signapk"`
}

// HackingTools :HumptyDumpty
type HackingTools struct {
	HumptyDumpty string `json:"humpty-dumpty"`
}

// Setup : It does the Settings set up if we don't know where to look for the external tools required to do the repackaging and the attacks.
// You can use force=true to setup the settings whatever the values in the usersettings.json file.
func Setup(force bool) {
	us, err := LoadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return
	}

	// Check adb
	isAdbInstalled := isAdbInstalled()
	if isAdbInstalled == false {
		logging.PrintlnError("Adb is mandatory for this app to run.\nInstall it (\"sudo apt-get install adb\"")
		return
	}

	// Check apktool
	isApktoolInstalled, apktoolpath := IsApktoolInstalled()
	if isApktoolInstalled == false {
		logging.Print(logging.Green("Where is located ApkTool?") + " (copy and paste the absolute path to your apktool executable, should look like \"/home/Developpement/HackingTools/DecompilingAndroidAppUtils/apktool/apktool.jar\"\n" + logging.Blue(">  "))
		apktuserentry := terminal.Waitfor()
		us.Tools.Apktool = apktuserentry
	} else {
		us.Tools.Apktool = apktoolpath
	}

	// Check Jadx
	isJadxInstalled, jadxpath := IsJadxInstalled()
	if isJadxInstalled == false {
		logging.Print(logging.Green("Where is located Jadx?") + " (copy and paste the absolute path to your jadx executable, should look like \"/home/Developpement/HackingTools/DecompilingAndroidAppUtils/jadx/bin/jadx\"\n" + logging.Blue(">  "))
		jadxuserentry := terminal.Waitfor()
		us.Tools.Jadx = jadxuserentry
	} else {
		us.Tools.Jadx = jadxpath
	}

	// Check SignApk
	isSignApkInstalled, signapkpath := IsSignApkInstalled()
	if isSignApkInstalled == false {
		logging.Print(logging.Green("Where is located SignApk?") + " (copy and paste the absolute path to your signapk jar file, should look like \"/home/Developpement/HackingTools/SignApkUtils/signapk.jar\"\n" + logging.Blue(">  "))
		signuserentry := terminal.Waitfor()
		us.Tools.SignApk = signuserentry
	} else {
		us.Tools.SignApk = signapkpath
	}

	// Check Humpty-Dumpty
	isHumptyDumptyInstalled, hdpath := IsHumptyDumptyInstalled()
	if isHumptyDumptyInstalled == false {
		logging.Print(logging.Green("Where is located humpty-dumpty ?") + " (copy and paste the absolute path to your humpty-dumpty.sh file, should look like \"/home/Developpement/HackingTools/humpty-dumpty-android-master/humpty.sh\"\n" + logging.Blue(">  "))
		hduserentry := terminal.Waitfor()
		us.HackingTools.HumptyDumpty = hduserentry
	} else {
		us.HackingTools.HumptyDumpty = hdpath
	}

	saveUsrSettings(us)
}

// It saves the User defined settings into a Json file.
// So that we can rely on these path when using these tools later on.
func saveUsrSettings(us UserSettings) error {
	bytes, err := json.MarshalIndent(us, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./.res/settings/usersettings.json", bytes, 0644)
}

// It reads the Settings file and update the variables accordingly.
// When using these tools, we can rely on the user defined ones instead of something hard coded.
func LoadUsrSettings() (UserSettings, error) {
	// TODO il faut déplacer le dossier .res dans /home pour éviter les pbs dans git avec le user root lorsque l'on partage les fichiers entre le container et le local.
	bytes, err := ioutil.ReadFile("./.res/settings/usersettings.json")
	if err != nil {
		return UserSettings{}, err
	}

	var s UserSettings
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		return UserSettings{}, err
	}

	version = s.Application.Version
	jadxpath = s.Tools.Jadx
	apktoolpath = s.Tools.Apktool
	signapkpath = s.Tools.SignApk
	humptydumptypath = s.HackingTools.HumptyDumpty

	return s, nil
}

// SetJadx : Set the path to the Jadx executable in the settings file so that wee could call it later.
func setJadx(path string) {
	us, _ := LoadUsrSettings()
	jadxpath = path
	us.Tools.Jadx = path
	saveUsrSettings(us)
}

// SetSignapk : Set the path to the SignApk executable in the settings file so that wee could call it later.
func setSignapk(path string) {
	us, _ := LoadUsrSettings()
	signapkpath = path
	us.Tools.SignApk = path
	saveUsrSettings(us)
}

// SetApktool : Set the path to the apktool executable in the settings file so that wee could call it later.
func setApktool(path string) {
	us, _ := LoadUsrSettings()
	apktoolpath = path
	us.Tools.Apktool = path
	saveUsrSettings(us)
}

// SetHumptyDumpty : Set the path to the humpty-dumpty shell script in the settings file so that wee could call it later.
func setHumptyDumpty(path string) {
	us, _ := LoadUsrSettings()
	humptydumptypath = path
	us.HackingTools.HumptyDumpty = path
	saveUsrSettings(us)
}

// Jadx :
func Jadx() string {
	return jadxpath
}

// SignApk :
func SignApk() string {
	return signapkpath
}

// ApkTool :
func ApkTool() string {
	return apktoolpath
}

// HumptyDumpty :
func HumptyDumpty() string {
	return humptydumptypath
}
