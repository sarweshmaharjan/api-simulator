package utils

import "encoding/json"

// ToJSON to JSON string
func ToJSON(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
