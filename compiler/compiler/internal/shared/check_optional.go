package shared

import "strings"

func CheckOptionalString(src *string) (string, bool) {
	if src == nil {
		return "", false
	}
	if strings.TrimSpace(*src) == "" {
		return "", false
	}

	return *src, true
}
