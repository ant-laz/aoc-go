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
	// the fmt package, which contains functions for formatting text, including printing to the console.
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/fmt
	"fmt"
	"strings"

	// the bufio package is used to read text file
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/bufio#Scanner
	"bufio"
	// the os package is used to open files on operating system
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/os
	"os"
	// the strconv package is used to convert from str to other types, e.g. int
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/
	"strconv"
	// the slices package is used to sort slices
	// This package is one of the standard library packages you got when you installed Go.
	// https://pkg.go.dev/slices
	"slices"
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

	// Creating some Go int slices to hold the left & right list of numbers
	// We know the capacity needs to be 1,000 because that's the number
	/// of lines in the input text file
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	// Read the file one line at a time
	// separate columnn entries, but splitting each line into two
	// convert column entries from string to int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		listentries := strings.Split(line, "   ")
		l, _ := strconv.Atoi(listentries[0])
		left = append(left, l)
		r, _ := strconv.Atoi(listentries[1])
		right = append(right, r)
	}

	// sort the right & left values
	slices.Sort(left)
	slices.Sort(right)

	// compute & sum the pairwise differences
	diffsum := 0
	for idx, _ := range left {
		diff := 0
		leftval := left[idx]
		rightval := right[idx]
		if leftval >= rightval {
			diff = leftval - rightval
		} else {
			diff = rightval - leftval
		}
		diffsum += diff
		fmt.Println("idx = ", idx, "sum = ", diffsum, "left[idx] = ", leftval, "right[idx] = ", rightval)
	}

	// return the sum

	fmt.Println("the answer is = ", diffsum)

}

// utilty function to help handle errors that could occur
// error is a built-in type in Go
// panic is a built-in function in Go
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
