package appendarraysv3

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "tibco-append-arrays-v3",
    "type": "flogo:activity",
    "ref": "github.com/git-suraj/flogo-components/activity/appendarraysv3",
    "version": "0.1.1",
    "title": "Append Arrays v3",
    "description": "append two arrays",
    "author": "suraj",
    "homepage": "https://github.com/git-suraj/flogo-components/tree/master/activity/appendarraysv3",
    "inputs": [
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
    "outputs": [
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
