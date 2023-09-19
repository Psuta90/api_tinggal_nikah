package utils

import (
	"encoding/json"
)

func StructToMap(data interface{}) (map[string]interface{}, error) {

	nm := make(map[string]interface{})

	datastr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(datastr, &nm); err != nil {
		return nil, err
	}

	return nm, nil
}
