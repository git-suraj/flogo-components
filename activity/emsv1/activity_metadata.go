package emsv1

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "emsv1",
  "type": "flogo:activity",
  "ref": "github.com/git-suraj/flogo-components/activity/emsv1",
  "version": "0.0.1",
  "title": "Publish EMS Message",
  "description": "Send message to EMS server",
  "homepage":"https://github.com/git-suraj/flogo-components/tree/master/activity/emsv1",
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
