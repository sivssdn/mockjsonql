package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sivssdn/mockgql/gql"
)

func main() {
	http.HandleFunc("/graphql", queryResolver)
	go http.ListenAndServe(":4000", nil)
	log.Println("Server started at port 4000 🚀🚀")
	<-make(chan int)
}

//resolver logic for http requests
func queryResolver(res http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type")
	res.Header().Set("Content-Type", "application/json")
	if req.Method == "OPTIONS" {
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	result := gql.GetQueriedJSON(string(body))
	res.Write(result)
}
