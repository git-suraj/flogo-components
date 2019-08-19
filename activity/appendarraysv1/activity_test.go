package appendarraysv1

import (
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

func TestStructAppend(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	/*array1 := []request{
		request{"a1n1", "a1v1"},
		request{"a1n2", "a1v2"},
	}
	array2 := []request{
		request{"a2n1", "a2v1"},
		request{"a2n2", "a2v2"},
	}
	tc.SetInput("array1", array1)
	tc.SetInput("array2", array2)

	act.Eval(tc)
	value := tc.GetOutput("output").([]request)
	object := reflect.ValueOf(value)
	if object.Len() != 4 {
		t.Fail()
	}*/

	a1 := `{
		"arr":[
			{
				"name":"1",
				"value":"a"
			}
		]
	}`

	a2 := `{
		"arr":[
			{
				"name":"1",
				"value":"a"
			}
		]
	}`
	tc.SetInput("array1", a1)
	tc.SetInput("array2", a2)
	act.Eval(tc)
	value := tc.GetOutput("output").(string)
	fmt.Println(value)
	if len(value) < 4 {
		t.Fail()
	}

}
