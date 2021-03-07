package sdk

import (
	"encoding/json"
)

func (sdk SDK) Unmarshal(out interface{}) error {
	var varMap = map[string]interface{}{}
	for _, variable := range sdk.TelemetryVars.Vars {
		varMap[variable.Name] = variable.Value
	}
	marshal, err := json.Marshal(varMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(marshal, out)
}
