package attacks

import (
	"github.com/shosta/androSecTest/command"
	"github.com/shosta/androSecTest/folder"
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/terminal"
)

func copyDump(pkgname string) {
	srcDir := "/home/shosta/ShostaSyncBox/Developpement/HackingTools/humpty-dumpty-android-master/dumps/" + pkgname
	destDir := InsecStorageDirPath(pkgname)
	logging.PrintlnDebug("Source : " + srcDir)
	logging.PrintlnDebug("Dest : " + destDir)

	logging.PrintlnDebug("Delete \"dump\" folder if it exists")
	folder.Delete(destDir)

	logging.PrintlnDebug("Copy dump folder to proper location")
	folder.CopyDir(srcDir, destDir)
}

// humpty.sh -a com.pixplicity.example
func pullLocalStorage(pkgname string) {
	logging.Println(logging.Green("Pull every files from the local storage of the \"" + pkgname + "\" package."))

	logging.Println("In progress...")
	var cmd = "/home/shosta/ShostaSyncBox/Developpement/HackingTools/humpty-dumpty-android-master/humpty.sh -a " + pkgname
	command.RunAlias(cmd)

	logging.Println(logging.Bold("Done"))
}

// DoInsecureStorage :
func DoInsecureStorage(pkgname string) {
	logging.Println(logging.Green("Test Insecure Storage"))
	pullLocalStorage(pkgname)

	logging.Println("Did you use all the features of the application? [y]es [n]o")
	terminal.Waitfor()

	// pullLocalStorage(pkgname)
	copyDump(pkgname)
	logging.Println(logging.Bold("Done"))
}
