package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

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

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	http.HandleFunc(
		"/sayHello",
		func(w http.ResponseWriter, req *http.Request) {
			log.Println("received a /sayHello request")
			log.Println("making a request to the world service")
			resp, err := worldClient.SayHello(context.Background(), &world.HelloReq{Msg: "hello!"})
			if err != nil {
				fmt.Println(err)
			}
			if resp != nil {
				log.Println("received", resp)
			}
		},
	)

	fmt.Println("listening on", lis.Addr().String())

	if err := http.Serve(lis, nil); err != nil {
		fmt.Println(err)
	}

	// for {
	// 	resp, err := worldClient.SayHello(context.Background(), &world.HelloReq{})
	// 	if err != nil {
	// 		fmt.Println("err:", err)
	// 	}
	// 	if resp != nil {
	// 		fmt.Println("resp:", resp)
	// 	}
	// 	time.Sleep(1 * time.Second)
	// }
}
