package controllers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/shlason/go-forum/pkg/structs"
)

func handleInternalErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: fmt.Sprintf("%s\n%s", err, debug.Stack()), Data: nil})
}
