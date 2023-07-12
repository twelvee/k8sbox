// Package utils is a useful utils that boxie use. Methods are usually exported
package utils

import (
	"encoding/json"
)

// StructToMap Converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &newMap)
	return newMap, err
}
