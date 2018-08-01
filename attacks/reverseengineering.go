package attacks

import (
	"github.com/shosta/androSecTest/command"
	grep "github.com/shosta/androSecTest/command/grep"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/settings"
	"github.com/shosta/androSecTest/variables"
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
