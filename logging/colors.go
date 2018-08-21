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

// Package logging : Provides printing on the terminal using several colors.
package logging

import "github.com/shosta/androSecTest/variables"

// Orange  : 
func Orange(str string) string {
	return variables.Orange + str + variables.Endc
}

// Green : 
func Green(str string) string {
	return variables.Green + str + variables.Endc
}

// Red : 
func Red(str string) string {
	return variables.Red + str + variables.Endc
}

// Blue : 
func Blue(str string) string {
	return variables.Blue + str + variables.Endc
}

// Bold :
func Bold(str string) string {
	return variables.Bold + str + variables.Endc
}
