package gql

import (
	"encoding/json"
	"regexp"
	"strings"
)

//segregate separates the string query into operationName, variable & query
func segregate(fullQuery string) map[string]interface{} {
	operationName := getSubstring("operationName:", fullQuery, ",")
	variables := getSubstring("variables:{", fullQuery, "},")
	jsonSchema := segregateQuery(fullQuery)
	return map[string]interface{}{"operationName": operationName, "variables": variables, "schema": jsonSchema}
}

//segregateQuery converts the raw string to it's equivalent JSON
func segregateQuery(fullQuery string) interface{} {
	endDelimiter := "++,+"
	query := getSubstring("{\\n", fullQuery+endDelimiter, endDelimiter)
	query = getSubstring("{\\n", query+endDelimiter, endDelimiter)
	jsonQueryString := queryToJSONSchema(query)
	var jsonQuery interface{}
	json.Unmarshal([]byte(jsonQueryString), &jsonQuery)
	return jsonQuery
}

//queryToJSON receives string gql query and returns its json (string) equivalent
func queryToJSONSchema(query string) string {
	var re = regexp.MustCompile(",[ ]*\"[ ]*}")
	jsonQuery := "{\"" + strings.ReplaceAll(query, "{\\n", "\":{\"")
	jsonQuery = strings.ReplaceAll(jsonQuery, "}\\n", "},\"")
	jsonQuery = strings.ReplaceAll(jsonQuery, "\\n", "\":\"\",\"")
	jsonQuery = re.ReplaceAllString(jsonQuery, "}")
	lastBracketIndex := strings.LastIndex(jsonQuery, "}")
	return jsonQuery[0 : lastBracketIndex-1]
}
