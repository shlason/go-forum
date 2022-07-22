package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/utils"
)

type auth struct {
	Signup http.Handler
	Login  http.Handler
	Logout http.Handler
}

func signup(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)
	if user.Email == "" || user.Name == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: "params not enough", Data: nil})
		return
	}
	err := user.ReadByEmail()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		formatResponseBody(w, responseBody{Msg: "email already exist", Data: nil})
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s", err), Data: nil})
		return
	}
	err = user.ReadByName()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		formatResponseBody(w, responseBody{Msg: "user name already exist", Data: nil})
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s", err), Data: nil})
		return
	}
	err = user.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s", err), Data: nil})
		return
	}
	formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s", err), Data: user})
}

var accountTypes = map[string]string{
	"email": "email",
	"name":  "name",
}

func login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)
	if (user.Email == "" && user.Name == "") || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: "params not enough", Data: nil})
		return
	}

	var (
		err         error
		accountType string
		rpwd        string = user.Password
	)

	if user.Email != "" {
		accountType = accountTypes["email"]
	}
	if user.Name != "" {
		accountType = accountTypes["name"]
	}

	switch accountType {
	case accountTypes["eamil"]:
		err = user.ReadByEmail()
	case accountTypes["name"]:
		err = user.ReadByName()
	}
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s not found", accountType), Data: nil})
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s", err), Data: nil})
		return
	}
	if !utils.CheckPasswordHash(user.Password, rpwd) {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: "password incorrect", Data: nil})
		return
	}
	// TODO: session and set cookie
	formatResponseBody(w, responseBody{Msg: "success", Data: user})
}

func logout(w http.ResponseWriter, r *http.Request) {}

var Auth = auth{
	Signup: http.HandlerFunc(signup),
	Login:  http.HandlerFunc(login),
	Logout: http.HandlerFunc(logout),
}
