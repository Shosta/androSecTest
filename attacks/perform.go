package attacks

import (
	"github.com/shosta/androSecTest/logging"
	"github.com/shosta/androSecTest/terminal"
)

func Do(pkgname string) {
	logging.Print(logging.Blue("What kind of attacks do you want to perform? [") +
		logging.Red("A") +
		logging.Blue("]ll, [") +
		logging.Red("I") +
		logging.Blue("]nsecure logging\n> "))
	usrinput := terminal.Waitfor()

	if usrinput == "i" {
		DoInsecureLog(pkgname)
	}
}
