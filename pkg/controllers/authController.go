package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
	"github.com/shlason/go-forum/pkg/constants"
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
	err := utils.ParseBody(r, user)

	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: "params invalid or not enough", Data: nil})
		return
	}
	err = user.ReadByEmail()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		formatResponseBody(w, responseBody{Msg: "email already exist", Data: nil})
		return
	}
	if err != sql.ErrNoRows {
		handleInternalErr(w, err)
		return
	}
	err = user.ReadByName()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		formatResponseBody(w, responseBody{Msg: "user name already exist", Data: nil})
		return
	}
	if err != sql.ErrNoRows {
		handleInternalErr(w, err)
		return
	}
	err = user.Create()
	if err != nil {
		handleInternalErr(w, err)
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
	err := utils.ParseBody(r, user)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if (user.Email == "" && user.Name == "") || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: "params invalid or not enough", Data: nil})
		return
	}

	var (
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
		formatResponseBody(w, responseBody{Msg: fmt.Sprintf("user %s not found", accountType), Data: nil})
		return
	}
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if !utils.CheckPasswordHash(user.Password, rpwd) {
		w.WriteHeader(http.StatusBadRequest)
		formatResponseBody(w, responseBody{Msg: "password incorrect", Data: nil})
		return
	}

	session := models.Session{
		UserID: user.ID,
	}

	// TODO: sql INSERT ... ON DUPLICATE KEY UPDATE 目前遇到更新不了的問題
	// 應該優化這一段，感覺沒必要 query 兩次 (第一次先讀來判斷 第二次依據結果建立或更新)，以及因為跑兩次導致後續的 error handle 很不優雅
	err = session.ReadByUserID()

	session.UUID = uuid.New().String()
	session.Expiry = time.Now().Add(constants.Cookie.SessionTokenExpiry)

	if err == sql.ErrNoRows {
		err := session.Create()
		if err != nil {
			handleInternalErr(w, err)
			return
		}
	} else if err != nil {
		handleInternalErr(w, err)
		return
	} else {
		err := session.UpdateByUserID()
		if err != nil {
			handleInternalErr(w, err)
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     constants.Cookie.SessionTokenName,
		Value:    session.UUID,
		Expires:  session.Expiry,
		HttpOnly: true,
		Path:     "/",
	})

	formatResponseBody(w, responseBody{Msg: "success", Data: user})
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie(constants.Cookie.SessionTokenName)
	session := models.Session{
		UUID:   c.Value,
		Expiry: time.Now(),
	}
	err := session.UpdateByUUID()
	if err != nil {
		handleInternalErr(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     constants.Cookie.SessionTokenName,
		Value:    "",
		Expires:  session.Expiry,
		HttpOnly: true,
		Path:     "/",
	})

	formatResponseBody(w, responseBody{Msg: "success", Data: nil})
}

func handleInternalErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	formatResponseBody(w, responseBody{Msg: fmt.Sprintf("%s\n%s", err, debug.Stack()), Data: nil})
}

var Auth = auth{
	Signup: http.HandlerFunc(signup),
	Login:  http.HandlerFunc(login),
	Logout: http.HandlerFunc(logout),
}
