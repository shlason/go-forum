package controllers

import (
	"database/sql"
	"encoding/json"
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
		return
	}
	err := user.ReadByEmail()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = user.ReadByName()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = user.Create()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(user)
	w.Write(res)
}

func login(w http.ResponseWriter, r *http.Request) {}

func logout(w http.ResponseWriter, r *http.Request) {}

var Auth = auth{
	Signup: http.HandlerFunc(signup),
	Login:  http.HandlerFunc(login),
	Logout: http.HandlerFunc(logout),
}
