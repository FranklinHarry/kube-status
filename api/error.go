package api

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

type ApiError struct {
	Error string `json:"error"`
}

func logAndRespondWithError(w http.ResponseWriter, code int, format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	glog.Error(err)
	glog.Flush()
	respondWithError(w, err, code)
}

//Returns the error in a json format
//if something wrong happens it will write the error on the response writer as a normal string
func respondWithError(w http.ResponseWriter, err error, code int) {
	apiError := ApiError{err.Error()}
	if res, err := json.Marshal(apiError); err != nil {
		http.Error(w, string(res), code)
	}
	http.Error(w, err.Error(), code)
}
