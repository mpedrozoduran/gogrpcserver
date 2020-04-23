package timemanager

import (
	"context"
	pb "github.com/mpedrozoduran/gogrpctest/timeproto"
)

type TimeStruct struct {
	pb.UnimplementedTimeManagerServer
}

func (t *TimeStruct) GetTime(ctx context.Context, request *pb.TimeRequest) (*pb.TimeResponse, error) {
	return &pb.TimeResponse{ServerTime: "2020-04-23 14:17:00"}, nil
}
