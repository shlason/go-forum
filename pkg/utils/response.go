package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Msg  string      `json:"msg"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func FormatResponseBody(w http.ResponseWriter, body ResponseBody) {
	res, _ := json.Marshal(body)
	w.Write(res)
}
