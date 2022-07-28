package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// type LogServiceServer interface{

// 	Message(context.Context, *MessageRequest) (MessageResponse, error)
// }

func main() {
	conn, err := net.Listen("tcp", "localhost:8080")

	if err != nil {

		fmt.Println(err)
	}

	gserver := grpc.NewServer()

	er1 := gserver.Serve(conn)

	if er1 != nil {
		fmt.Println(er1)
	} else {
		fmt.Println("SUCCESSFULLY COnnected")
	}

}
