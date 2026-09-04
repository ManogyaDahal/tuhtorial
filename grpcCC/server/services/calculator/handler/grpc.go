package handler

import (
	"context"

	pb "ManogyaDahal/server/proto/gen/calculator"
	"ManogyaDahal/server/services/calculator/service"
)

type GrpcCalculatorHandler struct {
	pb.UnimplementedCalculatorServer
	svc service.CalculatorService
}

func NewGrpcHandlers(svc service.CalculatorService) *GrpcCalculatorHandler {
	return &GrpcCalculatorHandler{svc: svc}
}

func (h *GrpcCalculatorHandler) AddNumber(ctx context.Context, req *pb.InputNumReq) (*pb.OutNumResp, error) {
	return h.svc.AddNumber(ctx, req)
}

func (h *GrpcCalculatorHandler) MulNumber(ctx context.Context, req *pb.InputTwoNumReq) (*pb.OutNumResp, error) {
	return h.svc.MulNumber(ctx, req)
}

func (h *GrpcCalculatorHandler) SubNumber(ctx context.Context, req *pb.InputNumReq) (*pb.OutNumResp, error) {
	return h.svc.SubNumber(ctx, req)
}

func (h *GrpcCalculatorHandler) DivNumber(ctx context.Context, req *pb.InputTwoNumReq) (*pb.OutNumResp, error) {
	return h.svc.DivNumber(ctx, req)
}
