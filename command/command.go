package command

import (
	"fmt"
	"os/exec"

	"github.com/shosta/androSecTest/logging"
)

// Run :
func Run(command string, args []string) string {
	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err) + ": " + string(output))
		return ""
	}
	return string(output)
}

// RunCmd :
func RunCmd(cmd string) string {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		logging.PrintlnError("error occured")
		logging.PrintlnError(fmt.Sprint(err))
	}
	fmt.Printf("%s", out)

	return string(out)
}

// func Start(cmd string, args []string) string {
// 	command := exec.Command(cmd, args...)
// 	err := command.Start()
// 	if err != nil {

// 		log.Fatal(err)
// 	}
// 	log.Printf("Waiting for command to finish...")
// 	err = command.Wait()
// 	log.Printf("Command finished with error: %v", err)
// }

// RunAlias : Run a command that is defined as an alias in ~/.bashrc or ~/.bash_aliases files.
// The aliasCmd is the entire command you want to run.
func RunAlias(aliasCmd string) string {
	out, err := exec.Command("/bin/bash", "-c", aliasCmd).Output()
	if err != nil {
		logging.PrintlnError("error occured")
		logging.PrintlnError(err)
	}
	// logging.PrintlnVerbose(string(out))

	return string(out)
}
