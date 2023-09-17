package common

import "encoding/json"

// ToJSON to JSON string
func ToJSONString(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// ToJSON to JSON string
func ToJSON(data interface{}) []byte {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return jsonBytes
}
