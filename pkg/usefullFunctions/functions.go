package usefullFunctions

import (
	"encoding/json"
	"net/http"
)

func Validate(from, to int32) bool {
	if from > to || to < 0 || from < 0 {
		return false
	} else {
		return true
	}
}

func Calculate(x1, x2 int, count int32) []int32 {
	reSlice := make([]int32, count)
	for i := 0; int32(i) < count; i++ {
		reSlice[i] = int32(x1 + x2)
		x1, x2 = x2, x2+x1
	}
	return reSlice
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}
