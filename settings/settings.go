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
	us, err := loadUsrSettings()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err))
		return
	}

	isApktoolInstalled, apktoolpath := IsApktoolInstalled()
	us.Tools.Apktool = apktoolpath

	if us.Tools.Jadx == "" || us.Tools.SignApk == "" || us.HackingTools.HumptyDumpty == "" || force == true {
		// Ask user to fill in the tools paths.
		logging.Print(logging.Green("Where is located Jadx?") + " (copy and paste the absolute path to your jadx executable, should look like \"/home/user/hacking/tools/jadx/jadx\"\n" + logging.Blue(">  "))
		jadxuserentry := terminal.Waitfor()	
		us.Tools.Jadx = jadxuserentry

		logging.Print(logging.Green("Where is located SignApk?") + " (copy and paste the absolute path to your signapk jar file, should look like \"/home/user/hacking/tools/signapk/sign.jar\"\n" + logging.Blue(">  "))
		signuserentry := terminal.Waitfor()
		us.Tools.SignApk = signuserentry

		if isApktoolInstalled == false {
			logging.Print(logging.Green("Where is located ApkTool?") + " (copy and paste the absolute path to your apktool executable, should look like \"/usr/local/bin/apktool\"\n" + logging.Blue(">  "))
			apktuserentry := terminal.Waitfor()
			us.Tools.Apktool = apktuserentry
		} else {
			us.Tools.Apktool = apktoolpath
		}

		logging.Print(logging.Green("Where is located Humpty-Dumpty?") + " (copy and paste the absolute path to your humpty-dumpty shell file, should look like \"/home/user/hacking/tools/humpty-dumpty/humpty.sh\"\n" + logging.Blue(">  "))
		hduserentry := terminal.Waitfor()
		us.HackingTools.HumptyDumpty = hduserentry

		saveUsrSettings(us)
	}
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
func loadUsrSettings() (UserSettings, error) {
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
	us, _ := loadUsrSettings()
	jadxpath = path
	us.Tools.Jadx = path
	saveUsrSettings(us)
}

// SetSignapk : Set the path to the SignApk executable in the settings file so that wee could call it later.
func setSignapk(path string) {
	us, _ := loadUsrSettings()
	signapkpath = path
	us.Tools.SignApk = path
	saveUsrSettings(us)
}

// SetApktool : Set the path to the apktool executable in the settings file so that wee could call it later.
func setApktool(path string) {
	us, _ := loadUsrSettings()
	apktoolpath = path
	us.Tools.Apktool = path
	saveUsrSettings(us)
}

// SetHumptyDumpty : Set the path to the humpty-dumpty shell script in the settings file so that wee could call it later.
func setHumptyDumpty(path string) {
	us, _ := loadUsrSettings()
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
