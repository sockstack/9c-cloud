package auth

import (
	"github.com/gin-gonic/gin"
	"log"
)

var Engine = newAuth()

type engine struct {
	server *gin.Engine
	Address string
}

type R func(engine *gin.Engine)

func newAuth() *engine {
	return &engine{server: gin.New()}
}

func (e *engine) Route(route R) *engine {
	route(e.server)
	return e
}

func (e *engine) Use(handlerFunc ...gin.HandlerFunc) *engine {
	e.server.Use(handlerFunc...)
	return e
}

func (e *engine) Run(handles ...OptionHandleFunc) {
	for _, h := range handles {
		h(e)
	}

	if err := e.server.Run(e.Address); err != nil {
		log.Fatalln(err)
	}
}

type OptionHandleFunc func(e *engine)

func WithAddress(address string) OptionHandleFunc {
	return func(e *engine) {
		e.Address = address
	}
}