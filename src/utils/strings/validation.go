package strings

import "unicode"

func IsValidInt(value string) bool {
	for _, c := range value {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsValidFloat(value string) bool {
	for _, c := range value {
		if !unicode.IsDigit(c) && c != '.' {
			return false
		}
	}
	return true
}

func IsValidBool(value string) bool {
	return value == "true" || value == "false" || value == "1" || value == "0"
}
