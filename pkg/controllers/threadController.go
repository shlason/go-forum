package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
	"github.com/shlason/go-forum/pkg/utils"
)

type thread struct {
	GetThreads            http.Handler
	CreateThread          http.Handler
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

type threadPayload struct {
	Subject string `json:"subject"`
}

func createThread(w http.ResponseWriter, r *http.Request) {
	session, err := getSession(r)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	payload := &threadPayload{}
	err = utils.ParseBody(r, payload)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if payload.Subject == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "subject empty", Data: nil})
		return
	}
	thread := models.Thread{
		UserID:  session.UserID,
		Subject: payload.Subject,
	}
	err = thread.Create()
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: nil})
}

func getThreadById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	threadID := params["threadID"]
	tid, err := strconv.Atoi(threadID)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	thread := models.Thread{
		ID: tid,
	}
	err = thread.ReadByID()
	if err != nil {
		if err != sql.ErrNoRows {
			handleInternalErr(w, err)
			return
		}
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: nil})
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: thread})
}

func updateThread(w http.ResponseWriter, r *http.Request) {

}

func getThreadRelatedPosts(w http.ResponseWriter, r *http.Request) {

}

var Thread = thread{
	GetThreads:            http.HandlerFunc(getThreads),
	CreateThread:          http.HandlerFunc(createThread),
	GetThreadById:         http.HandlerFunc(getThreadById),
	UpdateThread:          http.HandlerFunc(updateThread),
	GetThreadRelatedPosts: http.HandlerFunc(getThreadRelatedPosts),
}
