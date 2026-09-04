package handler

import (
	"context"
	pb "ManogyaDahal/server/proto/gen/calculator"
)


type GrpcCalculatorHandler struct { 
	pb.UnimplementedCalculatorServer	
}

func NewGrpcHandlers() *GrpcCalculatorHandler { 
	return &GrpcCalculatorHandler{}
}

// Provide sum of numbers
func (h *GrpcCalculatorHandler) AddNumber (ctx context.Context, req *pb.InputNumReq) (*pb.OutNumResp, error) {
	var sum int64 = 0
	for _, value := range req.InputNum{ 
		sum = sum + value
	}
	return &pb.OutNumResp{
		Out: sum,
	}, nil
}

// provides multiplication of numbers
func (h *GrpcCalculatorHandler) MulNumber(ctx context.Context, req *pb.InputNumReq) (*pb.OutNumResp, error) {
	var mul int64 = 1
	for _, value := range req.InputNum{ 
		mul = mul * value
	}
	return &pb.OutNumResp{
		Out: mul,
	}, nil
}

// Subtracts between two numbers
func (h *GrpcCalculatorHandler) SubNumber(ctx context.Context, req *pb.InputTwoNumReq) (*pb.OutNumResp, error) {
	return &pb.OutNumResp{
		Out: req.Input_A - req.Input_B,
	}, nil
}

// Divides between two numbers
func (h *GrpcCalculatorHandler) DivNumber(ctx context.Context, req *pb.InputTwoNumReq) (*pb.OutNumResp, error) {
	return &pb.OutNumResp{
		Out: req.Input_A / req.Input_B,
	}, nil
}
