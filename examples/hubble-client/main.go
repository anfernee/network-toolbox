package main

import (
	"context"
	"flag"

	observerpb "github.com/cilium/cilium/api/v1/observer"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
)

// Create flag to configure grpc server address
var address = flag.String("grpc-server", "localhost:4245", "grpc server address")

func main() {

	flag.Parse()

	// create grpc connection
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// create grpc client for hubble
	client := observerpb.NewObserverClient(conn)

	cl, err := client.GetFlows(context.Background(), &observerpb.GetFlowsRequest{})
	if err != nil {
		panic(err)
	}

	for {
		resp, err := cl.Recv()
		if err != nil {
			break
		}

		spew.Dump(resp)
	}
}
