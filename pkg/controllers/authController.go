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

type signupPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: swaggo 文件看熟，把 request payload 和 response 會有什麼都寫清楚
// signup godoc
// @Summary 	 Create account
// @Description  Create account by email, name, password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param    	 Payload  body   signupPayload true "email, name, password"
// @Success      200      		 {object}  structs.ResponseBody{data=models.User}
// @Failure      400      		 {object}  structs.ResponseBody
// @Failure      404      		 {object}  structs.ResponseBody
// @Failure      409      		 {object}  structs.ResponseBody
// @Failure      500      		 {object}  structs.ResponseBody
// @Router       /signup [post]
// TODO: 將 Request Body 的 Struct 和 Model 的 Struct 解耦來更彈性一點，以及控制資料的可見度
func signup(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	payload := &signupPayload{}
	err := utils.ParseBody(r, payload)

	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if payload.Email == "" || payload.Name == "" || payload.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "params invalid or not enough", Data: nil})
		return
	}
	// TODO: 在 payload struct 新增方法來做和 model 之間資料的轉換
	user.Email = payload.Email
	user.Name = payload.Name
	user.Password = payload.Password
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

type loginPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// login godoc
// @Summary 	 Login
// @Description  Login by email or name and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param    	 Payload  body   loginPayload true "email or name and password"
// @Success      200      		 {object}  structs.ResponseBody{data=models.User}
// @Failure      400      		 {object}  structs.ResponseBody
// @Failure      404      		 {object}  structs.ResponseBody
// @Failure      500      		 {object}  structs.ResponseBody
// @Router       /login [post]
func login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	payload := &loginPayload{}
	err := utils.ParseBody(r, payload)
	if err != nil {
		handleInternalErr(w, err)
		return
	}
	if (payload.Email == "" && payload.Name == "") || payload.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		structs.WriteResponseBody(w, structs.ResponseBody{Msg: "params invalid or not enough", Data: nil})
		return
	}
	// TODO: 在 payload struct 新增方法來做和 model 之間資料的轉換
	user.Email = payload.Email
	user.Name = payload.Name
	user.Password = payload.Password
	var (
		accountType  string
		rpwd         string = user.Password
		accountTypes        = map[string]string{
			"email": "email",
			"name":  "name",
		}
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

// logout godoc
// @Summary 	 Logout
// @Description  Logout set session expiry now
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200      		 {object}  structs.ResponseBody
// @Failure      404      		 {object}  structs.ResponseBody
// @Failure      500      		 {object}  structs.ResponseBody
// @Router       /logout [post]
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
