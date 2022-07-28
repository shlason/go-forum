package controllers

import (
	"net/http"
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
	Email     string    `json:"eamil"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
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
}

func patchUsers(w http.ResponseWriter, r *http.Request) {}

var User = user{
	GetUsers:   http.HandlerFunc(getUsers),
	PatchUsers: http.HandlerFunc(patchUsers),
}
