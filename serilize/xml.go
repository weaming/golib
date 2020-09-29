package serilize

import (
	"encoding/xml"
	"strings"
)

func XML(val interface{}, indent uint) ([]byte, error) {
	var b []byte
	var err error

	if Pretty {
		b, err = xml.MarshalIndent(val, "", strings.Repeat(" ", int(indent)))
	} else {
		b, err = xml.Marshal(val)
	}

	if err != nil {
		return []byte{}, nil
	}
	return b, nil
}
