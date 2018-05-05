package serilize

import (
	"encoding/xml"
)

func XML(val interface{}) ([]byte, error) {
	var b []byte
	var err error

	if Pretty {
		b, err = xml.MarshalIndent(val, "", "  ")
	} else {
		b, err = xml.Marshal(val)
	}

	if err != nil {
		return []byte{}, nil
	}
	return b, nil
}
