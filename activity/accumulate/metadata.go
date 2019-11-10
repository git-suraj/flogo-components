package accumulate

import (
	"github.com/project-flogo/core/data/coerce"
)

type Input struct {
	Type string                 `md:"type"`
	Key string                 `md:"key"`
	Schema     map[string]interface{} `md:"item"`
}

type Output struct {
	Schema     []interface{} `md:"output"`
}

func (o *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":   o.Type,
		"item":       o.Schema,
		"key" : o.Key,
	}
}

func (o *Input) FromMap(values map[string]interface{}) error {
	var err error
	o.Type, err = coerce.ToString(values["type"])
	if err != nil {
		return err
	}
	o.Key, err = coerce.ToString(values["key"])
	if err != nil {
		return err
	}
	o.Schema, err = coerce.ToObject(values["item"])
	if err != nil {
		return err
	}

	return nil
}


func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Schema,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	var err error

	o.Schema, err = coerce.ToArray(values["output"])
	if err != nil {
		return err
	}

	return nil
}

