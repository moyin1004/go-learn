package main

import (
	"context"
	"fmt"
	examplepb "go-learn/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func PressGrpc() {
	ctx := context.Background()
	conn, err := grpc.NewClient("192.168.1.37:8010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := examplepb.NewHttpServiceClient(conn)
	rsp, err := c.TestGrpc(ctx, &examplepb.TestGrpcReq{
		Id:   2,
		Data: "3",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp)
	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// design.TestMap()
	// design.TestPipeline()
	for i := 0; i < 2; i++ {
		go PressGrpc()
	}
	time.Sleep(10 * time.Second)
}
