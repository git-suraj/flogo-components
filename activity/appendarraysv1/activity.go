package appendarraysv1

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

// Request the request format
type Request struct {
	Arr []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"arr"`
}

// Eval implements activity.Activity.Eval
func (a *CounterActivity) Eval(context activity.Context) (bool, error) {

	a1 := context.GetInput("array1").(string)
	a2 := context.GetInput("array2").(string)
	array1 := Request{}
	json.Unmarshal([]byte(a1), &array1)
	array2 := Request{}
	json.Unmarshal([]byte(a2), &array2)
	arrayop := append(array1.Arr, array2.Arr...)
	res := Request{}
	res.Arr = arrayop
	b, err := json.Marshal(res)
	if err != nil {
		return false, err
	}
	context.SetOutput("output", string(b))
	return true, nil
}
