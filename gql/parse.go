package gql

import "fmt"

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
	fmt.Println(query)
}
