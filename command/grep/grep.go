package grep

import (
	"github.com/shosta/androSecTest/command"
	"github.com/shosta/androSecTest/logging"
)

// Motif : Check if the files that are in the dest folder have some urls.
// grep command used is : grep -Eo motif -R .
func Motif(motif string, src string, dest string) {
	logging.PrintlnDebug("Src : " + src)
	logging.PrintlnDebug("Dest : " + dest)

	cmd := "grep " + motif + " -R " + src + " > " + dest + "/grep-" + motif + ".txt"
	logging.PrintlnDebug("Run command : " + cmd)
	command.RunAlias(cmd)
}

// HTTP : Check if the files that are in the dest folder have some urls.
// grep command used is : grep -Eo '(http|https)://[^/"]+' -R .
func HTTP(src string, dest string) {
	Motif("'(http|https)://[^/\"]+'", src, dest)
}

// Admin : Check if the files that are in the dest folder have the following patterns in it, "admin", "adm".
// It stores the results in a specific file, named as followed grep-"motif".txt, in the "dest" folder.
func Admin(src string, dest string) {
	Motif("admin", src, dest)
	Motif("adm", src, dest)
}

// Passwd : Check if the files that are in the dest folder have the following patterns in it, "pass", "passwd", "password".
// It stores the results in a specific file, named as followed grep-"motif".txt, in the "dest" folder.
func Passwd(src string, dest string) {
	Motif("password", src, dest)
	Motif("passwd", src, dest)
	Motif("pass", src, dest)
}

// Key : Check if the files that are in the dest folder have the following patterns in it, "key".
// It stores the results in a specific file, named as followed grep-"motif".txt, in the "dest" folder.
func Key(src string, dest string) {
	Motif("key", src, dest)
}
