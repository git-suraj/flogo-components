package appendcomplexarrays

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

// Eval implements activity.Activity.Eval
func (a *CounterActivity) Eval(context activity.Context) (bool, error) {
	log.Debug("started......")
	arrIP := context.GetInput("inputarray").([]interface{})
	log.Debugf("array ip %v ..........", arrIP)
	itemIP := context.GetInput("item").(map[string]interface{})
	log.Debugf("item ip %v ..........", itemIP)
	arrOP := append(arrIP, itemIP)
	context.SetOutput("output", arrOP)
	log.Debug("end......")
	return true, nil
}
