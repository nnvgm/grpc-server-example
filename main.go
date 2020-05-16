package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/nnvgm/grpc-common-example/proto/math"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *math.SumRequest) (*math.SumResponse, error) {
	a := req.GetA()
	b := req.GetB()

	sum := a + b

	return &math.SumResponse{
		Sum: sum,
	}, nil
}

func main() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	math.RegisterMathServer(s, &server{})

	log.Println("Start GRPC server")

	err = s.Serve(ln)

	if err != nil {
		log.Fatal(err)
	}
}
