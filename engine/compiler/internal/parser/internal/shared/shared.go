package shared

import (
	"fmt"
)

const (
	KEY_NAME         string = "name"
	KEY_CATEGORY     string = "category"
	KEY_DESCRIPTION  string = "description"
	KEY_TYPE         string = "type"
	KEY_ARRAY        string = "array"
	KEY_PROPS        string = "props"
	KEY_REQUEST      string = "request"
	KEY_RESPONSES    string = "responses"
	KEY_HEADERS      string = "headers"
	KEY_MESSAGE_BODY string = "body"
	KEY_STATUS_CODE  string = "status_code"
)

func ParseComponentMetadata(fields map[string]any) (ParsedComponentMetadata, error) {
	metadata := ParsedComponentMetadata{}

	mapName, exists, err := GetStringFromMap(fields, KEY_NAME)
	if exists {
		if err != nil {
			return ParsedComponentMetadata{}, fmt.Errorf("error parsing %s: %+v", KEY_NAME, err)
		}

		metadata.Name = &mapName
	}

	category, _, err := GetStringFromMap(fields, KEY_CATEGORY)
	if err == nil {
		metadata.Category = &category
	}

	description, _, err := GetStringFromMap(fields, KEY_DESCRIPTION)
	if err == nil {
		metadata.Description = &description
	}

	return metadata, nil
}

func GetStringFromMap(mp map[string]any, key string) (string, bool, error) {
	if val, ok := mp[key]; ok {
		if str, ok := val.(string); ok {
			return str, true, nil
		}

		return "", true, fmt.Errorf("invalid type for %s: %+v", key, val)
	}

	return "", false, fmt.Errorf("missing %s", key)
}

type ParsedComponentMetadata struct {
	Name        *string `json:"name,omitempty" toml:"name,omitempty"`
	Category    *string `json:"category,omitempty" toml:"category,omitempty"`
	Description *string `json:"description,omitempty" toml:"description,omitempty"`
}
