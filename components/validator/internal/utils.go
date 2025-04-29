package internal

import (
	"strings"
)

func isValid(str string) bool {
	return strings.TrimSpace(str) != ""
}
