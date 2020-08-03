package gql

import (
	"encoding/json"
	"fmt"
	"strings"
)

//GetQueriedJSON accepts raw query string and retuns quried data from result json
func GetQueriedJSON(rawQuery string) []byte {
	query := segregate(rawQuery)
	schema := query["schema"]
	resolverData := getResolverData(query["operationName"].(string))
	result := queryJSON(schema, resolverData, query["variables"])
	resultString, _ := json.Marshal(result)
	return resultString
}

//accepts schema and data from mock resolver and returns filtered data based on schema
func queryJSON(schema, inputData, filterVariables interface{}) interface{} {
	var dataArray, results []map[string]interface{}
	var result, data map[string]interface{}
	schemaObj := schema.(map[string]interface{})
	err := json.Unmarshal([]byte(inputData.(string)), &dataArray)
	if err == nil {
		//array json
		for _, value := range dataArray {
			result = filterJSON(schemaObj, value, filterVariables)
			results = append(results, result)
		}
		return results
	}
	//non-array json
	err = json.Unmarshal([]byte(inputData.(string)), &data)
	if err != nil {
		//malformed/empty json
		fmt.Println("JSON data input from resolvers.json is empty or malformed")
		return make(map[string]interface{}, 1) //returning empty object
	}
	return filterJSON(schemaObj, data, filterVariables)
}

//filters mock keys based on schema for a mockData object
func filterJSON(schemaObj, data map[string]interface{}, filterVariables interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if !isObjValid(data, filterVariables) {
		return result
	}
	for mockKey, mockValue := range data {
		nestedSchema, exists := schemaObj[strings.ReplaceAll(mockKey, " ", "")]
		if exists {
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
				result[mockKey] = queryJSON(nestedSchema, string(jsonValueObj), make(map[string]string))
			}
		}
	}
	return result
}

//checks if the query variables match the mock data object
func isObjValid(data map[string]interface{}, variables interface{}) bool {
	queryVariables := variables.(map[string]string)
	if len(queryVariables) < 1 {
		return true
	}
	isValid := 1
	var stringFieldValue string
	for variable, variableValue := range queryVariables {
		fieldValue, exists := data[variable]
		switch fieldValue.(type) {
		case string:
			stringFieldValue = fieldValue.(string)
		default:
			stringFieldValue = fmt.Sprint(fieldValue)
		}
		if exists && strings.EqualFold(stringFieldValue, variableValue) {
			isValid = isValid * 1
			continue
		}
		isValid = 0
	}
	if isValid == 1 {
		return true
	}
	return false
}
