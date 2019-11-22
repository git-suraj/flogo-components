package addarray

import (
	"encoding/json"
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("add-array")

type p struct {
	val []float64
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
	inputData, _ := data.CoerceToString(context.GetInput("input").(*data.ComplexObject).Value)
	log.Infof("%v----------%T", inputData, inputData)
	inputBytes := []byte(inputData)
	o := make(map[string]interface{})
	json.Unmarshal(inputBytes, &o)
	o1 := o["val"].([]interface{})
	sum := 0.0
	for _, v := range o1 {
		sum += v.(float64)
	}
	log.Debugf("sum %v", sum)
	context.SetOutput("output", sum)
	return true, nil
}
