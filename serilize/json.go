package serilize

import "encoding/json"

func JSON(p interface{}) ([]byte, error) {
	var b []byte
	var err error

	if Pretty {
		b, err = json.MarshalIndent(p, "", "  ")
	} else {
		b, err = json.Marshal(p)
	}

	if err != nil {
		return []byte{}, nil
	}
	return b, nil
}
