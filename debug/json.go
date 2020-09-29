package debug

import (
	"encoding/json"
	"log"
)

func LogJSON(prefix string, v interface{}) error {
	bin, e := json.Marshal(v)
	if e != nil {
		return e
	}
	log.Println(prefix, string(bin))
	return nil
}
