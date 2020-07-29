package gql

import (
	"encoding/json"
	"fmt"
)

//GetQueriedJSON accepts raw query string and retuns quried data from result json
func GetQueriedJSON(rawQuery string) {
	// segregate(rawQuery)
	// fmt.Println(segregate(rawQuery))
	query := segregate(rawQuery)
	schema := query["schema"]
	// fmt.Println(schema)
	tmpData := `[{
        "    _id": "",
        "    availability": "",
        "    onDate": "",
        "    isOnLateStandBy": "",
        "    user ": {
            "      _id": "",
            "      name ": {
                "        full": "",
                "        initials": "",
                "        __typename": ""
            },
            "      roles": "",
            "      color": "",
            "      status": "",
            "      __typename": ""
        },
        "    __typename": ""
    }]`

	queryJSON(schema, tmpData)
}

func queryJSON(schema, inputData interface{}) {
	var data []map[string]interface{}
	json.Unmarshal([]byte(inputData.(string)), &data)
	// fmt.Println(data)
	fmt.Println(getKeys(data[0]))
}

