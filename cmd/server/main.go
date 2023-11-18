package server

import (
	"context"
	"distributed-kvstore/store"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "path/to/your/protobufs" // replace with your actual protobuf path
)

type kvServer struct {
	pb.UnimplementedKeyValueStoreServer
	store *store.Store
}

func NewServer(store *store.Store) *kvServer {
	return &kvServer{store: store}
}

func (s *kvServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	s.store.Put(req.Key, req.Value)
	return &pb.PutResponse{}, nil
}

func (s *kvServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	value, _ := s.store.Get(req.Key)
	return &pb.GetResponse{Value: value}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	kvStore := store.New()
	kvSvc := NewServer(kvStore)

	pb.RegisterKeyValueStoreServer(grpcServer, kvSvc)
	log.Println("Server is running on port 5000")
	grpcServer.Serve(lis)
}
