package appendarraysv1

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "tibco-append-arrays",
    "type": "flogo:activity",
    "ref": "github.com/git-suraj/flogo-components/activity/appendarraysv1",
    "version": "0.1.0",
    "title": "Append Arrays",
    "description": "append two arrays",
    "author": "suraj",
    "homepage": "https://github.com/git-suraj/flogo-components/tree/master/activity/appendarraysv1",
    "inputs": [
    {
      "name": "array1",
      "type": "array",
      "required": true
    },
    {
      "name": "array2",
      "type": "array",
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
