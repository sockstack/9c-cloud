package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	services []interface{}
	svr *grpc.Server
}

func NewServer() *server {
	svr := grpc.NewServer() //创建gRPC服务

	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(svr)
	return &server{
		svr: svr,
	}
}

func (s *server) Server() *grpc.Server {
	return s.svr
}

func (s *server) Run(address string) {
	lis, err := net.Listen("tcp", address)  //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	// 将监听交给gRPC服务处理
	err = s.svr.Serve(lis)
	if  err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
