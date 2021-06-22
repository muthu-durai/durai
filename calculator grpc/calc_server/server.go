package main

import (
	"context"
	"grpc/calculator/calcpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calc(ctx context.Context, in *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	num1, num2 := in.GetNum1(), in.GetNum2()
	result := num1 + num2
	res := &calcpb.CalcResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Can't Listen To The Port: %v", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed To Server: %v", err)
	}
}
