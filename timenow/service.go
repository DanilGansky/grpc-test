package timenow

import (
	"context"
	"log"
	"time"

	"github.com/littlefut/grpc-test/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct{}

func NewService() api.TimenowServiceServer {
	return &service{}
}

func (s *service) Now(ctx context.Context, req *api.TimenowRequest) (*api.TimenowResponse, error) {
	loc, err := time.LoadLocation(req.Timezone)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid timezone: %v", err)
	}

	log.Printf("showing local time for %s timezone\n", req.Timezone)
	return &api.TimenowResponse{Timestamp: time.Now().In(loc).Format(time.RFC1123)}, nil
}
