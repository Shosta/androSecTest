package androidpkg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/shosta/androSecTest/command/adb"
	"github.com/shosta/androSecTest/logging"
)

// Display a list of packages and let the user choose one of them.
// It returns the package name, the user chose.
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

// Allow the user to select the package he wants to work on through a simple part of the package name.
func Package(pkgnamepart string) string {
	var pkgs []string = adb.ListPackages(pkgnamepart)

	var pkgname string = choose(pkgs)

	return pkgname
}
