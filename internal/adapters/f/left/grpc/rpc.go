package grpc

import (
	"context"
	"errors"

	"github.com/404th/hex_arch/internal/adapters/f/left/grpc/pb"
)

func (ga Adapter) GetAddition(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error) {
	var response *pb.Answer

	if req.GetA() == 0 || req.GetB() == 0 {
		return response, errors.New("Unsatisfied inputs!")
	}

	result, err := ga.api.GetAddition(req.GetA(), req.GetB())
	if err != nil {
		return response, err
	}

	response.Result = result

	return response, nil
}

func (ga Adapter) GetSubtraction(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error) {
	var response *pb.Answer

	if req.GetA() == 0 || req.GetB() == 0 {
		return response, errors.New("Unsatisfied inputs!")
	}

	result, err := ga.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return response, err
	}

	response.Result = result

	return response, nil
}

func (ga Adapter) GetMultiplication(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error) {
	var response *pb.Answer

	if req.GetA() == 0 || req.GetB() == 0 {
		return response, errors.New("Unsatisfied inputs!")
	}

	result, err := ga.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return response, err
	}

	response.Result = result

	return response, nil
}

func (ga Adapter) GetDivision(ctx context.Context, req *pb.OperationDetails) (*pb.Answer, error) {
	var response *pb.Answer

	if req.GetA() == 0 || req.GetB() == 0 {
		return response, errors.New("Unsatisfied inputs!")
	}

	result, err := ga.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return response, err
	}

	response.Result = result

	return response, nil
}
