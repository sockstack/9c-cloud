package service

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type SessionService struct {
	store sessions.Store
}

func NewSessionService() *SessionService {
	service := &SessionService{}
	service.init()
	return service
}

func (s *SessionService) init()  {
	s.store = sessions.NewCookieStore([]byte("test"))
}

func (s *SessionService) Get(r *http.Request, key string) (session *sessions.Session, err error) {
	session, err = s.store.Get(r, key)
	return
}

func (s *SessionService) Save(w http.ResponseWriter, r *http.Request, key string, data interface{}) (err error) {
	session, err := s.Get(r, key)
	if err != nil {
		return
	}
	session.Values[key] = data
	err = s.store.Save(r, w, session)
	return
}
