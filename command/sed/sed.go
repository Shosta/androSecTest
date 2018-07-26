package sed

import (
	"github.com/shosta/androSecTest/command"
)

// Replace : Replace a motif in a file through a sed bash command.
// Here is an example of a command: cmd = "sed -i -e 's/src/replace/' filePath"
// and the function arguments :
// filePath := "/absolute/Path/to/file.ext"
// src 		:= "isAdmin=false"
// replace 	:= "isAdmin=true"
func Replace(filePath string, src string, replace string) {

	cmdName := "sed"
	cmdArgs := []string{
		"-i",
		"-e",
		"s/" + src + "/" + replace + "/",
		filePath,
	}
	command.Run(cmdName, cmdArgs)
}
