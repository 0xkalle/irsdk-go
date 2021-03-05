package sdk

import (
	"fmt"
	"reflect"
)

const tagName = "validate" // Todo rename
func (sdk SDK) Unmarshal(out interface{}) error {
	rv := reflect.ValueOf(out)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("argument invalid")
	}
	v := reflect.ValueOf(out)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	//json.Unmarshal()
	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)
		name := v.Type().Field(i).Name
		// Skip if tag is not defined or ignored
		if tag != "" && tag != "-" {
			name = tag
		}
		getVar, err := sdk.GetVar(name)
		if err != nil {
			return err
		}
		if getVar.Name == "" {
			continue
		}
		v.Field(i).Set(reflect.ValueOf(getVar.Value))

	}
	return nil
}
