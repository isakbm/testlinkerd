package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/testlinkerd/pkg/world"
	"google.golang.org/grpc"
)

type worldServer struct{}

func (s *worldServer) SayHello(_ context.Context, _ *world.HelloReq) (*world.HelloResp, error) {
	log.Println("got Hello request")
	return &world.HelloResp{Msg: "World"}, nil
}

func main() {

	server := grpc.NewServer()
	world.RegisterWorldServer(server, &worldServer{})

	lis, err := net.Listen("tcp", ":50040")
	if err != nil {
		panic(err)
	}

	fmt.Println("starting server, listening on", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		fmt.Println("err:", err)
	}
}
