package primary

import (
	"context"

	"github.com/daibertdiego/golang_hex_architecture/src/adapters/framework/primary/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	answer := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return answer, status.Error(codes.InvalidArgument, "missing requirement")
	}

	arithAdapter, err := grpca.api.GetAddition(req.A, req.B)
	if err != nil {
		return answer, status.Error(codes.Internal, "unexpected error")
	}

	answer = &pb.Answer{
		Value: arithAdapter.Answer,
	}

	return answer, nil

}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	answer := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return answer, status.Error(codes.InvalidArgument, "missing requirement")
	}

	arithAdapter, err := grpca.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return answer, status.Error(codes.Internal, "unexpected error")
	}

	answer = &pb.Answer{
		Value: arithAdapter.Answer,
	}

	return answer, nil

}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	answer := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return answer, status.Error(codes.InvalidArgument, "missing requirement")
	}

	arithAdapter, err := grpca.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return answer, status.Error(codes.Internal, "unexpected error")
	}

	answer = &pb.Answer{
		Value: arithAdapter.Answer,
	}

	return answer, nil

}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	answer := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return answer, status.Error(codes.InvalidArgument, "missing requirement")
	}

	arithAdapter, err := grpca.api.GetDivision(req.A, req.B)
	if err != nil {
		return answer, status.Error(codes.Internal, "unexpected error")
	}

	answer = &pb.Answer{
		Value: arithAdapter.Answer,
	}

	return answer, nil

}
