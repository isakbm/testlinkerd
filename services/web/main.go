package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/testlinkerd/pkg/world"
	"google.golang.org/grpc"
)

var (
	worldTarget = flag.String("target", ":50040", "specify target")
	lisPort     = flag.String("lisPort", "8080", "override the listen port")
	creds       = flag.String("creds", "/creds", "specify directory of credential keypair")
	servername  = flag.String("servername", "web", "override servername")
)

func main() {

	flag.Parse()

	cc, err := grpc.Dial(
		*worldTarget,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	worldClient := world.NewWorldClient(cc)

	cert, err := tls.LoadX509KeyPair(*creds+"/tls.crt", *creds+"/tls.key")
	if err != nil {
		panic(err)
	}

	lis, err := tls.Listen("tcp", ":"+*lisPort, &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		panic(err)
	}

	// lis, err := net.Listen("tcp", ":"+*lisPort)
	// if err != nil {
	// 	panic(err)
	// }

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
