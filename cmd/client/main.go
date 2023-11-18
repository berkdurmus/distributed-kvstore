package client

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "path/to/your/protobufs" // replace with your actual protobuf path
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKeyValueStoreClient(conn)

	// Put operation
	_, err = c.Put(context.Background(), &pb.PutRequest{Key: "exampleKey", Value: "exampleValue"})
	if err != nil {
		log.Fatalf("could not put: %v", err)
	}

	// Get operation
	getResp, err := c.Get(context.Background(), &pb.GetRequest{Key: "exampleKey"})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Got value: %s for key: %s", getResp.Value, "exampleKey")
}
