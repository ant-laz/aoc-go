/*
   Copyright 2025 Google LLC

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

// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
package main

import (
	// Import the popular fmt package, which contains functions for formatting text, including printing to the console.
	// https://pkg.go.dev/fmt
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/std
	"fmt"
)

// the main function, the entry point for the program
// on the terminal running go run . will start execution here
func main() {
	// PrintLn function from package fmt
	// https://pkg.go.dev/fmt#Println
	fmt.Println("Solution is working")
}
