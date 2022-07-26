package middlewares

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/shlason/go-forum/pkg/constants"
	"github.com/shlason/go-forum/pkg/models"
	"github.com/shlason/go-forum/pkg/structs"
)

func Auth() middlewareHandler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie(constants.Cookie.SessionTokenName)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				structs.WriteResponseBody(w, structs.ResponseBody{Msg: "Unauthorized", Data: nil})
				return
			}
			session := models.Session{
				UUID: c.Value,
			}
			err = session.ReadByUUID()
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusUnauthorized)
					structs.WriteResponseBody(w, structs.ResponseBody{Msg: "Unauthorized", Data: nil})
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if !time.Now().Before(session.Expiry) {
				w.WriteHeader(http.StatusUnauthorized)
				structs.WriteResponseBody(w, structs.ResponseBody{Msg: "Unauthorized", Data: nil})
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
