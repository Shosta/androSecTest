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

// Package main : THe main file that starts the program
package main

import (
	arg "github.com/alexflint/go-arg"
	"github.com/shosta/androSecTest/androidpkg"
	"github.com/shosta/androSecTest/attacks"
	"github.com/shosta/androSecTest/config"
	"github.com/shosta/androSecTest/devices"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/settings"
	"github.com/shosta/androSecTest/terminal"
)

func main() {
	var args struct {
		Settings    bool   `arg:"-s" help:"set up the user settings"`
		Package     string `arg:"-p" help:"package name"`
		Dest        string `arg:"-d" help:"destination folder absolute path"`
		Attacksonly bool   `arg:"-a" help:"perform only attacks"`
		Verbose     bool   `arg:"-v" help:"verbosity level"`
	}
	arg.MustParse(&args)
	settings.Setup(args.Settings)

	if !devices.IsConnected() {
		logging.Println(logging.Green("No device is connected.") + "\nPlease " + logging.Bold("connect a device to your computer") + " prior to any penetration testing.")
		return
	}

	pkgname := ""
	if args.Package == "" {
		// TODO : Wait for the user input to get the package.
		logging.PrintDebug("No package provided.\nPlease provide the name of the package you want to test.\n" + logging.Blue(">  "))
		pkgname = terminal.Waitfor()
	} else {
		pkgname = androidpkg.Package(args.Package)
	}

	if args.Verbose {
		config.IsVerboseLogRequired = true
	}
	logging.Println(pkgname)

	if args.Attacksonly == false {
		androidpkg.Savelocal(pkgname)
		androidpkg.Setup(pkgname)
	}

	attacks.Do(pkgname)
}
