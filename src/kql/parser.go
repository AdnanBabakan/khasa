package kql

import (
	"fmt"
	"strings"
)

func TokenizeInput(input string) (map[string]interface{}, error) {
	// split by space except if it's in quotes
	quoted := false

	split := strings.FieldsFunc(input, func(r rune) bool {
		if r == '\'' || r == '"' {
			quoted = !quoted
		}
		return !quoted && r == ' '
	})

	command := split[0]

	if !IsCommandValid(command) {
		return nil, InvalidCommandError(command)
	}

	arguments := split[1:]

	if len(arguments) < GetExpectedArguments(command)["min"] {
		return nil, NotEnoughArgumentsError(command, GetExpectedArguments(command)["min"], len(arguments))
	}

	if len(arguments) > GetExpectedArguments(command)["max"] && GetExpectedArguments(command)["max"] != -1 {
		return nil, TooManyArgumentsError(command, GetExpectedArguments(command)["max"], len(arguments))
	}

	fmt.Println(SetCommandArgumentParser(arguments))

	return map[string]interface{}{
		"command":   command,
		"arguments": arguments,
	}, nil
}
