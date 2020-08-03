package gql

import (
	"fmt"
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

func checkErr(err error, message string) {
	if err != nil {
		fmt.Println(message, "\n", err)
	}
}
