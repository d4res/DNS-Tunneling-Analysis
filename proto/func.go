package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
)

func IsEval(domain string) bool {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := NewDNSProcessorClient(conn)
	r, err := c.IsEval(context.Background(), &Request{Domain: domain})
	if err != nil {
		panic(err)
	}

	return r.Res
}
