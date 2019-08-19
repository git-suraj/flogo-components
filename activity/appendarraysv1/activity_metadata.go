package appendarraysv1

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "tibco-append-arrays",
    "type": "flogo:activity",
    "ref": "https://github.com/git-suraj/flogo-components/activity/appendarraysv1",
    "version": "0.0.1",
    "title": "Append Arrays v1",
    "description": "append two arrays",
    "homepage":"https://github.com/git-suraj/flogo-components/tree/master/activity/appendarraysv1",
    "settings":[
    ],
    "input":[
      {
        "name": "array1",
        "type": "string",
        "required": true
      },
      {
        "name": "array2",
        "type": "string",
        "required": true
      }
    ],
    "output": [
      {
        "name": "output",
        "type": "array",
        "required": true
      }
    ]
  }  
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
