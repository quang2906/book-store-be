package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stripe/stripe-go"
)

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func GetEvent(w http.ResponseWriter, req *http.Request) (eventRes *stripe.Event, err error) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	event := stripe.Event{}
	err = json.Unmarshal(payload, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func CORSCheck(handler func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	res := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		handler(w, req)
	}
	return res
}
