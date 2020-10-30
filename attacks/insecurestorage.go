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

// Package attacks : Tests the insecure storage
package attacks

import (
	"github.com/Shosta/androSecTest/command"
	"github.com/Shosta/androSecTest/folder"
	"github.com/Shosta/androSecTest/logging"
	"github.com/Shosta/androSecTest/terminal"
)

func copyDump(pkgname string) {
	srcDir := "/home/shosta/ShostaSyncBox/Developpement/HackingTools/humpty-dumpty-android-master/dumps/" + pkgname
	destDir := InsecStorageDirPath(pkgname) + "dumps"
	logging.PrintlnDebug("Source : " + srcDir)
	logging.PrintlnDebug("Dest : " + destDir)

	logging.PrintlnDebug("Delete \"dumps\" folder if it exists")
	folder.Delete(destDir)

	logging.PrintlnDebug("Copy \"dump\" folder to proper location")
	folder.CopyDir(srcDir, destDir)
}

// humpty.sh -a com.pixplicity.example
func pullLocalStorage(pkgname string) {
	logging.Println(logging.Green("Pull every files from the local storage of the \"" + pkgname + "\" package."))

	logging.Println("Work in progress...")
	var cmd = "/home/shosta/ShostaSyncBox/Developpement/HackingTools/humpty-dumpty-android-master/humpty.sh -a " + pkgname
	command.RunAlias(cmd)

	logging.Println(logging.Bold("Done"))
}

// DoInsecureStorage :
// TODO : Understand why Insecure Storage is not working. Even if we have a debuggable application.
func DoInsecureStorage(pkgname string) {
	logging.Println(logging.Green("Test Insecure Storage"))

	logging.Println(logging.Blue("Did you use all the features of the application?") + "[" + logging.Red("y") + "]es [" + logging.Red("n") + "]o")
	terminal.Waitfor()

	pullLocalStorage(pkgname)
	copyDump(pkgname)
	logging.Println(logging.Bold("Done"))
}
