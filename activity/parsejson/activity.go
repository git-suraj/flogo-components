package parsejson

import (
	"encoding/json"
	"strings"
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
	aip := context.GetInput("input").(string)
	ip := strings.Replace(aip, "\\", "", -1)
	//dummyip := "{\"d\":\"f\"}"
	log.Infof("ip %v ..........", ip)
	//log.Infof("dummy ip %v ..........", dummyip)
	jsonMapIP := make(map[string]interface{})
	//jsonMapDummyIP := make(map[string]interface{})
	//json.Unmarshal([]byte(dummyip), &jsonMapDummyIP)
	json.Unmarshal([]byte(ip), &jsonMapIP)
	log.Infof("ip map %v ..........", jsonMapIP)
	//log.Infof("dummy map %v ..........", jsonMapDummyIP)
	context.SetOutput("output", jsonMapIP)
	log.Info("end......")
	return true, nil
}
