package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
	"github.com/shlason/go-forum/pkg/utils"
)

type post struct {
	GetPosts    http.Handler
	GetPostByID http.Handler
	UpdatePost  http.Handler
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	p := models.Post{}
	posts, err := p.ReadAll()
	if err != nil && err != sql.ErrNoRows {
		handleInternalErr(w, err)
		return
	}
	if posts == nil {
		posts = make([]models.Post, 0)
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: posts})
}

func getPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postUUID := params["postUUID"]
	post := models.Post{
		UUID: postUUID,
	}
	err := post.ReadByUUID()
	if err != nil {
		if err != sql.ErrNoRows {
			handleInternalErr(w, err)
			return
		}
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: nil})
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: post})
}

type patchPostPayload struct {
	Content string `json:"content"`
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postUUID := params["postUUID"]
	post := models.Post{
		UUID: postUUID,
	}
	err := post.ReadByUUID()
	if err != nil {
		if err != sql.ErrNoRows {
			handleInternalErr(w, err)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "post not found", Data: nil})
		return
	}
	session, err := getSession(r)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if post.UserID != session.UserID {
		w.WriteHeader(http.StatusUnauthorized)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "unauthorized", Data: nil})
		return
	}
	payload := &patchPostPayload{}
	err = utils.ParseBody(r, payload)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if payload.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "invalid content", Data: nil})
		return
	}
	post.Content = payload.Content
	err = post.UpdateByUUID()
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: post})
}

var Post = post{
	GetPosts:    http.HandlerFunc(getPosts),
	GetPostByID: http.HandlerFunc(getPostByID),
	UpdatePost:  http.HandlerFunc(updatePost),
}
