package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
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

func updatePost(w http.ResponseWriter, r *http.Request) {

}

var Post = post{
	GetPosts:    http.HandlerFunc(getPosts),
	GetPostByID: http.HandlerFunc(getPostByID),
	UpdatePost:  http.HandlerFunc(updatePost),
}
