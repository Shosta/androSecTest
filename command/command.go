package command

import (
	"fmt"
	"os/exec"

	"github.com/shosta/androSecTest/logging"
)

func Run(command string, args []string) string {
	cmd := exec.Command(command, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logging.PrintlnError(fmt.Sprint(err) + ": " + string(output))
		return ""
	}
	return string(output)
}

func RunCmd(cmd string) string {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		logging.PrintlnError("error occured")
		logging.PrintlnError(fmt.Sprint(err))
	}
	fmt.Printf("%s", out)

	return string(out)
}

// func RunAlias(cmd string, args []string) string {
// 	binary, lookErr := exec.LookPath("/bin/bash")
// 	if lookErr != nil {
// 		logging.PrintlnDebug("look")
// 		panic(lookErr)
// 	}

// 	args2 := []string{"-c", "signapk", cmd}

// 	env := os.Environ()

// 	execErr := syscall.Exec(binary, args2, env)
// 	if execErr != nil {
// 		logging.PrintlnDebug("sys")
// 		panic(execErr)
// 	}

// 	return "done"

// 	testCmd := exec.Command("java", "-jar", "/home/shosta/ShostaSyncBox/Developpement/HackingTools/SignApkUtils/sign.jar", cmd+".apk")
// 	testOut, err := testCmd.Output()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(testOut))

// 	return string(testOut)"
// }
