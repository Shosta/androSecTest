/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Mozilla Public License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.mozilla.org/en-US/MPL/

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package attacks : does all the attacks at once.
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
	usrInput := terminal.Waitfor()

	if usrInput == "a" {
		DoInsecureLog(pkgname)
		DoInsecureStorage(pkgname)
		DoReverse(pkgname)
	}
}
