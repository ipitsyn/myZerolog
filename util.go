package myzerolog

import (
	"fmt"
)

func prepareMessage(args []interface{}) string {
	if len(args) == 0 {
		return ""
	}

	return fmt.Sprint(args...)
}

func prepareFormattedMessage(format string, args []interface{}) string {
	if len(args) == 0 {
		return format
	}

	return fmt.Sprintf(format, args...)
}
