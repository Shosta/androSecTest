// Package log contains utility functions for working with logs.
package logging

import (
	"fmt"

	"github.com/shosta/androSecTest/variables"
)

func PrintlnDebug(log string) {
	if variables.IsDebug {
		fmt.Println(Orange("[info] ") + log)
	}
}

func PrintlnError(log string) {
	if variables.IsDebug {
		fmt.Println(Red("[error] ") + log)
	}
}

// Print the log to the terminal if the
func PrintlnVerbose(log string) {
	if variables.IsVerboseLogRequired {
		fmt.Println(log)
	}
}

func Println(log string) {
	fmt.Println(log)
}
