package main

import (
	"fmt"
	"net/http"

	"github.com/sivssdn/mockgql/gql"
)

func main() {
	http.HandleFunc("/graphql", queryResolver)
	go http.ListenAndServe(":3000", nil)
	fmt.Println("Server started at port 3000 🚀🚀")
	<-make(chan int)
}

func queryResolver(res http.ResponseWriter, req *http.Request) {
	queries := [3]string{
		"{operationName:null,variables:{},query:fragment userFields on User {\\n  id\\n  emails {\\n    address\\n    verified\\n    __typename\\n  }\\n  username\\n  __typename\\n}\\n\\n{\\n  getUser {\\n    ...userFields\\n    __typename\\n  }\\n}\\n}",
		"{operationName:Me,variables:{},query:query Me {\\n  me {\\n    _id\\n    emails {\\n      address\\n      __typename\\n    }\\n    username\\n    roles\\n    name {\\n      initials\\n      full\\n      first\\n      last\\n      __typename\\n    }\\n    color\\n    area {\\n      _id\\n      name\\n      countryCode\\n      currencySymbol\\n      timezone\\n      __typename\\n    }\\n    status\\n    timezone\\n    __typename\\n  }\\n}\\n}",
		"{operationName:availabilitiesByArea,variables:{_id:144_1,onDate:mockOnDate_1},query:query availabilitiesByArea($areaId: ObjectId!, $weekAndYear: String, $onDates: [DateTime!]) {\\n  availabilitiesByArea(areaId: $areaId, weekAndYear: $weekAndYear, onDates: $onDates) {\\n    _id\\n    availability\\n    onDate\\n    isOnLateStandBy\\n    user {\\n      _id\\n      name {\\n        full\\n        initials\\n        __typename\\n      }\\n      status\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n}",
	}
	res.Header().Set("Content-Type", "application/json")
	result := gql.GetQueriedJSON(queries[2])
	res.Write(result)
}
