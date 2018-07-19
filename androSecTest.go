/*
Copyright 2018 Rémi Lavedrine.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/shosta/androSecTest/androidpkg"
	dependency "github.com/shosta/androSecTest/command/dependency"
	"github.com/shosta/androSecTest/variables"

	"github.com/alexflint/go-arg"
)

var args struct {
	Package string `arg:"-p" help:"package name"`
	Dest    string `arg:"-d" help:"destination folder absolute path"`
	Verbose bool   `arg:"-v" help:"verbosity level"`
}

func main() {
	arg.MustParse(&args)
	pkgname := ""
	if args.Package == "" {
		// Wait for the user input to get the package.
		pkgname = "com.orange.wifiorange"
	}

	if args.Verbose {
		variables.IsVerboseLogRequired = true
	}

	dependency.AreAllReady()

	pkgname := androidpkg.Savelocal(args.AppName)
	androidpkg.Setup(pkgname)
}
