package gql

import (
	"encoding/json"
	"fmt"
	"strings"
)

//GetQueriedJSON accepts raw query string and retuns quried data from result json
func GetQueriedJSON(rawQuery string) string {
	query := segregate(rawQuery)
	schema := query["schema"]

	tmpData := `[{
        "_id": "144_1",
        "    availability": "mockAavailabilityValue_1",
        "onDate": "mockOnDate_1",
        "isOnLateStandBy": "onisOnLateStandBy_1",
        "user ": {
            "      _id": "user_id_1",
            "      name ": {
                "        full": "name_full_1",
                "        initials": "intitials_1",
                "        __typename": "__typename_1"
            },
            "      roles": "roles_1",
            "      color": "color_1",
            "      status": "status_1",
            "      __typename": "__typename_1"
        },
        "    __typename": "__typename_outer_1"
    }]`

	result := queryJSON(schema, tmpData)
	resultString, _ := json.Marshal(result)
	return string(resultString)
}

func queryJSON(schema, inputData interface{}) interface{} {
	var dataArray, results []map[string]interface{}
	var result, data map[string]interface{}
	schemaObj := schema.(map[string]interface{})

	err := json.Unmarshal([]byte(inputData.(string)), &dataArray)
	if err == nil {
		//array json
		for _, value := range dataArray {
			result = diffJSON(schemaObj, value)
			results = append(results, result)
		}
		return results
	}
	//non-array json
	err = json.Unmarshal([]byte(inputData.(string)), &data)
	if err != nil {
		//malformed json
		fmt.Println("JSON data input is malformed")
	}
	return diffJSON(schemaObj, data)
}

func diffJSON(schemaObj, data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for mockKey, mockValue := range data {
		nestedSchema, ok := schemaObj[strings.ReplaceAll(mockKey, " ", "")]
		if ok == true {
			// result[mockKey] = mockValue
			switch mockValue.(type) {
			case string:
				result[mockKey] = mockValue
			case map[string]interface{}, []map[string]interface{}:
				jsonValueObj, err := json.Marshal(mockValue)
				if err != nil {
					result[mockKey] = "malformed json"
					continue
				}
				result[mockKey] = queryJSON(nestedSchema, string(jsonValueObj))
			}
		}
	}
	return result
}
