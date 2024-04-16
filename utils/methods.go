package utils

import "strings"

func CheckString(text *string) bool {
	if text == nil || len(strings.TrimSpace(*text)) == 0 {
		return false
	}
	return true
}
