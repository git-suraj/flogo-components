package appendarraysv2

import (
	"encoding/json"
	"fmt"
	"testing"

	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestStringAppend(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	array1 := []interface{}{"a", "b"}
	array2 := []interface{}{"c", "d"}
	tc.SetInput("array1", array1)
	tc.SetInput("array2", array2)

	act.Eval(tc)

	value := tc.GetOutput("output").([]interface{})

	if len(value) != 4 {
		t.Fail()
	}
}

func TestIntAppend(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	array1 := []interface{}{1, 2}
	array2 := []interface{}{3, 4}
	tc.SetInput("array1", array1)
	tc.SetInput("array2", array2)

	act.Eval(tc)

	value := tc.GetOutput("output").([]interface{})

	if len(value) != 4 {
		t.Fail()
	}
}

type Req1 []struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func TestStructAppend(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	s1 := `[
			{
				"name":"1",
				"value":"a"
			}
		]`

	s2 := `[
			{
				"name":"2",
				"value":"b"
			}
		]`
	var a1 []interface{}
	var a2 []interface{}
	json.Unmarshal([]byte(s1), &a1)
	json.Unmarshal([]byte(s2), &a2)
	tc.SetInput("array1", a1)
	tc.SetInput("array2", a2)
	fmt.Println(act)
	act.Eval(tc)
	value := tc.GetOutput("output").([]interface{})
	if len(value) != 2 {
		t.Fail()
	}
}
