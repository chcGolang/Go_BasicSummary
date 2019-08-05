package main

import (
	"Go_BasicSummary/rpc_demo/go_rpc/proto_file"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request *proto_file.String, reply *proto_file.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

const HelloServiceName = "HelloService"

func RegisterHelloService(rpcService interface{}) {
	rpc.RegisterName(HelloServiceName, rpcService)
}
