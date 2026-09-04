package service

import (
	"context"

	pb "ManogyaDahal/server/proto/gen/calculator"
)

// Defines what service do our interface implements
type CalculatorService interface {
	AddNumber(context.Context, *pb.InputNumReq) (*pb.OutNumResp, error)
	MulNumber(context.Context, *pb.InputTwoNumReq) (*pb.OutNumResp, error)
	SubNumber(context.Context, *pb.InputNumReq) (*pb.OutNumResp, error)
	DivNumber(context.Context, *pb.InputTwoNumReq) (*pb.OutNumResp, error)
}

// service struct currently holds nothing
// but has implementation of interfaces i.e struct methodes
type svc struct{}

// returns new service
func NewService() CalculatorService {
	return &svc{}
}

func (s *svc) AddNumber(ctx context.Context, req *pb.InputNumReq) (*pb.OutNumResp, error) {
	var sum int64
	for _, v := range req.InputNum {
		sum += v
	}
	return &pb.OutNumResp{Out: sum}, nil
}

func (s *svc) MulNumber(ctx context.Context, req *pb.InputTwoNumReq) (*pb.OutNumResp, error) {
	return &pb.OutNumResp{Out: req.Input_A * req.Input_B}, nil
}

func (s *svc) SubNumber(ctx context.Context, req *pb.InputNumReq) (*pb.OutNumResp, error) {
	return &pb.OutNumResp{Out: req.InputNum[0] - req.InputNum[1]}, nil
}

func (s *svc) DivNumber(ctx context.Context, req *pb.InputTwoNumReq) (*pb.OutNumResp, error) {
	return &pb.OutNumResp{Out: req.Input_A / req.Input_B}, nil
}
