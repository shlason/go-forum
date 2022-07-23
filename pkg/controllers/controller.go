package controllers

import (
	"encoding/json"
	"net/http"
)

type responseBody struct {
	Msg  string      `json:"msg"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func formatResponseBody(w http.ResponseWriter, body responseBody) {
	res, _ := json.Marshal(body)
	w.Write(res)
}
