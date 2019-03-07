package main

import (
	"context"
	"google.golang.org/grpc/grpclog"

	"github.com/zenoss/zenkit"
	// proto "github.com/zenoss/zing-proto/go/{{Name}}"
	"google.golang.org/grpc"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := zenkit.RunGRPCServer(ctx, "{{Name}}", func(svr *grpc.Server) error {

		// Fill this in with your service details

		// proto.Register{{Name | title}}ServiceServer(svr, New{{Name | title}}Service())

		return nil

	})
	if err != nil {
		grpclog.Errorf("Error running GRPC server: %s", err.Error())
	}
}
