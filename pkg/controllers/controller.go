package controllers

import (
	"encoding/json"
	"net/http"
)

type responseBody struct {
	Msg  string
	Code string
	Data interface{}
}

func formatResponseBody(w http.ResponseWriter, body responseBody) {
	res, _ := json.Marshal(body)
	w.Write(res)
}
