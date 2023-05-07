package plugins

import (
	"context"

	"github.com/RGood/v8-rpc/internal/generated/protos/clock"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Clock is a clock plugin
type Clock struct {
	clock.UnimplementedClockServer
}

func NewClock() *Clock {
	return &Clock{}
}

// Now returns a current timestamp
func (c *Clock) Now(ctx context.Context, _ *clock.TimeRequest) (*clock.TimeResponse, error) {
	return &clock.TimeResponse{
		Id:   uuid.NewString(),
		Time: timestamppb.Now(),
	}, nil
}
