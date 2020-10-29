package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"net/http"
)

type Oauth2 struct {
	server *server.Server
}

func NewOauth2() *Oauth2 {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(server.ClientFormHandler)
	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println(err)
		return
	})
	s.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println(re)
	})
	
	s.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		return "123", nil
	})

	return &Oauth2{server: s}
}

func (o *Oauth2) HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) (err error) {
	err = o.server.HandleAuthorizeRequest(w, r)
	return
}

func (o *Oauth2) HandleTokenRequest(w http.ResponseWriter, r *http.Request) (err error)  {
	err = o.server.HandleTokenRequest(w, r)
	return
}

func (o *Oauth2) Unwrap(err error) {
	if err != nil {
		panic(err)
	}
}
