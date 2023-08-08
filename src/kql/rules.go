package kql

import "khasa/src/utils/slices"

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

var setCommandOptionValues = map[string][]interface{}{
	"TYPE": []interface{}{"EXACT:STRING", "EXACT:INT", "EXACT:FLOAT", "EXACT:BOOL"},
	"TTL":  []interface{}{"TYPE:INT"},
}

func IsSetCommandOptionValid(option string) bool {
	return slices.Contains(setCommandOptions, option)
}
