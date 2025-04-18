package main

import (
	"context"

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

	ctx := zenkit.LoggerToContext(context.Background(), log)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := zenkit.WaitUntilEnvoyReady(log); err != nil {
		log.WithError(err).Fatal("waiting for envoy failed")
	}

	cancelInstrumentation, err := zenkit.InstallInstrumentation(ctx, log)
	if err != nil {
		log.WithError(err).Error("could not install instrumentation")
	}

	defer cancelInstrumentation()

	err := zenkit.RunGRPCServer(ctx, ServiceName, func(svr *grpc.Server) error {
		// Fill this in with your service details

		// proto.Register{{Name | title}}ServiceServer(svr, New{{Name | title}}Service())

		return nil
	})
	// Replace RunGRPCServer above with this if you're exposing a public API.
	//
	// err := zenkit.RunGRPCServerWithEndpoint(ctx, ServiceName, func(svr *grpc.Server) error {
	// 	proto.Register{{Name | title}}ServiceServer(svr, serviceServer)
	// 	return nil
	// }, proto.Register{{Name | title}}ServiceHandlerFromEndpoint)
	if err != nil {
		log.WithError(err).Fatal("error running gRPC server")
	}
}
