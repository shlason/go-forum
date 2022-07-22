package controllers

import (
	"encoding/json"
	"net/http"
)

type responseBody struct {
	msg  string
	code string
	data interface{}
}

func formatResponseBody(w http.ResponseWriter, body responseBody) {
	res, _ := json.Marshal(body)
	w.Write(res)
}
