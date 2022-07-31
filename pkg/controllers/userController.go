package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
)

type user struct {
	GetUsers   http.Handler
	PatchUsers http.Handler
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

func patchUsers(w http.ResponseWriter, r *http.Request) {}

var User = user{
	GetUsers:   http.HandlerFunc(getUsers),
	PatchUsers: http.HandlerFunc(patchUsers),
}
