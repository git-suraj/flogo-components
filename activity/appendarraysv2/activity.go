package appendarraysv2

import (
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

	array1 := context.GetInput("array1").([]interface{})
	log.Debug("Incoming array1: ")
	log.Debug(array1)

	array2 := context.GetInput("array2").([]interface{})
	log.Debug("Incoming array2: ")
	log.Debug(array2)

	arrayop := append(array1, array2...)
	context.SetOutput("output", arrayop)
	return true, nil
}
