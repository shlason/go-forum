package controllers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/shlason/go-forum/pkg/constants"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
)

func handleInternalErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: fmt.Sprintf("%s\n%s", err, debug.Stack()), Data: nil})
}

func getSession(r *http.Request) (models.Session, error) {
	c, _ := r.Cookie(constants.Cookie.SessionTokenName)
	session := models.Session{
		UUID: c.Value,
	}
	err := session.ReadByUUID()
	return session, err
}
