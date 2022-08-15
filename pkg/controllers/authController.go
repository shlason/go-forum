package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/shlason/go-forum/pkg/constants"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
	"github.com/shlason/go-forum/pkg/utils"
)

type auth struct {
	Signup http.Handler
	Login  http.Handler
	Logout http.Handler
}

// TODO: swaggo 文件看熟，把 request payload 和 response 會有什麼都寫清楚
// signup godoc
// @Summary 	 Create account
// @Description  Create account by email, name, password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param		 email 	  body string true "user email"
// @Param		 name 	  body string true "user name"
// @Param		 password body string true "user password"
// @Success      200      {object}  structs.ResponseBody{data=models.User}
// @Failure      400      {object}  structs.ResponseBody
// @Failure      404      {object}  structs.ResponseBody
// @Failure      500      {object}  structs.ResponseBody
// @Router       /signup [post]
// TODO: 將 Request Body 的 Struct 和 Model 的 Struct 解耦來更彈性一點，以及控制資料的可見度
func signup(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := utils.ParseBody(r, user)

	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if user.Email == "" || user.Name == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "params invalid or not enough", Data: nil})
		return
	}
	err = user.ReadByEmail()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "email already exist", Data: nil})
		return
	}
	if err != sql.ErrNoRows {
		handleInternalErr(w, err)
		return
	}
	err = user.ReadByName()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "user name already exist", Data: nil})
		return
	}
	if err != sql.ErrNoRows {
		handleInternalErr(w, err)
		return
	}
	err = user.Create()
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	structs.WriteResponseBody(w, structs.ResponseBody{Msg: fmt.Sprintf("%s", err), Data: user})
}

var accountTypes = map[string]string{
	"email": "email",
	"name":  "name",
}

func login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := utils.ParseBody(r, user)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if (user.Email == "" && user.Name == "") || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "params invalid or not enough", Data: nil})
		return
	}

	var (
		accountType string
		rpwd        string = user.Password
	)

	if user.Email != "" {
		accountType = accountTypes["email"]
	}
	if user.Name != "" {
		accountType = accountTypes["name"]
	}

	switch accountType {
	case accountTypes["eamil"]:
		err = user.ReadByEmail()
	case accountTypes["name"]:
		err = user.ReadByName()
	}
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: fmt.Sprintf("user %s not found", accountType), Data: nil})
		return
	}
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if !utils.CheckPasswordHash(user.Password, rpwd) {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "password incorrect", Data: nil})
		return
	}

	session := models.Session{
		UserID: user.ID,
	}

	// TODO: sql INSERT ... ON DUPLICATE KEY UPDATE 目前遇到更新不了的問題
	// 應該優化這一段，感覺沒必要 query 兩次 (第一次先讀來判斷 第二次依據結果建立或更新)，以及因為跑兩次導致後續的 error handle 很不優雅
	err = session.ReadByUserID()

	session.UUID = uuid.New().String()
	session.Expiry = time.Now().Add(constants.Cookie.SessionTokenExpiry)

	if err == sql.ErrNoRows {
		err := session.Create()
		if err != nil {
			handleInternalErr(w, err)
			return
		}
	} else if err != nil {
		handleInternalErr(w, err)
		return
	} else {
		err := session.UpdateByUserID()
		if err != nil {
			handleInternalErr(w, err)
			return
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name:     constants.Cookie.SessionTokenName,
		Value:    session.UUID,
		Expires:  session.Expiry,
		HttpOnly: true,
		Path:     "/",
	})

	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: user})
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie(constants.Cookie.SessionTokenName)
	session := models.Session{
		UUID:   c.Value,
		Expiry: time.Now(),
	}
	err := session.UpdateByUUID()
	if err != nil {
		handleInternalErr(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     constants.Cookie.SessionTokenName,
		Value:    "",
		Expires:  session.Expiry,
		HttpOnly: true,
		Path:     "/",
	})

	structs.WriteResponseBody(w, structs.ResponseBody{Msg: "success", Data: nil})
}

var Auth = auth{
	Signup: http.HandlerFunc(signup),
	Login:  http.HandlerFunc(login),
	Logout: http.HandlerFunc(logout),
}
