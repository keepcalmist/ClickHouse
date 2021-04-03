package restServer

import (
	"encoding/json"
	"github.com/keepcalmist/grpcFibonacci/pkg/usefullFunctions"
	"log"
	"net/http"
)

type (
	fiboRequest struct {
		X int32 `json:"x"`
		Y int32 `json:"y"`
	}
	fiboResponce struct {
		Digits []int32 `json:"digits"`
	}
)

func (s *server) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("[", r.Method, "]", r.URL.Path)
	request := &fiboRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		usefullFunctions.Respond(w, r, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	if !usefullFunctions.Validate(request.X, request.Y) {
		usefullFunctions.Respond(w, r, http.StatusBadRequest, nil)
		log.Println(err)
		return
	}
	resp, err := s.fibo.CalculateDigits(r.Context(), request.X, request.Y)
	if err != nil {
		usefullFunctions.Respond(w, r, http.StatusBadRequest, err)
		log.Println(err)
		return
	}
	respFibo := &fiboResponce{Digits: resp}
	usefullFunctions.Respond(w, r, http.StatusOK, respFibo)
	return
}
