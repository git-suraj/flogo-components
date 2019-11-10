package getcache

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

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

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("region", "EU")
	tc.SetInput("tsc", "yBSD4qpCgX8hwbQA+5HvwkVYz2xb/cnAblRrRiHNh4Oh3SCBcU5VY0grYH4IlAsKO8DodomNGVCo6KpArWOoBr7IQjuD/wW5e8/QNQ1sS9/aYRQROaB0iLU/4dMcVS6PRWCN89hab1koKLn5cXjIkUGF2q/F1OJzENYDGY2nBlShhrY9eZWHdJEuNYSYud81b2P7r2nCW/ocWJb0PTBXwJXQIaq9Cw2e5jW2uyOSuTQQn4OFmfkFr1dc2iL9GhQSvVxw+1C/6lwYzjDJxj9g9BzxrrLwMgq7vjRWr4ug20x7PtKKYIJxWrStI0d3PW8uNbg3qj20V9dbvC1uPPOBv8yKk71sIDVKfvo+slgQOhQt8FtE182Yxd7pv+0KLvwfGoUWcYRUSIDZZXfF6/bn/Hr6xiw7sihYNaCrN47EBpBo0nodPEjJfatjkbf0Gjb9G26OJVlX0iTditrxYVoCDq80Mjzp308yD5nIm9txptNkwjGdaVnbzTLT/1vHe6qzy8pUG+aM+EqwJz0lUv6A1S3K6HokvtUYy4tFQAbdGK6t1EyXj2LW84TSUCQCRNuptn/eFdjHy4BrR8OUqUhPIe/UgituHXKRdeDuEZnR+SySONcpdr2WDXeEiKNwxKljcT1KVxR4xyqUts75q/oPicYsg0ZeWjxTUVzaV3oHOyI0I2m3DIlGX9xqc1TihYZqmW9KJhUAWMqKRFELc2W6Obs/kAQcAYZTlUXcaBEgFmSCPT02KwqNR28VOLlrgIHNjJh9EvWg3JryMZX5S/Wjelvj8Ez5p56Kh0WCLl7TOrePYuBhd95XbBHrzWXma/ELuKKdpSfbL4f9Ze8DFVAcbUTDFy6AnmnmIZ7IeWTInGirZspAU+OzjCvh4UNdH+U5f7PU6rRD8o3fqMTkuC+xPOwz8C2by2+g9OHtMOj/vDbT1L13Q7W6UtUPKLOYqVqfCzQMk7XI7HpkyNsjw2fcOkgVm2UCd9ZnHfhC5qQmQ5N/953+0uDkHEoGVkJNHL7m")
	tc.SetInput("domain", "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")
	tc.SetInput("id", "21704")

	//setup attrs

	act.Eval(tc)

	//check result attr
	// output := tc.GetOutput("result")
}

type ConfigProperties map[string]string

var gprops ConfigProperties

func ReadPropertiesFile(filepath string) (ConfigProperties, error) {
	config := ConfigProperties{}

	if len(filepath) == 0 {
		return config, nil
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
