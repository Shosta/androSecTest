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

// Package logging contains utility functions to work with logs.
package logging

import (
	"fmt"

	"github.com/Shosta/androSecTest/config"
)

// PrintlnDebug :
func PrintlnDebug(log string) {
	if config.IsDebug {
		fmt.Println(Orange("[dbg:info] ") + log)
	}
}

// PrintDebug :
func PrintDebug(log string) {
	if config.IsDebug {
		fmt.Print(Orange("[dbg:info] ") + log)
	}
}

// PrintlnError :
func PrintlnError(err interface{}) {
	if config.IsDebug {
		fmt.Println(Red("[error] ") + fmt.Sprint(err))
	}
}

// PrintlnVerbose : Print the log to the terminal if the configuration's "IsVerboseLogRequired" value is set to "true". 
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
