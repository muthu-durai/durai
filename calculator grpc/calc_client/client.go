package main

import (
	"context"
	"grpc/calculator/calcpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't Connect To the Server: %v", err)
	}
	defer con.Close()
	c := calcpb.NewCalcServiceClient(con)
	unary(c)
}

func unary(c calcpb.CalcServiceClient) {

	req := &calcpb.CalcRequest{
		Num1: 5,
		Num2: 10,
	}

	res, err := c.Calc(context.Background(), req)
	if err != nil {
		log.Fatalf("Can't Call Calc Function: %v", err)
	}
	log.Printf("Response From The Server: %v ", res.Result)
}
