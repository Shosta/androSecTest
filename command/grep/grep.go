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

// Package grep : It executes grep commands on a computer.
package grep

import (
	"github.com/Shosta/androSecTest/command"
	"github.com/Shosta/androSecTest/logging"
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
