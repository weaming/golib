package response

import (
	"net/http"

	"github.com/weaming/golib/serilize"
)

func JSON(w http.ResponseWriter, val interface{}, code ...int) {
	b, err := serilize.JSON(val, 0)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	w.Write(b)
}
