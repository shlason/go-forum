package structs

import (
	"encoding/json"
	"net/http"
)

type ResponseBody struct {
	Msg  string      `json:"msg"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func WriteResponseBody(w http.ResponseWriter, body ResponseBody) {
	res, _ := json.Marshal(body)
	w.Write(res)
}
