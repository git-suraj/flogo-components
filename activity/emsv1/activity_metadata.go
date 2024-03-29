package emsv1

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/project-flogo/legacybridge"
)

var jsonMetadata = `{
  "name": "emsv1",
  "type": "flogo:activity",
  "ref": "github.com/git-suraj/flogo-components/activity/emsv1",
  "version": "0.0.1",
  "title": "Publish EMS message",
  "description": "Publish EMS message",
  "author": "Suraj <supillai@tibco.com>",
  "homepage": "https://github.com/git-suraj/flogo-components/tree/master/activity/emsv1",
  "display": {
    "description": "Publish EMS message",
    "category": "Utility",
    "uid": "suraj-ems",
    "visible": true
  },
  "inputs":[
    {
      "name": "content",
      "type": "string",
      "required": true
    },
    {
      "name": "destination",
      "type": "string",
      "required": true
    },
    {
      "name": "serverUrl",
      "type": "string",
      "required": true
    },
    {
      "name": "user",
      "type": "string",
      "required": true
    },
    {
      "name": "password",
      "type": "string",
      "required": false
    },
    {
      "name": "exchangeMode",
      "type": "string",
      "required": true,
      "allowed": ["send-only","send-receive"]
    },
    {
      "name": "deliveryDelay",
      "type": "integer",
      "required": true
    },
    {
      "name": "deliveryMode",
      "type": "string",
      "required": true,
      "allowed" : ["persistent","non_persistent","reliable"]
    },
    {
      "name": "expiration",
      "type": "integer",
      "required": true
    },
    {
      "name": "tracing",
      "type": "any",
      "required": false
    }
  ],
  "outputs": [
    {
      "name": "response",
      "type": "string"
    },
    {
      "name": "tracing",
      "type": "any"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	//activity.Register(NewActivity(md))
	legacybridge.RegisterLegacyActivity(NewActivity(md))
}
