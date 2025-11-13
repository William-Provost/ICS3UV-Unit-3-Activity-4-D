// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-13
// Fileoverview: This program tells you if a given year is a leap year.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var yearAsString string = ""
	var yearAsNumber int = 0

	var reader = bufio.NewReader(os.Stdin)

	// input year
	fmt.Print("Please enter a year: ")
	yearAsString, _ = reader.ReadString('\n')
	yearAsString = strings.TrimSpace(yearAsString)
	yearAsNumber, _ = strconv.Atoi(yearAsString)

	// check if leap year
	if (yearAsNumber%4 == 0 && yearAsNumber%100 != 0) || yearAsNumber%400 == 0 {
		fmt.Println(yearAsNumber, "is a leap year.")
	} else {
		fmt.Println(yearAsNumber, "is not a leap year.")
	}

	fmt.Println("\nDone.")
}