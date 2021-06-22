package main

import (
	"context"
	"fmt"
	"grpc/message/messagepb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

// Unary

func (*server) message(cntx context.Context, req *messagepb.messageRequest) (*messagepb.messageResponse, error) {
	firstName := req.messageing.GetFirstName()
	result := "Hello " + firstName
	res := &messagepb.messageResponse{
		Result: result,
	}
	return res, nil
}

// Server Streaming

func (*server) messageManyTime(req *messagepb.messageManyTimeRequest, stream messagepb.messageService_messageManyTimeServer) error {
	firstName := req.messageing.GetFirstName()
	fmt.Println(req)
	fmt.Println("req.Getmessageing() :", req.Getmessageing())
	fmt.Println("req.messageing", req.messageing)
	fmt.Println("stream", stream)
	for i := 0; i <= 10; i++ {
		result := "Hello " + firstName + " " + strconv.Itoa(i)
		res := &messagepb.messageManyTimeResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

// Client Streaming

func (*server) Longmessage(stream messagepb.messageService_LongmessageServer) error {
	result := " "
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&messagepb.LongmessageResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Can't Read Streams From The Client: %v", err)
		}
		firstName := res.messageing.FirstName
		result += "Hello " + firstName
	}
}

// Bidirection Streaming

func (*server) messageEveryone(stream messagepb.messageService_messageEveryoneServer) error {

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Can't Recevie From The Client: %v", err)
		}

		firstName := res.messageing.FirstName
		result := "Hello " + firstName

		stream.Send(&messagepb.messageEveryOneResponse{
			Result: result,
		})

	}
}

func main() {
	fmt.Println("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // Default Port For gRPC
	if err != nil {
		log.Fatalf("Can't Listen To The Port: %v", err)
	}
	s := grpc.NewServer()
	messagepb.RegistermessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed To Server: %v", err)
	}
}
