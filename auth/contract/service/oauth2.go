package service

import "net/http"

type IOauth2 interface {
	HandleAuthorizeRequest(http.ResponseWriter, *http.Request) (err error)
	HandleTokenRequest(http.ResponseWriter, *http.Request) (err error)
	Unwrap(err error)
}
