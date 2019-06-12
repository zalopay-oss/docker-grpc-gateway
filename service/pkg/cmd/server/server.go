package cmd

import (
	"context"
	"net"
	"os"
	"os/signal"
	service_proto "service/pkg/api"
	"service/pkg/serviceA"
	"service/pkg/serviceB"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

//RunServer run server
func RunServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	servA := serviceA.NewServiceA()
	servB := serviceB.NewServiceB()

	server := grpc.NewServer()
	service_proto.RegisterServiceAServer(server, servA)
	service_proto.RegisterServiceBServer(server, servB)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Info("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Info("Start service A, service B port " + port + " ...")
	return server.Serve(listen)
}
