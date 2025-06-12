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
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/fmt
	"fmt"
	// Import bufio package to be used to read text file
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/bufio#Scanner
	"bufio"
	// Used to access files on operating system
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/os
	"os"
)

// the main function, the entry point for the program
// on the terminal running go run . will start execution here
func main() {
	// Open the file with the puzzle input
	// check for errors
	file, err := os.Open("input.txt")
	checkError(err)

	// We defer the closing of that file with closeFile.
	// This will be executed at the end of the enclosing function (main)
	// This ensure we do not "leak" connections to files
	defer file.Close()

	// Read the file one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

// utilty function to help handle errors that could occur
// error is a built-in type in Go
// panic is a built-in function in Go
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
