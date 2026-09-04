// defines what services do it holds
package service 

import (
	"context"

	pb "ManogyaDahal/server/proto/gen/calculator"
)

// Defines what services do the calculator hold
type CalculatorService interface {
	AddNumber(context.Context, *pb.InputNumReq) (*pb.OutNumResp, error)
	MulNumber(context.Context, *pb.InputNumReq) (*pb.OutNumResp, error)
	SubNumber(context.Context, *pb.InputTwoNumReq) (*pb.OutNumResp, error)
	DivNumber(context.Context, *pb.InputTwoNumReq) (*pb.OutNumResp, error)
}

// currently no database integration so it don't have any fields.
type svc struct { 
	// repository
	//repo repo.Querier
}

func NewService() *svc {
	return &svc{}
}
