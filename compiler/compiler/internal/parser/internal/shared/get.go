package shared

import "fmt"

// TODO: refactor out
func GetStringFromMap(mp map[string]any, key string) (string, bool, error) {
	if val, ok := mp[key]; ok {
		if str, ok := val.(string); ok {
			return str, true, nil
		}

		return "", true, fmt.Errorf("invalid type for %s: %+v", key, val)
	}

	return "", false, fmt.Errorf("missing %s", key)
}
