package getcache

import (
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/koding/cache"
)

const (
	ivId  = "id"
)

var activityLog = logger.GetLogger("getcache")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	activityLog.Info("Start execution get cache....")

	if context.GetInput(ivId) == nil {
		activityLog.Error("get-cache-4001: id is not configured")
		return false, activity.NewError("id is not configured", "get-cache-4001", nil)
	}
	id := context.GetInput(ivId).(string)
	cache :=cache.N
	activityLog.Debug(fmt.Sprintf("Call getworkitems with region: %s, tsc,: %s, domain: %s, sandbox: %s and id: %s", region, tsc, domain, sandbox, id))
	response, err := work.GetWorkItems(region, tsc, domain, sandbox, id)

	if err != nil {
		activityLog.Error("get-workitems-5001: Error calling getworkitems")
		return false, activity.NewError("Error calling getworkitems", "get-workitems-5001", nil)
	}

	// Set response
	outputField := &data.ComplexObject{Value: response}
	context.SetOutput(ovResult, outputField)

	activityLog.Info("Finished execution LiveApps Get WorkItems Activity")

	return true, nil
}
