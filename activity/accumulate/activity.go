/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package accumulate

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

var activityMd = activity.ToMetadata(&Input{}, &Output{})

func init() {
	_ = activity.Register(&MyActivity{}, New)
	jsonArrayMap = make( map[string][]interface{})
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	return &MyActivity{logger: log.ChildLogger(ctx.Logger(), "append")}, nil
}

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	logger log.Logger
}

func (*MyActivity) Metadata() *activity.Metadata {
	return activityMd
}

var jsonArrayMap map[string][]interface{}


func (appendAct *MyActivity) Eval(context activity.Context) (done bool, err error) {

	appendAct.logger.Debugf("Executing Append Activity")

	input := &Input{}

	err = context.GetInputObject(input)
	if err != nil {
		return false, err
	}

	key := input.Key

	inputType := input.Type

	switch inputType {
	case "Append":

		var jsonArray []interface{}

		jsonArray,ok := jsonArrayMap[key]

		if ok {
			jsonArray = append( jsonArray , input.Schema)
			jsonArrayMap[key] = jsonArray


		} else {
			jsonArray = append( jsonArray , input.Schema)
			jsonArrayMap[key] = jsonArray

		}

	case "Clear" :

		_ , ok := jsonArrayMap[key]

		if ok {
			delete(jsonArrayMap, key )
		}


	case "Get":

		output := &Output{}

		jsonArray,ok := jsonArrayMap[key]
		if ok {
			output.Schema = jsonArray
		}

		err = context.SetOutputObject(output)
		if err != nil {
			return false, err
		}

	}
	return true, nil
}
