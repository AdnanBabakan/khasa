package kql

import (
	"errors"
	"fmt"
)

func InvalidCommandError(command string) error {
	return errors.New(fmt.Sprintf("Invalid command: %s", command))
}

func NotEnoughArgumentsError(command string, needed int, provided int) error {
	return errors.New(fmt.Sprintf("Not enough arguments for command: %s. Needed %d arguments, but %d provided.", command, needed, provided))
}

func TooManyArgumentsError(command string, needed int, provided int) error {
	return errors.New(fmt.Sprintf("Too many arguments for command: %s. Needed %d arguments, but %d provided.", command, needed, provided))
}

func InvalidOptionError(command string, options string) error {
	return errors.New(fmt.Sprintf("Invalid option for command: %s. Option %s does not exist.", command, options))
}

func InvalidOptionValueError(command string, option string, value string) error {
	return errors.New(fmt.Sprintf("Invalid value for option %s for command %s. Value %s is not valid.", option, command, value))
}
