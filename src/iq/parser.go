package iq

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBeInteger = "must be an integer.\n"
	mustBeGreaterThan0 = "must be strictly greater than 0.\n"
	mustBePositive = "must be positive.\n"
	mustBeBetween = "must be an integer between 0 and 200.\n"
	smallerThanIq1 = "is smaller than IQ1.\n"
	tooLarge = "is too large.\n"
)

// U - mean
var U int
// S - standardDeviation
var S int
// IQ1 - minimum IQ
var IQ1 int
// IQ2 - maximum IQ
var IQ2 int

func printError(valueName string, errorMessage string) {
	fmt.Printf("Error: '%s' %s\n", valueName, errorMessage)
}

func isInteger(arg string) bool {
	var re = regexp.MustCompile("[-+]?\\d+")

	match := re.FindString(arg)
	if len(arg) != len(match) {
		return false
	}
	return true
}

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if (arg == "-h") {
			return true
		}
	}
	return false
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 2 {
		fmt.Println(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > 4 {
		fmt.Println(tooManyArgs)
		return false
	}
	valueNames := [4]string{"u", "s", "IQ1", "IQ2"}
	for i, arg := range argsWithoutProg {

		// Check integer
		if !isInteger(arg) {
			printError(valueNames[i], mustBeInteger)
			return false
		}
		integer, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		if integer < 0 {
			printError(valueNames[i], mustBePositive)
			return false
		}

		// Check and assign u
		if valueNames[i] == "u" {
			if (integer > 200) {
				printError(valueNames[i], tooLarge)
				return false
			}
			U = integer
		}

		// Check and assign s
		if valueNames[i] == "s" {
			if integer <= 0 {
				printError(valueNames[i], mustBeGreaterThan0)
				return false
			}
			S = integer
		}

		// Check and assign IQs
		if valueNames[i] == "IQ1" || valueNames[i] == "IQ2" {
			if integer < 0 || integer > 200 {
				printError(valueNames[i], mustBeBetween)
				return false
			}
		}
		if valueNames[i] == "IQ1" {
			IQ1 = integer
		}
		if valueNames[i] == "IQ2" {
			if integer < IQ1 {
				printError(valueNames[i], smallerThanIq1)
				return false
			}
			IQ2 = integer
		}
	}
	return true
}