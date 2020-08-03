package main

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/grpc"

	"github.com/zenoss/zenkit/v5"
	// proto "github.com/zenoss/zing-proto/v11/go/cloud/{{Name}}"
)

const (
	// ServiceName is the name if this microservice.
	ServiceName = "{{Name}}"
)

func main() {
	zenkit.InitConfig(ServiceName)

	log := zenkit.Logger(ServiceName)
	ctx, cancel := context.WithCancel(ctxlogrus.ToContext(context.Background(), log))
	defer cancel()

	err := zenkit.RunGRPCServer(ctx, ServiceName, func(svr *grpc.Server) error {

		// Fill this in with your service details

		// proto.Register{{Name | title}}ServiceServer(svr, New{{Name | title}}Service())

		return nil

	})
	if err != nil {
		log.WithError(err).Fatal("error running gRPC server")
	}
}
