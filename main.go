package main

import (
	"fmt"
	"github.com/mpedrozoduran/gogrpcserver/timemanager"
	pb "github.com/mpedrozoduran/gogrpcserver/timeproto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = 9091
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTimeManagerServer(grpcServer, &timemanager.TimeStruct{})
	grpcServer.Serve(lis)
}
