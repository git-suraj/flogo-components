package emsv1

/*
#cgo darwin CFLAGS: -I.
#cgo darwin CFLAGS: -I/home/include/tibems
#cgo darwin LDFLAGS: -L/home/lib -ltibems64

#include <tibems.h>
tibemsDestination castToDestination(tibemsTemporaryQueue queue) {
  return (tibemsDestination)queue;
}
tibems_bool castToBool(int value) {
	return (tibems_bool)value;
}
tibems_long castToLong(int value) {
  return (tibems_long)value;
}
tibems_int castToInt(int value) {
  return (tibems_int)value;
}

 */
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


// Eval implements activity.Activity.Eval
func (a *CounterActivity) Eval(context activity.Context) (bool, error) {
	log.Debug("**********************************")
	log.Debugf("***************** start")
	a1 := context.GetInput("array1").(string)
	a2 := context.GetInput("array2").(string)
	log.Debugf("***************** a1 %v ", a1)
	log.Debugf("***************** a2 %v ", a2)
	context.SetOutput("output", "arrayop")
	log.Debugf("***************** end")
	return true, nil
}
