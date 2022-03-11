package ports

import (
	"context"

	"github.com/404th/hex_arch/internal/adapters/f/left/grpc/pb"
)

type GRPCPorts interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error)
}
