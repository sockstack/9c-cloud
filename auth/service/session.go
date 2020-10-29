package service

import "github.com/gorilla/sessions"

type SessionService struct {
	store sessions.Store
}

func NewSessionService() *SessionService {
	return &SessionService{
		store: sessions.NewCookieStore([]byte("test")),
	}
}

func (s *SessionService) Get(key string)  {

}

func (s *SessionService) Save(key string, data interface{})  {

}
