
fetch("http://localhost:3000/graphql", {
    "headers": {
      "accept": "*/*",
      "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
      "authorization": "",
      "content-type": "application/json",
      "sec-fetch-dest": "empty",
      "sec-fetch-mode": "cors",
      "sec-fetch-site": "same-origin",
      "cookie": "keystone.uid=s%3A5c0928f49869f20004b5fbd2%3ABxSwYBG1RcAHa6lJ7QOkHTdZSdk21ZsCTHwBPijfMGY.VGn3yO32M%2FOsBsmDLsTBiZkMU7unehmF1I0wQJRptkI; this.sid=s%3AZPqZQt7kxxcFGih0wiU6swQxfXykJ6fa.evZFtnc9QHFE9s6DUYL%2F1OboJqX21YAf2sbY6unpLSI"
    },
    "referrer": "http://localhost:3000/login?message=Session%20expired!%20Please%20login%20again.",
    "referrerPolicy": "no-referrer-when-downgrade",
    "body": "{\"operationName\":null,\"variables\":{},\"query\":\"fragment userFields on User {\\n  id\\n  emails {\\n    address\\n    verified\\n    __typename\\n  }\\n  username\\n  __typename\\n}\\n\\n{\\n  getUser {\\n    ...userFields\\n    __typename\\n  }\\n}\\n\"}",
    "method": "POST",
    "mode": "cors"
  });
  


fetch("http://localhost:3000/graphql", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
    "authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InRva2VuIjoiNTViMDJiMjU0NjQ5NzI0MDZjNGU2MzA5NzM1ZDU0NjE4YjY2NTExYzE0NjIyMjlhYzk5ZTMyMGQ2NjBhNjY0MmEwYTUzODA4NDZmMjI1ODJiYjE5ODkiLCJpc0ltcGVyc29uYXRlZCI6ZmFsc2UsInVzZXJJZCI6IjVkODg5YjU3OTBkZTc3MDAwNDU0YjBmMiJ9LCJpYXQiOjE1OTU1MDk0MjMsImV4cCI6MTU5NTUxNDgyM30.KKv4WlO8ODpuFczFDu9wS6xED5u7_eK67ZvqqZcp154",
    "content-type": "application/json",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "cookie": "keystone.uid=s%3A5c0928f49869f20004b5fbd2%3ABxSwYBG1RcAHa6lJ7QOkHTdZSdk21ZsCTHwBPijfMGY.VGn3yO32M%2FOsBsmDLsTBiZkMU7unehmF1I0wQJRptkI; this.sid=s%3AZPqZQt7kxxcFGih0wiU6swQxfXykJ6fa.evZFtnc9QHFE9s6DUYL%2F1OboJqX21YAf2sbY6unpLSI; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InRva2VuIjoiNTViMDJiMjU0NjQ5NzI0MDZjNGU2MzA5NzM1ZDU0NjE4YjY2NTExYzE0NjIyMjlhYzk5ZTMyMGQ2NjBhNjY0MmEwYTUzODA4NDZmMjI1ODJiYjE5ODkiLCJpc0ltcGVyc29uYXRlZCI6ZmFsc2UsInVzZXJJZCI6IjVkODg5YjU3OTBkZTc3MDAwNDU0YjBmMiJ9LCJpYXQiOjE1OTU1MDk0MjMsImV4cCI6MTU5NTUxNDgyM30.KKv4WlO8ODpuFczFDu9wS6xED5u7_eK67ZvqqZcp154"
  },
  "referrer": "http://localhost:3000/login?message=Session%20expired!%20Please%20login%20again.",
  "referrerPolicy": "no-referrer-when-downgrade",
  "body": "{\"operationName\":\"Me\",\"variables\":{},\"query\":\"query Me {\\n  me {\\n    _id\\n    emails {\\n      address\\n      __typename\\n    }\\n    username\\n    roles\\n    name {\\n      initials\\n      full\\n      first\\n      last\\n      __typename\\n    }\\n    color\\n    area {\\n      _id\\n      name\\n      countryCode\\n      currencySymbol\\n      timezone\\n      __typename\\n    }\\n    status\\n    timezone\\n    __typename\\n  }\\n}\\n\"}",
  "method": "POST",
  "mode": "cors"
});


fetch("http://localhost:3000/graphql", {
  "headers": {
    "accept": "*/*",
    "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
    "authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InRva2VuIjoiNTViMDJiMjU0NjQ5NzI0MDZjNGU2MzA5NzM1ZDU0NjE4YjY2NTExYzE0NjIyMjlhYzk5ZTMyMGQ2NjBhNjY0MmEwYTUzODA4NDZmMjI1ODJiYjE5ODkiLCJpc0ltcGVyc29uYXRlZCI6ZmFsc2UsInVzZXJJZCI6IjVkODg5YjU3OTBkZTc3MDAwNDU0YjBmMiJ9LCJpYXQiOjE1OTU1MDk0MjMsImV4cCI6MTU5NTUxNDgyM30.KKv4WlO8ODpuFczFDu9wS6xED5u7_eK67ZvqqZcp154",
    "content-type": "application/json",
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "cookie": "keystone.uid=s%3A5c0928f49869f20004b5fbd2%3ABxSwYBG1RcAHa6lJ7QOkHTdZSdk21ZsCTHwBPijfMGY.VGn3yO32M%2FOsBsmDLsTBiZkMU7unehmF1I0wQJRptkI; this.sid=s%3AZPqZQt7kxxcFGih0wiU6swQxfXykJ6fa.evZFtnc9QHFE9s6DUYL%2F1OboJqX21YAf2sbY6unpLSI; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InRva2VuIjoiNTViMDJiMjU0NjQ5NzI0MDZjNGU2MzA5NzM1ZDU0NjE4YjY2NTExYzE0NjIyMjlhYzk5ZTMyMGQ2NjBhNjY0MmEwYTUzODA4NDZmMjI1ODJiYjE5ODkiLCJpc0ltcGVyc29uYXRlZCI6ZmFsc2UsInVzZXJJZCI6IjVkODg5YjU3OTBkZTc3MDAwNDU0YjBmMiJ9LCJpYXQiOjE1OTU1MDk0MjMsImV4cCI6MTU5NTUxNDgyM30.KKv4WlO8ODpuFczFDu9wS6xED5u7_eK67ZvqqZcp154"
  },
  "referrer": "http://localhost:3000/incidents",
  "referrerPolicy": "no-referrer-when-downgrade",
  "body": "{\"operationName\":\"availabilitiesByArea\",\"variables\":{\"areaId\":\"5d88f1b89259f80004fe9d51\",\"weekAndYear\":\"302020\"},\"query\":\"query availabilitiesByArea($areaId: ObjectId!, $weekAndYear: String, $onDates: [DateTime!]) {\\n  availabilitiesByArea(areaId: $areaId, weekAndYear: $weekAndYear, onDates: $onDates) {\\n    _id\\n    availability\\n    onDate\\n    isOnLateStandBy\\n    user {\\n      _id\\n      name {\\n        full\\n        initials\\n        __typename\\n      }\\n      roles\\n      color\\n      status\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\"}",
  "method": "POST",
  "mode": "cors"
});



