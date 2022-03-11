package grpc

import (
	"net"

	"github.com/404th/hex_arch/adapters/f/left/grpc/pb"
	"github.com/404th/hex_arch/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPorts
}

func (sv *Adapter) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterArithmeticServiceServer(grpcServer, sv.api)

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
