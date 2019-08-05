package main

import (
	"Go_BasicSummary/rpc_demo/grpc_demo/middle"
	"Go_BasicSummary/rpc_demo/grpc_demo/protos"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "chc.com", // NOTE: this is required!
		RootCAs:      certPool,
	})

	authentication := middle.Authentication{
		User:     "chc",
		Password: "123456",
	}
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&authentication))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := protos.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &protos.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			// Grpc流
			if err := stream.Send(&protos.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}

}
