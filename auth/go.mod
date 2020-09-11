module github.com/sockstack/9c-cloud/auth

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/sockstack/9c-cloud/common v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/sockstack/9c-cloud/common => ../common
