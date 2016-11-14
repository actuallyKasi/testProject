package utils

import (
	"errors"
	"encoding/json"
)

func ObjectToJson(object interface{}) ([]byte, error) {
	ReturnObject, err := json.Marshal(object)
	if err != nil {
		return ReturnObject, errors.New("Marshalling the response failed!")
	}
	return ReturnObject, nil
}