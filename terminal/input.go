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

// Package terminal : Provides the features to read a user input from the terminal.
package terminal

import (
	"bufio"
	"os"

	"github.com/Shosta/androSecTest/logging"
)

// Waitfor a user input on the CLI.
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
