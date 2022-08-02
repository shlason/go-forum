package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/shlason/go-forum/pkg/constants"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
	"github.com/shlason/go-forum/pkg/utils"
)

type user struct {
	GetUsers  http.Handler
	PatchUser http.Handler
}

type userResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	user := models.User{}
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		users, err := user.ReadAll()
		if err != nil {
			handleInternalErr(w, err)
			return
		}
		var res []userResponse
		for _, u := range users {
			res = append(res, userResponse{
				ID:        u.ID,
				Name:      u.Name,
				Email:     u.Email,
				CreatedAt: u.CreatedAt,
				UpdatedAt: u.UpdatedAt,
			})
		}
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: res})
		return
	}
	user.ID, err = strconv.Atoi(userId)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	err = user.ReadByUserID()
	if err != nil {
		if err == sql.ErrNoRows {
			structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: nil})
			return
		}
		handleInternalErr(w, err)
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}})
}

type patchUserPayload struct {
	Password string `json:"password"`
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	var err error
	// Auth middleware 檢查過了，這邊忽略 error
	c, _ := r.Cookie(constants.Cookie.SessionTokenName)
	session := models.Session{
		UUID: c.Value,
	}
	err = session.ReadByUUID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payload := &patchUserPayload{}
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "need user id params", Data: nil})
		return
	}
	user := models.User{}
	user.ID, err = strconv.Atoi(userId)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if user.ID != session.UserID {
		w.WriteHeader(http.StatusUnauthorized)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "Unauthorized", Data: nil})
		return
	}
	err = utils.ParseBody(r, payload)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if payload.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "password invalid", Data: nil})
		return
	}
	user.Password = payload.Password
	err = user.UpdateUserPasswordByUserID()
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: nil})
}

var User = user{
	GetUsers:  http.HandlerFunc(getUsers),
	PatchUser: http.HandlerFunc(patchUser),
}
