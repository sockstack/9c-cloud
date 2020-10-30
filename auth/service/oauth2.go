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
	session *SessionService
	option *oauth2Option
}

func NewOauth2(handlers ...Oauth2OptionHandler) (oauth2 *Oauth2) {
	option := &oauth2Option{LoginAddress: "/login"}
	for _, h := range handlers {
		h(option)
	}

	oauth2 = &Oauth2{option: option}
	oauth2.init()
	return
}

func (o *Oauth2) init() {
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
	s.SetUserAuthorizationHandler(o.userAuthorizationHandler)
	s.SetResponseErrorHandler(o.responseErrorHandler)
	s.SetInternalErrorHandler(o.internalErrorHandler)

	o.server = s
	o.session = NewSessionService()
}

func (o *Oauth2) HandleAuthorizeRequest(w http.ResponseWriter, r *http.Request) (err error) {
	session, err := o.session.Get(r, "oauth2_id")
	if err != nil {
		return
	}
	if session.Values["oauth2_id"] == nil {
		if r.Form == nil {
			if err = r.ParseForm(); err != nil {
				return
			}
		}
		if err = o.session.Save(w, r, "rf", r.Form); err != nil {
			return
		}
		session, _ = o.session.Get(r, "rf")
		w.Header().Set("Location", o.option.LoginAddress)
		w.WriteHeader(http.StatusFound)
		return nil
	}
	err = o.server.HandleAuthorizeRequest(w, r)
	return
}

func (o *Oauth2) HandleTokenRequest(w http.ResponseWriter, r *http.Request) (err error)  {
	err = o.server.HandleTokenRequest(w, r)
	return
}

func (o *Oauth2) userAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return "123", nil
}

func (o *Oauth2) responseErrorHandler(re *errors.Response) {
	log.Println(re)
}

func (o *Oauth2) internalErrorHandler(err error) (re *errors.Response){
	return
}

func (o *Oauth2) Unwrap(err error) {
	if err != nil {
		panic(err)
	}
}

type oauth2Option struct {
	LoginAddress string
}

type Oauth2OptionHandler func(o *oauth2Option)

func WithLoginAddress(address string) Oauth2OptionHandler {
	return func(o *oauth2Option) {
		o.LoginAddress = address
	}
}
