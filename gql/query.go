package gql

import (
	"encoding/json"
)

//GetQueriedJSON accepts raw query string and retuns quried data from result json
func GetQueriedJSON(rawQuery string) {
	// segregate(rawQuery)
	// fmt.Println(segregate(rawQuery))
	query := segregate(rawQuery)
	schema := query["schema"]

	tmpData := `[{
        "    _id": "144_1",
        "    availability": "mockAavailabilityValue_1",
        "    onDate": "mockOnDate_1",
        "    isOnLateStandBy": "onisOnLateStandBy_1",
        "    user ": {
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

	queryJSON(schema, tmpData)
}

func queryJSON(schema, inputData interface{}) {
	var data, results []map[string]interface{}
	var result map[string]interface{}
	schemaObj := schema.(map[string]interface{})
	json.Unmarshal([]byte(inputData.(string)), &data)

	for _, value := range data {
		result = make(map[string]interface{})
		for mockKey, mockValue := range value {
			_, ok := schemaObj[mockKey]
			if ok == true {
				result[mockKey] = mockValue
			}
		}
		results = append(results, result)
	}

}
