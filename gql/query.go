package gql

import "fmt"

//GetQueriedJSON accepts raw query string and retuns quried data from result json
func GetQueriedJSON(rawQuery string) {
	fmt.Println(segregate(rawQuery))
}

func queryJSON(schema, data map[string]interface{}) {

}
