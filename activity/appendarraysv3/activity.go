package appendarraysv3

import (
	"encoding/json"
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-tibco-append-arrays")

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

type request struct {
	name  string
	value string
}

// Eval implements activity.Activity.Eval
func (a *CounterActivity) Eval(context activity.Context) (bool, error) {
	log.Debug("======================")
	a1 := context.GetInput("array1").(string)
	a2 := context.GetInput("array2").(string)
	log.Debug("++++++++++++++++++++++")
	var array1 []string
	json.Unmarshal([]byte(a1), &array1)
	log.Debug("----------------------")
	var array2 []string
	json.Unmarshal([]byte(a2), &array2)
	arrayop := append(array1, array2...)
	log.Debug("**********************")
	context.SetOutput("output", arrayop)
	log.Debug("00000000000000000000000")
	return true, nil
}
