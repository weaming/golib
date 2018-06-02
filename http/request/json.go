package request

import (
	"encoding/json"
	"io"
)

func DeserilizeJSON(body io.ReadCloser, p interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	err := decoder.Decode(p)
	return err
}
