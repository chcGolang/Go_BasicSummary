package main

import (
	"Go_BasicSummary/rpc_demo/go_rpc/proto_file"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	request := &proto_file.String{
		Value: "hello",
	}
	reply := &proto_file.String{
		Value: "hi",
	}
	err = client.Hello(request, reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request *proto_file.String, reply *proto_file.String) error
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	codec := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: codec}, nil
}

func (p *HelloServiceClient) Hello(request *proto_file.String, reply *proto_file.String) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
