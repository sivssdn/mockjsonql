package gql

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

//separates the string query into operationName, variable & query
func Segregate(fullQuery string) {
	operationName := getSubstring("operationName:", fullQuery, ",")
	variables := getSubstring("variables:{", fullQuery, "},")
	fmt.Println("operationName :", operationName)
	fmt.Println("variables :", variables)
	// fmt.Println(getSubstring("query:query", fullQuery, "}"))
}

func SegregateQuery(fullQuery string) {
	endDelimiter := "++,+"
	query := getSubstring("{\\n", fullQuery+endDelimiter, endDelimiter)
	query = getSubstring("{\\n", fullQuery+endDelimiter, endDelimiter)
	jsonQueryString := queryToJson(query)
	var jsonQuery interface{}
	json.Unmarshal([]byte(jsonQueryString), &jsonQuery)
	fmt.Println(jsonQuery)
}

func queryToJson(query string) string {
	var re = regexp.MustCompile(",[ ]*\"[ ]*}")
	jsonQuery := "{\"" + strings.ReplaceAll(query, "{\\n", "\":{\"")
	jsonQuery = strings.ReplaceAll(jsonQuery, "}\\n", "},\"")
	jsonQuery = strings.ReplaceAll(jsonQuery, "\\n", "\":\"\",\"")
	jsonQuery = re.ReplaceAllString(jsonQuery, "}")
	lastBracketIndex := strings.LastIndex(jsonQuery, "}")
	return jsonQuery[0:lastBracketIndex]
}
