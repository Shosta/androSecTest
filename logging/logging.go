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

func PrintlnError(err interface{}) {
	if variables.IsDebug {
		fmt.Println(Red("[error] ") + fmt.Sprint(err))
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

func Print(log string) {
	fmt.Print(log)
}
