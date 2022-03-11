package grpc

import (
	"net"

	"github.com/404th/hex_arch/internal/adapters/f/left/grpc/pb"
	"github.com/404th/hex_arch/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPorts
}

func (sv Adapter) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterArithmeticServiceServer(grpcServer, sv)

	if err = grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
