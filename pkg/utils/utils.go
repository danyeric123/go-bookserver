package utils

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func SendJsonResponse(w http.ResponseWriter, res []byte){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}