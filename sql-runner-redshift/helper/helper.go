package helper

import (
	"encoding/json"
)

// InArrayStr function to check if a string element is present in an array of string
func InArrayStr(array []string, testElement string) bool {
	for _, s := range array {
		if testElement == s {
			return true
		}
	}
	return false
}

// SQLParametersParser function to unmarshal SQL formatting parameters
func SQLParametersParser(input string) (map[string]interface{}, error) {
	var output map[string]interface{}
	err := json.Unmarshal([]byte(input), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
