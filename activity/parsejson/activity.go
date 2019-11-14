package parsejson

import (
	"encoding/json"
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("parse-json")

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
	log.Info("started......")
	ip1 := context.GetInput("input").(string)
	log.Infof("ip1 %v ..........", ip1)
	ip := "{\"d\":\"f\"}"
	op := p{d: "Sean"}
	log.Infof("ip %v ..........", ip)
	in := []byte(ip)
	jsonMap := make(map[string]interface{})
	json.Unmarshal(in, &jsonMap)
	log.Infof("raw %v ..........", jsonMap)
	log.Infof("op %v ..........", op)
	context.SetOutput("output", jsonMap)
	log.Info("end......")
	return true, nil
}
