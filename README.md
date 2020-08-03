## Mock GraphQL with JSON

A fast and simple mock for your graphql front-end without writing any back-end resolvers and services.
All you need to do is name your resolver and paste it's equivalent data under 'name' & 'data' key in resolver.json.

#### How to use:
- Fill 'name' & 'data' fields with string and JSON inside resolvers.json
- Build the app using `go build`
- Run your build file
- Send requests from your client to 'localhost:4000/graphql'