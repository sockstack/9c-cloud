package service

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type ISessionService interface {
	Get(r *http.Request, key string) (session *sessions.Session, err error)
	Save(w http.ResponseWriter, r *http.Request, key string, data interface{}) (err error)
}
