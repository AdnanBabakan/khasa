package kql

import "strings"

func SetCommandArgumentParser(args []string) (map[string]interface{}, error) {
	options := args[2:]

	parsedOptions := make(map[string]string)

	for _, option := range options {
		split := strings.Split(option, "=")
		if !IsSetCommandOptionValid(split[0]) {
			//return
		}
		parsedOptions[split[0]] = split[1]
	}

	return map[string]interface{}{
		"key":     args[0],
		"value":   args[1],
		"options": parsedOptions,
	}, nil
}
