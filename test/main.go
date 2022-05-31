package main

import (
	"DNSpcap/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewDNSProcessorClient(conn)
	r, err := c.IsEval(context.Background(), &proto.Request{Domain: "www.baidu.com"})
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Res)
}
