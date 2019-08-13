package appendarrays

import (
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
	array1 := []string{"a", "b"}
	array2 := []string{"c", "d"}
	tc.SetInput("array1", array1)
	tc.SetInput("array2", array2)

	act.Eval(tc)

	value := tc.GetOutput("output").([]int)

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
	array1 := []int{1, 2}
	array2 := []int{3, 4}
	tc.SetInput("array1", array1)
	tc.SetInput("array2", array2)

	act.Eval(tc)

	value := tc.GetOutput("output").([]string)

	if len(value) != 4 {
		t.Fail()
	}
}
