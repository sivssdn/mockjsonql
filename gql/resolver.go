package gql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type resolver struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

/*
* reads resolvers mock data from file and returns mock data corresponding to resolver's name
* @input resolverName string - name of the resolver for which data is to be returned
* @returns string equivalent of json result data
 */
func getResolverData(queryResolverName string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	file, err := ioutil.ReadFile("resolvers.json")
	if err != nil {
		panic("Couldn't read/find data from resolvers.json")
	}
	var resolvers []resolver
	err = json.Unmarshal(file, &resolvers)
	if err != nil {
		panic("Couldn't read/find data from resolvers.json")
	}
	var result []byte
	for _, resolver := range resolvers {
		if strings.EqualFold(resolver.Name, queryResolverName) {
			result, _ = json.Marshal(resolver.Data)
			break
		}
	}
	return string(result)
}
