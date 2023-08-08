package kql

import (
	"khasa/src/utils/slices"
	"khasa/src/utils/strings"
)

// General commands and their expected arguments
var commands = []string{
	"SET",
	"GET",
}

var expectedArguments = map[string][]int{
	"SET": []int{2, -1},
	"GET": []int{1, 1},
}

func IsCommandValid(command string) bool {
	return slices.Contains(commands, command)
}

func GetExpectedArguments(command string) map[string]int {
	return map[string]int{
		"min": expectedArguments[command][0],
		"max": expectedArguments[command][1],
	}
}

// Set command
var setCommandOptions = []string{
	"TYPE",
	"TTL",
}

var setCommandOptionValues = map[string][][]string{
	"TYPE": [][]string{
		[]string{"STRING", "EXACT"},
		[]string{"INT", "EXACT"},
		[]string{"FLOAT", "EXACT"},
		[]string{"BOOL", "EXACT"},
	},
	"TTL": [][]string{
		[]string{"INT", "TYPE"},
	},
}

func IsSetCommandOptionValid(option string) bool {
	return slices.Contains(setCommandOptions, option)
}

func IsSetCommandOptionValueValid(option string, value string) bool {
	for _, v := range setCommandOptionValues[option] {
		t := v[0]
		condition := v[1]

		if condition == "EXACT" && value == t {
			return true
		}

		if condition == "TYPE" {
			switch t {
			case "STRING":
				return true
			case "INT":
				if strings.IsValidInt(value) {
					return true
				}
			}
		}
	}

	return false
}
