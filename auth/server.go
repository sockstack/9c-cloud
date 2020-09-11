package main

import (
	"context"
	"crypto/md5"
	"fmt"
	test "github.com/sockstack/9c-cloud/auth/proto"
	server2 "github.com/sockstack/9c-cloud/common/server"
	"log"
)

// 业务实现方法的容器
type server struct{}

// 为server定义 DoMD5 方法 内部处理请求并返回结果
// 参数 (context.Context[固定], *test.Req[相应接口定义的请求参数])
// 返回 (*test.Res[相应接口定义的返回参数，必须用指针], error)
func (s *server) DoMD5(ctx context.Context, in *test.Req) (*test.Res, error) {
	log.Println("MD5方法请求JSON:"+in.JsonStr)
	return &test.Res{BackJson: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.JsonStr)))}, nil
}

func main() {
	newServer := server2.NewServer()
	s := server{}
	test.RegisterWaiterServer(newServer.Server(), &s)
	newServer.Run(":8028")
}
