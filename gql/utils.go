package gql

import (
	"strings"
)

func getSubstring(start, fullString, end string) string {
	startIndex := strings.Index(fullString, start)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(start)
	endIndex := strings.Index(fullString, end)
	if endIndex == -1 {
		return ""
	}
	return fullString[startIndex:endIndex]
}

//accepts an object and returns string array of it keys
func getKeys(object map[string]interface{}) []string {
	keys := make([]string, len(object))
	i := 0
	for index := range object {
		keys[i] = index
		i++
	}
	return keys
}