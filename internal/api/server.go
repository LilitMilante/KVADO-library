package api

import (
	"KVADO-library/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServer(srv proto.LibraryServer) *grpc.Server {
	creds := grpc.Creds(insecure.NewCredentials())
	grpcServer := grpc.NewServer(creds)
	proto.RegisterLibraryServer(grpcServer, srv)
	return grpcServer
}
