// Package logging contains utility functions for working with logs.
package logging

import (
	"fmt"

	"github.com/shosta/androSecTest/config"
)

// PrintlnDebug :
func PrintlnDebug(log string) {
	if config.IsDebug {
		fmt.Println(Orange("[dbg:info] ") + log)
	}
}

// PrintlnError :
func PrintlnError(err interface{}) {
	if config.IsDebug {
		fmt.Println(Red("[error] ") + fmt.Sprint(err))
	}
}

// PrintlnVerbose : Print the log to the terminal if the
func PrintlnVerbose(log string) {
	if config.IsVerboseLogRequired {
		fmt.Println(log)
	}
}

// Println :
func Println(log string) {
	fmt.Println(log)
}

// Print :
func Print(log string) {
	fmt.Print(log)
}
