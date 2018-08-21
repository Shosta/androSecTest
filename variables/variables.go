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

// Package variables : has all the global variables that are used in the program.
package variables

// Attacks folder names to have them gathered in one single place.
const (
	AttacksDir            = "/attacks"
	SourcePackageDir      = "/sourcePackage"
	UnzippedPackageDir    = "/unzippedPackage"
	DisassemblePackageDir = "/disassemblePackage"
	DecompiledPackageDir  = "/decompiledPackage"
	LeakagesDir           = "/leakages"
	DebuggablePackageDir  = "/debuggablePackage"
	InsecureBackupDir     = "/insecureBackup"
	InsecureLoggingDir    = "/insecureLogging"
	InsecureStorageDir    = "/insecureStorage"
)

// Color const to display color on the terminal command.
const (
	Header    = "\033[95m"
	Blue      = "\033[94m"
	Green     = "\033[92m"
	Orange    = "\033[93m"
	Red       = "\033[91m"
	Endc      = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
)
