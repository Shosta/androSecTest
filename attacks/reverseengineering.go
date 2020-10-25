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

// Package attacks : does the reverse engineering attack on the application through decompiling the application and looling at improper strings in it.
package attacks

import (
	"github.com/Shosta/androSecTest/command"
	grep "github.com/Shosta/androSecTest/command/grep"
	"github.com/Shosta/androSecTest/logging"
	"github.com/Shosta/androSecTest/settings"
	"github.com/Shosta/androSecTest/variables"
)

// Voici le commande à utiliser :
// ./ShostaSyncBox/Developpement/HackingTools/DecompilingAndroidApp/jadx/bin/jadx --deobf -d ~/android/security/com.orange.owtv/attacks/decodedPackage ~/android/security/com.orange.owtv/attacks/sourcePackage/com.orange.owtv.apk
func reverseApk(apkname string) {
	// TODO : Il faut changer le chemin absolu vers le binaire de jadx pour que cela soit rentré par l'utilisateur dans un fichier settings.

	cmd := settings.Jadx() + " " +
		UnzipDirPath(apkname) + "/classes.dex" + " " +
		"-d " + DecompiledPackageDirPath(apkname) + " " +
		"--deobf"

	logging.PrintlnDebug("Cmd : " + cmd)

	logging.Println("Decompiling apk to " + logging.Bold(apkname+"/attacks/decodedPackage/") + "\nWork in progress...")
	command.RunAlias(cmd)
	logging.PrintlnDebug(cmd)
}

// DoReverse : Reverse the ".apk" to the ".java" files.
// Try to deobfuscate code while reversing it.
// Then it performs some research for specific leak in the codebase, looking for strings as "password", "admin", "key", etc. The results are stored in specific files.
func DoReverse(pkgname string) {
	logging.Println(logging.Green("Reverse apk"))
	reverseApk(pkgname)
	logging.Println(logging.Bold("Done"))

	logging.Println(logging.Green("Check for leakage in codebase") + "\nWork in progress...")
	checkForLeaks(pkgname)
	logging.Println(logging.Bold("Done"))
}

func checkForLeaks(pkgname string) {
	decoPkgPath := DecompiledPackageDirPath(pkgname)
	createLeakageDir(pkgname)
	grep.Passwd(decoPkgPath, decoPkgPath+variables.LeakagesDir)
	grep.Admin(decoPkgPath, decoPkgPath+variables.LeakagesDir)
	grep.Key(decoPkgPath, decoPkgPath+variables.LeakagesDir)
}
