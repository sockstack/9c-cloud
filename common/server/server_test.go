package server

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	newServer := NewServer()
	var _ serverInterface = newServer
}

func TestRun(t *testing.T) {
	newServer := NewServer()
	newServer.Run(":8080")
}