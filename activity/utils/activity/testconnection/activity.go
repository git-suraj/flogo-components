package testconnection

import (
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("append-complex-arrays")

type p struct {
	d string
}

// CounterActivity is a Counter Activity implementation
type CounterActivity struct {
	sync.Mutex
	metadata *activity.Metadata
}

// NewActivity creates a new CounterActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CounterActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *CounterActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// OP the output struct
type OP struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Region   string `json:"region"`
}

// Eval implements activity.Activity.Eval
func (a *CounterActivity) Eval(context activity.Context) (bool, error) {
	log.Debug("started......")
	var username, password, region string
	testConn := context.GetInput("testConnection").(map[string]interface{})
	connectionSettings := testConn["settings"].([]interface{})
	for _, v := range connectionSettings {
		setting := v.(map[string]interface{})
		if setting["name"] == "username" {
			username = setting["value"].(string)
		} else if setting["name"] == "password" {
			password = setting["value"].(string)
		} else if setting["name"] == "region" {
			region = setting["value"].(string)
		}
	}
	o := OP{Username: username, Password: password, Region: region}
	log.Debugf("op %v", o)
	context.SetOutput("user", username)
	context.SetOutput("password", password)
	context.SetOutput("region", region)
	log.Debug("end......")
	return true, nil
}
