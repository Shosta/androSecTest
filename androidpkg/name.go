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

// Package androidpkg : Scan the packages on the device that contains a string and allows the user to work on it.
package androidpkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Shosta/androSecTest/command/adb"
	"github.com/Shosta/androSecTest/logging"
)

// Display a list of packages and let the user choose one of them.
// It returns the package name the user chose.
func choose(pkgs []string) string {

	for i, pkg := range pkgs {
		if pkg != "" {
			fmt.Println(logging.Blue("[") + logging.Red(strconv.Itoa(i+1)) + logging.Blue("] ") + logging.Bold(strings.Split(pkg, ":")[1]))
		}
	}

	// Wait for input from user in order to choose which apk to retrive through adb
	logging.Println(logging.Blue("Which package do you want to investigate?"))

	var input string
	fmt.Scanln(&input)

	var i, err = strconv.Atoi(input)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return strings.Split(pkgs[i-1], ":")[1]
}

// Package Allow the user to select the package he wants to work on thanks to a simple part of the package name.
func Package(pkgnamepart string) string {
	var pkgs = adb.ListPackages(pkgnamepart)

	var pkgname = choose(pkgs)

	return pkgname
}
