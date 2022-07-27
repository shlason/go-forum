package controllers

import (
	"net/http"
	"time"
)

type user struct {
	GetUsers   http.Handler
	PatchUsers http.Handler
}

type userResponse struct {
	id        int
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	// userId := r.URL.Query().Get("userId")

}

func patchUsers(w http.ResponseWriter, r *http.Request) {}

var User = user{
	GetUsers:   http.HandlerFunc(getUsers),
	PatchUsers: http.HandlerFunc(patchUsers),
}
