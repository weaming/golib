package serilize

import (
	"encoding/json"
	"strings"
)

func JSON(p interface{}, indent uint) ([]byte, error) {
	var b []byte
	var err error

	if indent > 0 {
		b, err = json.MarshalIndent(p, "", strings.Repeat(" ", int(indent)))
	} else {
		b, err = json.Marshal(p)
	}

	if err != nil {
		return nil, err
	}
	return b, nil
}
