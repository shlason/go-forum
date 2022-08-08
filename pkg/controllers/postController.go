package controllers

import "net/http"

type post struct {
	GetPosts    http.Handler
	GetPostByID http.Handler
	UpdatePost  http.Handler
}

func getPosts(w http.ResponseWriter, r *http.Request) {

}

func getPostByID(w http.ResponseWriter, r *http.Request) {

}

func updatePost(w http.ResponseWriter, r *http.Request) {

}

var Post = post{
	GetPosts:    http.HandlerFunc(getPosts),
	GetPostByID: http.HandlerFunc(getPostByID),
	UpdatePost:  http.HandlerFunc(updatePost),
}
