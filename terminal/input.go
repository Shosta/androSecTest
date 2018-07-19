package terminal

import (
	"bufio"
	"os"

	"github.com/shosta/androSecTest/logging"
)

// Wait for a user input on the CLI.
// It returns the user input as a string.
func Waitfor() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		logging.PrintlnDebug("User wrote: " + scanner.Text())
		return scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}

	return ""
}
