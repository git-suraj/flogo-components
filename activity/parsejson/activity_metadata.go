package appendarraysv3

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `
"title": "Parse JSON",
"name": "Parse-JSON",
"author": "Suraj",
"ref": "github.com/git-suraj/flogo-components/activity/parsejson",
"homepage":"https://github.com/git-suraj/flogo-components/tree/master/activity/parsejson",
"type": "flogo:activity",
"version": "1.1.0",
"display": {
  "visible": true,
  "description": "This activity parses a json string into object",
  "category": "Utility"
},
"feature": {
  "iterator": {
    "type": "iterator",
    "enabled": false
  }
},
"inputs": [{
    "name": "item",
    "type": "complex_object",
    "display": {
      "description": "JSON data or schema for the input",
      "type": "texteditor",
      "syntax": "json",
      "name": "Input Schema",
      "mappable": false
    }
  },
  {
    "name": "input",
    "type": "string",
    "required": true
  }
],
"outputs": [{
  "name": "output",
  "type": "string"
}]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}