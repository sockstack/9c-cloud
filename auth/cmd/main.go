package main

import (
	"github.com/sockstack/9c-cloud/auth"
	"github.com/sockstack/9c-cloud/auth/routes"
)

func main() {
	auth.Engine.
		Route(routes.Api).
		Run(auth.WithAddress(":9096"))
}