package main

import (
	"context"
	"fmt"
	"grpc/message/messagepb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't Connect To Server: %v", err)
	}
	defer con.Close()
	c := messagepb.NewmessageServiceClient(con)
	// fmt.Println("Created Client Successfully: ", c)
	// Unary(c)
	// ServerStreaming(c)
	// ClientStreaming(c)
	BiDiStreaming(c)

}

func Unary(c messagepb.messageServiceClient) {

	req := &messagepb.messageRequest{
		messageing: &messagepb.messageing{
			FirstName: "Paveen",
			LastName:  "Kumar",
		},
	}

	res, err := c.message(context.Background(), req)
	if err != nil {
		log.Fatalf("Can't Call message Function On the Server: %v", err)
	}

	log.Printf("Response From message Function %v", res.Result)
}

// Server Streaming

func ServerStreaming(c messagepb.messageServiceClient) {

	req := &messagepb.messageManyTimeRequest{
		messageing: &messagepb.messageing{
			FirstName: "Paveen",
			LastName:  "Kumar",
		},
	}

	resStream, err := c.messageManyTime(context.Background(), req)

	// fmt.Println(context.Background())

	if err != nil {
		log.Fatalf("Can't Call The Function In the Server: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error While Reading The Steams: %v", err)
		}

		log.Printf("Stream Of Response From messageManyTime: %v", msg.GetResult())
	}
}

// Client Streaming

func ClientStreaming(c messagepb.messageServiceClient) {

	requests := []*messagepb.LongmessageRequest{
		&messagepb.LongmessageRequest{
			messageing: &messagepb.messageing{
				FirstName: "Paveen ",
			},
		},
		&messagepb.LongmessageRequest{
			messageing: &messagepb.messageing{
				FirstName: "Harish ",
			},
		},
		&messagepb.LongmessageRequest{
			messageing: &messagepb.messageing{
				FirstName: "Goutham ",
			},
		},
		&messagepb.LongmessageRequest{
			messageing: &messagepb.messageing{
				FirstName: "Sharan ",
			},
		},
		&messagepb.LongmessageRequest{
			messageing: &messagepb.messageing{
				FirstName: "Durai ",
			},
		},
	}
	stream, err := c.Longmessage(context.Background())

	if err != nil {
		log.Fatalf("Can't Send Stream To Server: %v\n", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending Request %v \n", req)
		stream.Send(req)
		time.Sleep(2000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Can't Receive Data From The Server: %v\n", err)
	}

	log.Printf("The Response From The Server: %v\n", res)
}

// BiDirectional Streaming

func BiDiStreaming(c messagepb.messageServiceClient) {
	stream, err := c.messageEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error While Creating Stream: %v", err)
	}

	requests := []*messagepb.messageEveryOneRequest{
		&messagepb.messageEveryOneRequest{
			messageing: &messagepb.messageing{
				FirstName: "Paveen ",
			},
		},
		&messagepb.messageEveryOneRequest{
			messageing: &messagepb.messageing{
				FirstName: "Harish ",
			},
		},
		&messagepb.messageEveryOneRequest{
			messageing: &messagepb.messageing{
				FirstName: "Goutham ",
			},
		},
		&messagepb.messageEveryOneRequest{
			messageing: &messagepb.messageing{
				FirstName: "Sharan ",
			},
		},
		&messagepb.messageEveryOneRequest{
			messageing: &messagepb.messageing{
				FirstName: "Durai ",
			},
		},
	}

	cWait := make(chan struct{})

	go func() {
		for _, val := range requests {
			stream.Send(val)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(cWait)
			}
			if err != nil {
				log.Printf("Can't Receive From The Server: %v", err)
				close(cWait)
			}
			fmt.Println("Respose From The Server: ", res.Result)

		}
	}()

	<-cWait

}
