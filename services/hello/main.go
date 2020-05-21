package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/testlinkerd/pkg/world"
	"google.golang.org/grpc"
)

var (
	target = flag.String("target", ":50040", "specify target")
)

func main() {

	flag.Parse()

	cc, err := grpc.Dial(
		*target,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	worldClient := world.NewWorldClient(cc)

	for {
		resp, err := worldClient.SayHello(context.Background(), &world.HelloReq{})
		if err != nil {
			fmt.Println("err:", err)
		}
		if resp != nil {
			fmt.Println("resp:", resp)
		}
		time.Sleep(1 * time.Second)
	}
}
