package controllers

import "net/http"

type thread struct {
	GetThreads            http.Handler
	GetThreadById         http.Handler
	UpdateThread          http.Handler
	GetThreadRelatedPosts http.Handler
}

func getThreads(w http.ResponseWriter, r *http.Request) {

}
func getThreadById(w http.ResponseWriter, r *http.Request) {

}
func updateThread(w http.ResponseWriter, r *http.Request) {

}
func getThreadRelatedPosts(w http.ResponseWriter, r *http.Request) {

}

var Thread = thread{
	GetThreads:            http.HandlerFunc(getThreads),
	GetThreadById:         http.HandlerFunc(getThreadById),
	UpdateThread:          http.HandlerFunc(updateThread),
	GetThreadRelatedPosts: http.HandlerFunc(getThreadRelatedPosts),
}
