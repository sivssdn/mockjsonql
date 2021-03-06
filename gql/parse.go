package gql

import (
	"encoding/json"
	"regexp"
	"strings"
)

//segregate separates the string query into operationName, variable & query
func segregate(fullQuery string) map[string]interface{} {
	fullQuery = strings.ReplaceAll(fullQuery, "\"", "")
	operationName := getSubstring("operationName:", fullQuery, ",")
	variables, jsonSchema := make(chan interface{}), make(chan interface{})
	go segregateVariables(fullQuery, variables)
	go segregateQuery(fullQuery, jsonSchema)
	return map[string]interface{}{"operationName": operationName, "variables": <-variables, "schema": <-jsonSchema}
}

//segregateQuery converts the raw string to it's equivalent JSON
func segregateQuery(fullQuery string, result chan interface{}) {
	endDelimiter := "++,+"
	query := getSubstring("{\\n", fullQuery+endDelimiter, endDelimiter)
	query = getSubstring("{\\n", query+endDelimiter, endDelimiter)
	jsonQueryString := queryToJSONSchema(query)
	var jsonQuery interface{}
	json.Unmarshal([]byte(jsonQueryString), &jsonQuery)
	result <- jsonQuery
}

//queryToJSON receives string gql query and returns its json (string) equivalent
func queryToJSONSchema(query string) string {
	var re = regexp.MustCompile(",[ ]*\"[ ]*}")
	jsonQuery := "{\"" + strings.ReplaceAll(query, "{\\n", "\":{\"")
	jsonQuery = strings.ReplaceAll(jsonQuery, "}\\n", "},\"")
	jsonQuery = strings.ReplaceAll(jsonQuery, "\\n", "\":\"\",\"")
	jsonQuery = re.ReplaceAllString(jsonQuery, "}")
	jsonQuery = strings.ReplaceAll(jsonQuery, " ", "")
	lastBracketIndex := strings.LastIndex(jsonQuery, "}")
	if lastBracketIndex == -1 {
		lastBracketIndex = 1
	}
	return jsonQuery[0 : lastBracketIndex-1]
}

//segregateVariables receives raw query string and returns key-value pair for the queried variables with their values
func segregateVariables(fullQuery string, result chan interface{}) {
	fullQuery = "{\"" + getSubstring("variables:{", fullQuery, "},") + "\"}"
	fullQuery = strings.ReplaceAll(fullQuery, ":", "\":\"")
	fullQuery = strings.ReplaceAll(fullQuery, ",", "\",\"")
	var variableObj map[string]string
	err := json.Unmarshal([]byte(fullQuery), &variableObj)
	checkErr(err, "Unable to read query variables")
	result <- variableObj
}
