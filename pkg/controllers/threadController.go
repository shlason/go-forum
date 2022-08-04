package controllers

import (
	"database/sql"
	"net/http"

	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
)

type thread struct {
	GetThreads            http.Handler
	GetThreadById         http.Handler
	UpdateThread          http.Handler
	GetThreadRelatedPosts http.Handler
}

func getThreads(w http.ResponseWriter, r *http.Request) {
	var threads []models.Thread
	thread := models.Thread{}
	threads, err := thread.ReadAll()
	// TODO: 待研究 - 沒有資料找不到時，應該要噴 sql.ErrNoRows 才對，但這邊不會有錯
	if err != nil && err != sql.ErrNoRows {
		handleInternalErr(w, err)
		return
	}
	if threads == nil {
		threads = make([]models.Thread, 0)
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: threads})
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
