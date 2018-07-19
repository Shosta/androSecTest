package attacks

import (
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/terminal"
)

// Do :
func Do(pkgname string) {
	logging.Print(logging.Blue("What kind of attacks do you want to perform? [") +
		logging.Red("a") +
		logging.Blue("]ll, [") +
		logging.Red("i") +
		logging.Blue("]nsecure logging\n> "))
	usrinput := terminal.Waitfor()

	if usrinput == "a" {
		DoInsecureLog(pkgname)
		DoInsecureStorage(pkgname)
	}
}
